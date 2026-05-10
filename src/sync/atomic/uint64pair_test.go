// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"sync"
	. "sync/atomic"
	"testing"
)

func TestUint64PairLoadStore(t *testing.T) {
	var p Uint64Pair

	v1, v2 := p.Load()
	if v1 != 0 || v2 != 0 {
		t.Fatalf("zero value: got (%d, %d), want (0, 0)", v1, v2)
	}

	p.Store(7, 11)
	if v1, v2 = p.Load(); v1 != 7 || v2 != 11 {
		t.Fatalf("after Store(7, 11): got (%d, %d), want (7, 11)", v1, v2)
	}

	p.Store(0, ^uint64(0))
	if v1, v2 = p.Load(); v1 != 0 || v2 != ^uint64(0) {
		t.Fatalf("after Store(0, ^0): got (%d, %d), want (0, %d)", v1, v2, ^uint64(0))
	}
}

func TestUint64PairSwap(t *testing.T) {
	var p Uint64Pair
	p.Store(1, 2)

	old1, old2 := p.Swap(3, 4)
	if old1 != 1 || old2 != 2 {
		t.Fatalf("Swap returned (%d, %d), want (1, 2)", old1, old2)
	}
	if v1, v2 := p.Load(); v1 != 3 || v2 != 4 {
		t.Fatalf("after Swap: got (%d, %d), want (3, 4)", v1, v2)
	}
}

func TestUint64PairCompareAndSwap(t *testing.T) {
	var p Uint64Pair
	p.Store(10, 20)

	if !p.CompareAndSwap(10, 20, 30, 40) {
		t.Fatal("CompareAndSwap should succeed when both halves match")
	}
	if v1, v2 := p.Load(); v1 != 30 || v2 != 40 {
		t.Fatalf("after successful CAS: got (%d, %d), want (30, 40)", v1, v2)
	}

	if p.CompareAndSwap(10, 40, 99, 99) {
		t.Fatal("CompareAndSwap should fail on low-half mismatch")
	}
	if p.CompareAndSwap(30, 20, 99, 99) {
		t.Fatal("CompareAndSwap should fail on high-half mismatch")
	}
	if v1, v2 := p.Load(); v1 != 30 || v2 != 40 {
		t.Fatalf("after failed CAS: got (%d, %d), want (30, 40)", v1, v2)
	}
}

func TestUint64PairConcurrent(t *testing.T) {
	// 32 goroutines each bump (lo, hi) -> (lo+1, hi-1) 1000 times via
	// CompareAndSwap. After all goroutines exit, the invariant
	// lo + hi == initialHi must hold iff every successful CAS updated
	// both halves together (i.e. CAS is truly 128-bit atomic).
	var p Uint64Pair
	const initialHi = uint64(0xdeadbeefcafebabe)
	p.Store(0, initialHi)

	const G, N = 32, 1000
	var wg sync.WaitGroup
	for g := 0; g < G; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < N; i++ {
				for {
					lo, hi := p.Load()
					if p.CompareAndSwap(lo, hi, lo+1, hi-1) {
						break
					}
				}
			}
		}()
	}
	wg.Wait()

	v1, v2 := p.Load()
	if v1 != G*N {
		t.Errorf("low half: got %d, want %d", v1, G*N)
	}
	if v2 != initialHi-G*N {
		t.Errorf("high half: got %#x, want %#x", v2, initialHi-G*N)
	}
}

func BenchmarkUint64PairCompareAndSwap(b *testing.B) {
	var p Uint64Pair
	for i := 0; i < b.N; i++ {
		for {
			lo, hi := p.Load()
			if p.CompareAndSwap(lo, hi, lo+1, hi-1) {
				break
			}
		}
	}
}

func BenchmarkUint64PairCompareAndSwapParallel(b *testing.B) {
	var p Uint64Pair
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for {
				lo, hi := p.Load()
				if p.CompareAndSwap(lo, hi, lo+1, hi-1) {
					break
				}
			}
		}
	})
}
