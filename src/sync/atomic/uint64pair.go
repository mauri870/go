// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

// A Uint64Pair is an atomic pair of uint64 values.
// The zero value is a pair of zeros.
//
// Uint64Pair must not be copied after first use.
type Uint64Pair struct {
	_ noCopy
	_ align128
	v [2]uint64
}

// addr returns the 16-byte-aligned *[2]uint64 inside x that holds the atomic
// pair. The struct's align128 field guarantees that x is 16-byte aligned, so
// v is always 16-byte aligned.
//
//go:nosplit
func (x *Uint64Pair) addr() *[2]uint64 {
	return &x.v
}

// Load atomically loads and returns the pair stored in x.
func (x *Uint64Pair) Load() (v1, v2 uint64) {
	return load128(x.addr())
}

// Store atomically stores the pair v1, v2 into x.
func (x *Uint64Pair) Store(v1, v2 uint64) {
	store128(x.addr(), v1, v2)
}

// Swap atomically stores new1, new2 into x and returns the old pair.
func (x *Uint64Pair) Swap(new1, new2 uint64) (old1, old2 uint64) {
	return swap128(x.addr(), new1, new2)
}

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint64Pair) CompareAndSwap(old1, old2, new1, new2 uint64) (swapped bool) {
	return cas128(x.addr(), old1, old2, new1, new2)
}

// The following functions are split across build tags:
//   - !race: bodies in uint64pair_norace.go (linknamed to internal/runtime/atomic)
//   - race:  declarations in uint64pair_race.go; bodies provided by TSan trampolines
//     in runtime/race_<arch>.s and runtime/race.go.
