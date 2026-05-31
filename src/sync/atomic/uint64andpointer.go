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
	_ align128

	// u holds the uint64 value (element [0]) and p holds the pointer (element [1]).
	u uint64
	p *T
}

// addr returns the 16-byte-aligned *[2]uint64 inside x that holds the atomic pair,
// where element [0] is the uint64 value and element [1] is the pointer value.
// The struct's align128 field guarantees that x is 16-byte aligned, so u is always
// 16-byte aligned.
//
//go:nosplit
func (x *Uint64AndPointer[T]) addr() *[2]uint64 {
	return (*[2]uint64)(unsafe.Pointer(&x.u))
}

// ptrSlot returns the pointer slot of the live pair as *unsafe.Pointer.
//
//go:nosplit
func (x *Uint64AndPointer[T]) ptrSlot() *unsafe.Pointer {
	return (*unsafe.Pointer)(unsafe.Pointer(&x.p))
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
	if writeBarrier.enabled {
		atomicwb(x.ptrSlot(), unsafe.Pointer(v2))
	}
	store128(x.addr(), v1, uint64(uintptr(unsafe.Pointer(v2))))
}

// Swap atomically stores (new1, new2) into x and returns the previous pair.
func (x *Uint64AndPointer[T]) Swap(new1 uint64, new2 *T) (old1 uint64, old2 *T) {
	addr := x.addr()
	pslot := x.ptrSlot()
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
	if writeBarrier.enabled {
		atomicwb(x.ptrSlot(), unsafe.Pointer(new2))
	}
	addr := x.addr()
	return cas128(addr,
		old1, uint64(uintptr(unsafe.Pointer(old2))),
		new1, uint64(uintptr(unsafe.Pointer(new2))),
	)
}
