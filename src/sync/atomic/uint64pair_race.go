// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package atomic

import _ "unsafe" // for linkname

// In race builds, cas128 is the runtime-provided TSan trampoline that
// forwards to __tsan_go_atomic128_compare_exchange. The Go-side body is
// in runtime/race.go; the asm trampoline is in runtime/race_<arch>.s.
//
//go:linkname cas128 runtime.raceCas128
//go:noescape
func cas128(addr *uint64, old1, old2, new1, new2 uint64) bool
