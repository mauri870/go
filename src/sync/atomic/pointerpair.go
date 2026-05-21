// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "unsafe"

// A PointerPair is an atomic pair of typed pointers (*T1, *T2).
// The zero value is a pair of nil pointers.
//
// PointerPair must not be copied after first use.
type PointerPair[T1, T2 any] struct {
	// Mention *T1, *T2 to disallow implicit conversion between PointerPair types
	// and to ensure the GC knows the fields hold pointers.
	// See go.dev/issue/56603 for details.
	_ [0]*T1
	_ [0]*T2

	_ noCopy

	// p10/p20 and p11/p21 are two candidate pairs for the 16-byte-aligned atomic
	// slot. addr() selects the pair whose first element is 16-byte aligned.
	// The unselected pair's fields are always nil.
	p10 *T1
	p20 *T2
	_   uint64 // shifts pair 1 so that p11 is 16-byte aligned when p10 is not
	p11 *T1
	p21 *T2
}

// addr returns the 16-byte-aligned *[2]unsafe.Pointer inside x that holds the
// atomic pair.
//
//go:nosplit
func (x *PointerPair[T1, T2]) addr() *[2]unsafe.Pointer {
	if uintptr(unsafe.Pointer(&x.p10))&15 == 0 {
		return (*[2]unsafe.Pointer)(unsafe.Pointer(&x.p10))
	}
	return (*[2]unsafe.Pointer)(unsafe.Pointer(&x.p11))
}

// Load atomically loads and returns the pair stored in x.
func (x *PointerPair[T1, T2]) Load() (v1 *T1, v2 *T2) {
	addr := x.addr()
	for {
		// Read each pointer through LoadPointer so the value provenance is
		// preserved for checkptr and the GC write barrier.
		v1 = (*T1)(LoadPointer(&addr[0]))
		v2 = (*T2)(LoadPointer(&addr[1]))
		// Verify the two reads are consistent via a no-op CAS.
		if cas128(
			(*[2]uint64)(unsafe.Pointer(addr)),
			uint64(uintptr(unsafe.Pointer(v1))),
			uint64(uintptr(unsafe.Pointer(v2))),
			uint64(uintptr(unsafe.Pointer(v1))),
			uint64(uintptr(unsafe.Pointer(v2))),
		) {
			return
		}
	}
}

// Store atomically stores the pair (v1, v2) into x.
func (x *PointerPair[T1, T2]) Store(v1 *T1, v2 *T2) {
	addr := x.addr()
	if writeBarrier.enabled {
		atomicwb(&addr[0], unsafe.Pointer(v1))
		atomicwb(&addr[1], unsafe.Pointer(v2))
	}
	store128(
		(*[2]uint64)(unsafe.Pointer(addr)),
		uint64(uintptr(unsafe.Pointer(v1))),
		uint64(uintptr(unsafe.Pointer(v2))),
	)
}

// Swap atomically stores (new1, new2) into x and returns the previous pair.
func (x *PointerPair[T1, T2]) Swap(new1 *T1, new2 *T2) (old1 *T1, old2 *T2) {
	addr := x.addr()
	if writeBarrier.enabled {
		atomicwb(&addr[0], unsafe.Pointer(new1))
		atomicwb(&addr[1], unsafe.Pointer(new2))
	}
	for {
		// Read the current pair through LoadPointer to maintain pointer provenance.
		old1 = (*T1)(LoadPointer(&addr[0]))
		old2 = (*T2)(LoadPointer(&addr[1]))
		if cas128(
			(*[2]uint64)(unsafe.Pointer(addr)),
			uint64(uintptr(unsafe.Pointer(old1))),
			uint64(uintptr(unsafe.Pointer(old2))),
			uint64(uintptr(unsafe.Pointer(new1))),
			uint64(uintptr(unsafe.Pointer(new2))),
		) {
			return
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *PointerPair[T1, T2]) CompareAndSwap(old1 *T1, old2 *T2, new1 *T1, new2 *T2) (swapped bool) {
	addr := x.addr()
	if writeBarrier.enabled {
		atomicwb(&addr[0], unsafe.Pointer(new1))
		atomicwb(&addr[1], unsafe.Pointer(new2))
	}
	return cas128(
		(*[2]uint64)(unsafe.Pointer(addr)),
		uint64(uintptr(unsafe.Pointer(old1))),
		uint64(uintptr(unsafe.Pointer(old2))),
		uint64(uintptr(unsafe.Pointer(new1))),
		uint64(uintptr(unsafe.Pointer(new2))),
	)
}
