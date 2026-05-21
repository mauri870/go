// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "unsafe"

// A Uint64AndPointer is an atomic pair of a uint64 value and a typed pointer *T.
// The zero value is a pair of (0, nil).
//
// Uint64AndPointer must not be copied after first use.
type Uint64AndPointer[T any] struct {
	// Mention *T to disallow implicit conversion between Uint64AndPointer types
	// and to ensure the GC knows that the pointer fields hold pointers.
	// See go.dev/issue/56603 for details.
	_ [0]*T

	_ noCopy

	// u0/p0 and u1/p1 are two candidate pairs for the 16-byte-aligned atomic
	// slot. addr() selects the pair whose first element is 16-byte aligned.
	// The unselected pair's fields are always zero/nil.
	// Within each pair: u holds the uint64 value (element [0]) and
	// p holds the pointer (element [1]).
	u0 uint64
	p0 *T
	_  uint64 // shifts pair 1 so that u1 is 16-byte aligned when u0 is not
	u1 uint64
	p1 *T
}

// addr returns the 16-byte-aligned *[2]uint64 inside x that holds the atomic pair,
// where element [0] is the uint64 value and element [1] is the pointer value.
//
//go:nosplit
func (x *Uint64AndPointer[T]) addr() *[2]uint64 {
	if uintptr(unsafe.Pointer(&x.u0))&15 == 0 {
		return (*[2]uint64)(unsafe.Pointer(&x.u0))
	}
	return (*[2]uint64)(unsafe.Pointer(&x.u1))
}

// ptrSlot returns the pointer slot of the live pair as *unsafe.Pointer.
//
//go:nosplit
func (x *Uint64AndPointer[T]) ptrSlot() *unsafe.Pointer {
	addr := x.addr()
	return (*unsafe.Pointer)(unsafe.Pointer(&addr[1]))
}

// Load atomically loads and returns the pair stored in x.
func (x *Uint64AndPointer[T]) Load() (v1 uint64, v2 *T) {
	addr := x.addr()
	pslot := (*unsafe.Pointer)(unsafe.Pointer(&addr[1]))
	for {
		// Read each component using the appropriate atomic load so that the
		// pointer provenance is preserved for checkptr and the GC.
		v1 = LoadUint64(&addr[0])
		v2 = (*T)(LoadPointer(pslot))
		// Verify the two reads are consistent via a no-op CAS.
		if cas128(addr,
			v1, uint64(uintptr(unsafe.Pointer(v2))),
			v1, uint64(uintptr(unsafe.Pointer(v2))),
		) {
			return
		}
	}
}

// Store atomically stores the pair (v1, v2) into x.
func (x *Uint64AndPointer[T]) Store(v1 uint64, v2 *T) {
	addr := x.addr()
	if writeBarrier.enabled {
		atomicwb((*unsafe.Pointer)(unsafe.Pointer(&addr[1])), unsafe.Pointer(v2))
	}
	store128(addr, v1, uint64(uintptr(unsafe.Pointer(v2))))
}

// Swap atomically stores (new1, new2) into x and returns the previous pair.
func (x *Uint64AndPointer[T]) Swap(new1 uint64, new2 *T) (old1 uint64, old2 *T) {
	addr := x.addr()
	pslot := (*unsafe.Pointer)(unsafe.Pointer(&addr[1]))
	if writeBarrier.enabled {
		atomicwb(pslot, unsafe.Pointer(new2))
	}
	for {
		// Read each component using the appropriate atomic load so that the
		// pointer provenance is preserved for checkptr and the GC.
		old1 = LoadUint64(&addr[0])
		old2 = (*T)(LoadPointer(pslot))
		if cas128(addr,
			old1, uint64(uintptr(unsafe.Pointer(old2))),
			new1, uint64(uintptr(unsafe.Pointer(new2))),
		) {
			return
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint64AndPointer[T]) CompareAndSwap(old1 uint64, old2 *T, new1 uint64, new2 *T) (swapped bool) {
	addr := x.addr()
	if writeBarrier.enabled {
		atomicwb((*unsafe.Pointer)(unsafe.Pointer(&addr[1])), unsafe.Pointer(new2))
	}
	return cas128(addr,
		old1, uint64(uintptr(unsafe.Pointer(old2))),
		new1, uint64(uintptr(unsafe.Pointer(new2))),
	)
}
