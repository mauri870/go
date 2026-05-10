// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "unsafe"

// A Uint64Pair is an atomic pair of uint64 values.
// The zero value is a pair of zeros.
//
// Uint64Pair must not be copied after first use.
type Uint64Pair struct {
	_ noCopy
	// 24 bytes of storage so that a 16-byte-aligned 16-byte region is
	// always available inside, regardless of how Uint64Pair is allocated.
	// addr selects that region.
	v [3]uint64
}

// addr returns the 16-byte-aligned *uint64 inside x.v that holds the
// atomic pair.
//
//go:nosplit
func (x *Uint64Pair) addr() *uint64 {
	if uintptr(unsafe.Pointer(&x.v[0]))&15 == 0 {
		return &x.v[0]
	}
	return &x.v[1]
}

// Load atomically loads and returns the pair stored in x.
func (x *Uint64Pair) Load() (v1, v2 uint64) {
	a := x.addr()
	pair := (*[2]uint64)(unsafe.Pointer(a))
	for {
		v1 = LoadUint64(&pair[0])
		v2 = LoadUint64(&pair[1])
		if cas128(a, v1, v2, v1, v2) {
			return
		}
	}
}

// Store atomically stores the pair v1, v2 into x.
func (x *Uint64Pair) Store(v1, v2 uint64) {
	a := x.addr()
	pair := (*[2]uint64)(unsafe.Pointer(a))
	for {
		old1 := LoadUint64(&pair[0])
		old2 := LoadUint64(&pair[1])
		if cas128(a, old1, old2, v1, v2) {
			return
		}
	}
}

// Swap atomically stores new1, new2 into x and returns the old pair.
func (x *Uint64Pair) Swap(new1, new2 uint64) (old1, old2 uint64) {
	a := x.addr()
	pair := (*[2]uint64)(unsafe.Pointer(a))
	for {
		old1 = LoadUint64(&pair[0])
		old2 = LoadUint64(&pair[1])
		if cas128(a, old1, old2, new1, new2) {
			return
		}
	}
}

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint64Pair) CompareAndSwap(old1, old2, new1, new2 uint64) (swapped bool) {
	return cas128(x.addr(), old1, old2, new1, new2)
}

// cas128 is split across build tags:
//   - !race: bodies in uint64pair_norace.go (linkname'd to
//     internal/runtime/atomic.Cas128)
//   - race:  declaration only in uint64pair_race.go; body provided as
//     a TSan trampoline in runtime/race_amd64.s calling
//     __tsan_go_atomic128_compare_exchange
