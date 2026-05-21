// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"sync"
	. "sync/atomic"
	"testing"
)

func TestUint64AndPointerLoadStore(t *testing.T) {
	type V struct{ x int }
	var p Uint64AndPointer[V]

	u, v := p.Load()
	if u != 0 || v != nil {
		t.Fatalf("zero value: got (%d, %v), want (0, nil)", u, v)
	}

	obj := &V{42}
	p.Store(7, obj)
	if u2, v2 := p.Load(); u2 != 7 || v2 != obj {
		t.Fatalf("after Store(7, obj): got (%d, %p), want (7, %p)", u2, v2, obj)
	}

	p.Store(0, nil)
	if u2, v2 := p.Load(); u2 != 0 || v2 != nil {
		t.Fatalf("after Store(0, nil): got (%d, %v), want (0, nil)", u2, v2)
	}
}

func TestUint64AndPointerSwap(t *testing.T) {
	type V struct{ x int }
	var p Uint64AndPointer[V]
	obj1, obj2 := &V{1}, &V{2}
	p.Store(10, obj1)

	oldU, oldV := p.Swap(20, obj2)
	if oldU != 10 || oldV != obj1 {
		t.Fatalf("Swap returned (%d, %p), want (10, %p)", oldU, oldV, obj1)
	}
	if u, v := p.Load(); u != 20 || v != obj2 {
		t.Fatalf("after Swap: got (%d, %p), want (20, %p)", u, v, obj2)
	}
}

func TestUint64AndPointerCompareAndSwap(t *testing.T) {
	type V struct{ x int }
	var p Uint64AndPointer[V]
	obj1, obj2 := &V{1}, &V{2}
	p.Store(10, obj1)

	if !p.CompareAndSwap(10, obj1, 20, obj2) {
		t.Fatal("CompareAndSwap should succeed when both halves match")
	}
	if u, v := p.Load(); u != 20 || v != obj2 {
		t.Fatalf("after successful CAS: got (%d, %p), want (20, %p)", u, v, obj2)
	}

	if p.CompareAndSwap(10, obj2, 30, nil) {
		t.Fatal("CompareAndSwap should fail on uint64 mismatch")
	}
	if p.CompareAndSwap(20, obj1, 30, nil) {
		t.Fatal("CompareAndSwap should fail on pointer mismatch")
	}
}

func TestUint64AndPointerConcurrent(t *testing.T) {
	// Goroutines increment the uint64 part via CAS while cycling the pointer
	// between two values. After completion, verify the uint64 count is consistent.
	type V struct{ x int }
	var p Uint64AndPointer[V]
	obj1, obj2 := &V{1}, &V{2}
	const initialU = uint64(0)
	p.Store(initialU, obj1)

	const G, N = 32, 1000
	var wg sync.WaitGroup
	for g := 0; g < G; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < N; i++ {
				for {
					u, v := p.Load()
					var newV *V
					if v == obj1 {
						newV = obj2
					} else {
						newV = obj1
					}
					if p.CompareAndSwap(u, v, u+1, newV) {
						break
					}
				}
			}
		}()
	}
	wg.Wait()

	u, _ := p.Load()
	if u != G*N {
		t.Errorf("uint64 counter: got %d, want %d", u, G*N)
	}
}

func BenchmarkUint64AndPointerCompareAndSwap(b *testing.B) {
	type V struct{ x int }
	obj1, obj2 := &V{1}, &V{2}
	var p Uint64AndPointer[V]
	p.Store(0, obj1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for {
			u, v := p.Load()
			var newV *V
			if v == obj1 {
				newV = obj2
			} else {
				newV = obj1
			}
			if p.CompareAndSwap(u, v, u+1, newV) {
				break
			}
		}
	}
}
