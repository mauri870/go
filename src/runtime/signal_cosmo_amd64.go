// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(mauri870): why this line is needed here?
// Perhaps the detection of `_cosmo.go` in filenames is not working?
//go:build cosmo

package runtime

import "unsafe"

type sigctxt struct {
	info *siginfo
	ctxt unsafe.Pointer
}

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) regs() *sigcontext {
	return (*sigcontext)(c.ctxt)
}

// TODO(mauri870): all these are missing implementations
func (c *sigctxt) rax() uint64 { return 0 }
func (c *sigctxt) rbx() uint64 { return 0 }
func (c *sigctxt) rcx() uint64 { return 0 }
func (c *sigctxt) rdx() uint64 { return 0 }
func (c *sigctxt) rdi() uint64 { return 0 }
func (c *sigctxt) rsi() uint64 { return 0 }
func (c *sigctxt) rbp() uint64 { return 0 }
func (c *sigctxt) rsp() uint64 { return 0 }
func (c *sigctxt) r8() uint64  { return 0 }
func (c *sigctxt) r9() uint64  { return 0 }
func (c *sigctxt) r10() uint64 { return 0 }
func (c *sigctxt) r11() uint64 { return 0 }
func (c *sigctxt) r12() uint64 { return 0 }
func (c *sigctxt) r13() uint64 { return 0 }
func (c *sigctxt) r14() uint64 { return 0 }
func (c *sigctxt) r15() uint64 { return 0 }

//go:nosplit
//go:nowritebarrierrec
func (c *sigctxt) rip() uint64 { return 0 }

func (c *sigctxt) rflags() uint64  { return 0 }
func (c *sigctxt) cs() uint64      { return 0 }
func (c *sigctxt) fs() uint64      { return 0 }
func (c *sigctxt) gs() uint64      { return 0 }
func (c *sigctxt) sigcode() uint64 { return 0 }
func (c *sigctxt) sigaddr() uint64 {
	return *(*uint64)(add(unsafe.Pointer(c.info), 16))
}

func (c *sigctxt) set_rip(x uint64)     {}
func (c *sigctxt) set_rsp(x uint64)     {}
func (c *sigctxt) set_sigcode(x uint64) {}
func (c *sigctxt) set_sigaddr(x uint64) {
	*(*uint64)(add(unsafe.Pointer(c.info), 16)) = x
}
