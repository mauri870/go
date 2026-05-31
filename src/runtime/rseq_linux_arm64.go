// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// _RSEQ_SIG is the architecture-specific signature value that must appear
// in the instruction stream immediately before each rseq abort handler.
// On arm64 this is the encoding of: brk #0x4de0
const _RSEQ_SIG = 0xd428bc00
