// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !race

package atomic

import _ "unsafe" // for linkname

// load128 atomically loads and returns the pair at *ptr.
//
//go:linkname load128 internal/runtime/atomic.Load128
func load128(ptr *[2]uint64) (lo, hi uint64)

// store128 atomically stores (lo, hi) at *ptr.
//
//go:linkname store128 internal/runtime/atomic.Store128
func store128(ptr *[2]uint64, lo, hi uint64)

// swap128 atomically stores (new1, new2) at *ptr and returns the old pair.
//
//go:linkname swap128 internal/runtime/atomic.Swap128
func swap128(ptr *[2]uint64, new1, new2 uint64) (old1, old2 uint64)

// cas128 atomically compares (*ptr) to (old1, old2) and, if equal, stores
// (new1, new2). ptr must be 16-byte aligned.
//
//go:linkname cas128 internal/runtime/atomic.Cas128
func cas128(ptr *[2]uint64, old1, old2, new1, new2 uint64) bool
