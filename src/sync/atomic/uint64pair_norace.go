// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !race

package atomic

import _ "unsafe" // for linkname

// cas128 atomically compares (*addr, *(addr+8)) to (old1, old2) and, if
// equal, stores (new1, new2). addr must be 16-byte aligned.
//
//go:linkname cas128 internal/runtime/atomic.Cas128
func cas128(addr *uint64, old1, old2, new1, new2 uint64) bool
