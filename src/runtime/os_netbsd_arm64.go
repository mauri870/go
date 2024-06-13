// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"internal/abi"
	"unsafe"
)

func lwp_mcontext_init(mc *mcontextt, stk unsafe.Pointer, mp *m, gp *g, fn uintptr) {
	// Machine dependent mcontext initialisation for LWP.
	mc.__gregs[_REG_ELR] = uint64(abi.FuncPCABI0(lwp_tramp))
	mc.__gregs[_REG_X31] = uint64(uintptr(stk))
	mc.__gregs[_REG_X0] = uint64(uintptr(unsafe.Pointer(mp)))
	mc.__gregs[_REG_X1] = uint64(uintptr(unsafe.Pointer(mp.g0)))
	mc.__gregs[_REG_X2] = uint64(fn)
}
