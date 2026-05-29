// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

// Load128 atomically loads the pair stored at *ptr and returns it.
// ptr must be 16-byte aligned.
//
//go:nosplit
func Load128(ptr *[2]uint64) (lo, hi uint64) {
	for {
		lo = Load64(&ptr[0])
		hi = Load64(&ptr[1])
		if Cas128(ptr, lo, hi, lo, hi) {
			return
		}
	}
}

// Store128 atomically stores (lo, hi) into *ptr.
// ptr must be 16-byte aligned.
//
//go:nosplit
func Store128(ptr *[2]uint64, lo, hi uint64) {
	for {
		old0 := Load64(&ptr[0])
		old1 := Load64(&ptr[1])
		if Cas128(ptr, old0, old1, lo, hi) {
			return
		}
	}
}

// Swap128 atomically stores (new1, new2) into *ptr and returns the old pair.
// ptr must be 16-byte aligned.
//
//go:nosplit
func Swap128(ptr *[2]uint64, new1, new2 uint64) (old1, old2 uint64) {
	for {
		old1 = Load64(&ptr[0])
		old2 = Load64(&ptr[1])
		if Cas128(ptr, old1, old2, new1, new2) {
			return
		}
	}
}

