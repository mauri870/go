// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux || (!amd64 && !arm64)

package runtime

func rseqRegister(mp *m)   {}
func rseqUnregister(mp *m) {}

//go:nosplit
func getcpuid() int32 { return -1 }
