// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"sync"
	. "sync/atomic"
	"testing"
)

func TestPointerPairLoadStore(t *testing.T) {
	type A struct{ x int }
	type B struct{ y int }
	var p PointerPair[A, B]

	v1, v2 := p.Load()
	if v1 != nil || v2 != nil {
		t.Fatalf("zero value: got (%v, %v), want (nil, nil)", v1, v2)
	}

	a := &A{1}
	b := &B{2}
	p.Store(a, b)
	if r1, r2 := p.Load(); r1 != a || r2 != b {
		t.Fatalf("after Store: got (%p, %p), want (%p, %p)", r1, r2, a, b)
	}

	p.Store(nil, nil)
	if r1, r2 := p.Load(); r1 != nil || r2 != nil {
		t.Fatalf("after Store(nil, nil): got (%v, %v), want (nil, nil)", r1, r2)
	}
}

func TestPointerPairSwap(t *testing.T) {
	type A struct{ x int }
	type B struct{ y int }
	var p PointerPair[A, B]

	a1, b1 := &A{1}, &B{1}
	a2, b2 := &A{2}, &B{2}
	p.Store(a1, b1)

	old1, old2 := p.Swap(a2, b2)
	if old1 != a1 || old2 != b1 {
		t.Fatalf("Swap returned (%p, %p), want (%p, %p)", old1, old2, a1, b1)
	}
	if r1, r2 := p.Load(); r1 != a2 || r2 != b2 {
		t.Fatalf("after Swap: got (%p, %p), want (%p, %p)", r1, r2, a2, b2)
	}
}

func TestPointerPairCompareAndSwap(t *testing.T) {
	type A struct{ x int }
	type B struct{ y int }
	var p PointerPair[A, B]

	a1, b1 := &A{1}, &B{1}
	a2, b2 := &A{2}, &B{2}
	p.Store(a1, b1)

	if !p.CompareAndSwap(a1, b1, a2, b2) {
		t.Fatal("CompareAndSwap should succeed when both match")
	}
	if r1, r2 := p.Load(); r1 != a2 || r2 != b2 {
		t.Fatalf("after successful CAS: got (%p, %p), want (%p, %p)", r1, r2, a2, b2)
	}

	if p.CompareAndSwap(a1, b2, a2, b2) {
		t.Fatal("CompareAndSwap should fail on first-element mismatch")
	}
	if p.CompareAndSwap(a2, b1, a2, b2) {
		t.Fatal("CompareAndSwap should fail on second-element mismatch")
	}
}

func TestPointerPairConcurrent(t *testing.T) {
	type Node struct{ val int }
	a := &Node{1}
	b := &Node{2}
	c := &Node{3}
	d := &Node{4}
	var p PointerPair[Node, Node]
	p.Store(a, b)

	const G = 32
	var wg sync.WaitGroup
	for g := 0; g < G; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				r1, r2 := p.Load()
				if r1 == a && r2 == b {
					p.CompareAndSwap(a, b, c, d)
				} else if r1 == c && r2 == d {
					p.CompareAndSwap(c, d, a, b)
				}
			}
		}()
	}
	wg.Wait()

	r1, r2 := p.Load()
	if (r1 != a || r2 != b) && (r1 != c || r2 != d) {
		t.Fatalf("Load returned inconsistent pair (%p, %p)", r1, r2)
	}
}

func BenchmarkPointerPairCompareAndSwap(b *testing.B) {
	type V struct{ x int }
	v1, v2 := &V{1}, &V{2}
	v3, v4 := &V{3}, &V{4}
	var p PointerPair[V, V]
	p.Store(v1, v2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for {
			r1, r2 := p.Load()
			var n1, n2 *V
			if r1 == v1 {
				n1, n2 = v3, v4
			} else {
				n1, n2 = v1, v2
			}
			if p.CompareAndSwap(r1, r2, n1, n2) {
				break
			}
		}
	}
}
