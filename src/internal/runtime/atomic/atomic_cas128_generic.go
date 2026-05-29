// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64 && !arm64

package atomic

import "unsafe"

// Cas128 atomically compares the 16 bytes at *ptr to (old1, old2) and,
// if equal, replaces them with (new1, new2). On architectures without a
// native 128-bit atomic instruction, this delegates to the lock-table
// fallback in atomic_cas128_native.go. ptr must be 16-byte aligned.
//
//go:nosplit
func Cas128(ptr *[2]uint64, old1, old2, new1, new2 uint64) bool {
	return goCas128(ptr, old1, old2, new1, new2)
}

// Cas128p atomically compares the 16 bytes at *ptr to (old1, old2) and,
// if equal, replaces them with (new1, new2). On architectures without a
// native 128-bit atomic instruction, this delegates to the lock-table
// fallback in atomic_cas128_native.go. ptr must be 16-byte aligned.
//
//go:nosplit
func Cas128p(ptr *[2]unsafe.Pointer, old1, old2, new1, new2 unsafe.Pointer) bool {
	if uintptr(unsafe.Pointer(ptr))&15 != 0 {
		panicUnaligned128()
	}
	_ = *ptr // fault on nil before taking the lock
	l := pairAddrLock((*[2]uint64)(unsafe.Pointer(ptr)))
	l.lock()
	ok := false
	if ptr[0] == old1 && ptr[1] == old2 {
		ptr[0] = new1
		ptr[1] = new2
		ok = true
	}
	l.unlock()
	return ok
}
