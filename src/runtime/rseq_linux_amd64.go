// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// _RSEQ_SIG is the architecture-specific signature value that must appear
// in the instruction stream immediately before each rseq abort handler.
// On amd64 this is the encoding of: nop dword [rax+0x53053053]
// (a no-op that is distinct from any valid instruction prefix sequence).
const _RSEQ_SIG = 0x53053053
