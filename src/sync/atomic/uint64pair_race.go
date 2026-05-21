// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

package atomic

import _ "unsafe" // for linkname

// In race builds the 128-bit operations are provided by the runtime race
// package, which forwards to TSan's atomic128 functions for correct race
// detection. The Go-side bodies are in runtime/race.go; the assembly
// trampolines are in runtime/race_<arch>.s.

// load128 is the TSan-instrumented 128-bit atomic load.
//
//go:linkname load128 runtime.raceLoad128
//go:noescape
func load128(ptr *[2]uint64) (lo, hi uint64)

// store128 is the TSan-instrumented 128-bit atomic store.
//
//go:linkname store128 runtime.raceStore128
//go:noescape
func store128(ptr *[2]uint64, lo, hi uint64)

// swap128 is the TSan-instrumented 128-bit atomic swap.
//
//go:linkname swap128 runtime.raceSwap128
func swap128(ptr *[2]uint64, new1, new2 uint64) (old1, old2 uint64)

// cas128 is the TSan-instrumented 128-bit compare-and-swap.
//
//go:linkname cas128 runtime.raceCas128
//go:noescape
func cas128(ptr *[2]uint64, old1, old2, new1, new2 uint64) bool
