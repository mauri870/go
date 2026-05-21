// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic

import "unsafe"

// writeBarrier is a copy of the runtime variable of the same name.
// The fields must match runtime.writeBarrier exactly.
//
//go:linkname writeBarrier runtime.writeBarrier
var writeBarrier struct {
	enabled bool    // compiler emits a check of this before calling write barrier
	pad     [3]byte // compiler uses 32-bit load for "enabled" field
	alignme uint64  // guarantee alignment so that compiler can use a 32 or 64-bit load
}

// atomicwb performs a write barrier before an atomic pointer write.
// The caller should guard the call with "if writeBarrier.enabled".
//
//go:linkname atomicwb runtime.atomicwb
//go:nosplit
func atomicwb(ptr *unsafe.Pointer, new unsafe.Pointer)
