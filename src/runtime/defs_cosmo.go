// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

/*
Input to cgo -godefs.

GOARCH=amd64 go tool cgo -godefs defs_cosmo.go >defs_cosmo_amd64.go

This is only a helper to create defs_cosmo_amd64.go
Go runtime functions require the "linux" name of fields (ss_sp, si_addr, etc)
However, auto generated cosmo structures have Go style names.

For example:

// C def
struct itimerval {
  struct timeval it_interval;
  struct timeval it_value;
};

// Auto generated Go code
type itimerval struct {
	Interval	timeval
	Value		timeval
}

TODO(mauri870): create a script to automatise defs_cosmo creation.
TODO(mauri870): read if "modifications made" in defs_aix is necessary,
they used this same trick.
*/

package runtime

/*
#include <cosmopolitan.h>
*/
import "C"

// TODO(mauri870): some constants are hardcoded here
// They result in _Cmacro_ definitions when imported from C, eg C.EINTR.
// Need to check how to lookup these errno values from cosmo.
// It also happens with symbolic definitions such as MADV_DONTNEED.
const (
	EINTR   = 0x4
	EFAULT  = 0xe
	_EAGAIN = 0x23
	_ENOMEM = 0xc

	_O_WRONLY  = 0x1
	O_NONBLOCK = 0x4
	_O_CREAT   = 0x200
	_O_TRUNC   = 0x400
	O_CLOEXEC  = 0x80000

	_PROT_NONE  = C.PROT_NONE
	_PROT_READ  = C.PROT_READ
	_PROT_WRITE = C.PROT_WRITE
	_PROT_EXEC  = C.PROT_EXEC

	_MAP_ANON    = 0x20
	_MAP_PRIVATE = 0x2
	_MAP_FIXED   = C.MAP_FIXED

	_MADV_DONTNEED   = 0x4
	_MADV_FREE       = 0x8
	_MADV_HUGEPAGE   = 0xe
	_MADV_NOHUGEPAGE = 0xf
	_MADV_COLLAPSE   = 0x19

	_SA_SIGINFO = 0x4
	_SA_RESTART = 0x10000000
	_SA_ONSTACK = 0x8000000

	_SI_USER = 0 // C.SI_USER

	_SIGHUP  = C.SIGHUP
	_SIGINT  = C.SIGINT
	_SIGQUIT = C.SIGQUIT
	_SIGILL  = C.SIGILL
	_SIGTRAP = C.SIGTRAP
	_SIGABRT = C.SIGABRT
	// SIGEMT    = C.SIGEMT
	_SIGFPE  = C.SIGFPE
	_SIGKILL = C.SIGKILL
	// SIGBUS    = C.SIGBUS
	_SIGSEGV = C.SIGSEGV
	// SIGSYS    = C.SIGSYS
	_SIGPIPE = C.SIGPIPE
	_SIGALRM = C.SIGALRM
	_SIGTERM = C.SIGTERM
	_SIGURG  = 0x17
	// SIGSTOP   = C.SIGSTOP
	// SIGTSTP   = C.SIGTSTP
	// SIGCONT   = C.SIGCONT
	// SIGCHLD   = C.SIGCHLD
	_SIGTTIN = C.SIGTTIN
	_SIGTTOU = C.SIGTTOU
	// SIGIO     = C.SIGIO
	_SIGXCPU   = C.SIGXCPU
	_SIGXFSZ   = C.SIGXFSZ
	_SIGVTALRM = C.SIGVTALRM
	_SIGPROF   = C.SIGPROF
	_SIGWINCH  = C.SIGWINCH
	// SIGINFO   = C.SIGINFO
	// SIGUSR1   = C.SIGUSR1
	// SIGUSR2   = C.SIGUSR2

	// FPE_INTDIV = C.FPE_INTDIV
	// FPE_INTOVF = C.FPE_INTOVF
	// FPE_FLTDIV = C.FPE_FLTDIV
	// FPE_FLTOVF = C.FPE_FLTOVF
	// FPE_FLTUND = C.FPE_FLTUND
	// FPE_FLTRES = C.FPE_FLTRES
	// FPE_FLTINV = C.FPE_FLTINV
	// FPE_FLTSUB = C.FPE_FLTSUB

	BUS_ADRALN = C.BUS_ADRALN
	BUS_ADRERR = C.BUS_ADRERR
	BUS_OBJERR = C.BUS_OBJERR

	SEGV_MAPERR = C.SEGV_MAPERR
	SEGV_ACCERR = C.SEGV_ACCERR

	_ITIMER_REAL    = C.ITIMER_REAL
	_ITIMER_VIRTUAL = C.ITIMER_VIRTUAL
	_ITIMER_PROF    = C.ITIMER_PROF

	_NSIG = C.NSIG

	// TODO(mauri870): these are defined as SYMBOLIC in cosmopolitan.h.
	// Need to find a way to import these constants.
	_FUTEX_WAIT         = 0
	_FUTEX_WAKE         = 1
	_FUTEX_PRIVATE_FLAG = 128
	_FUTEX_WAIT_PRIVATE = (_FUTEX_WAIT | _FUTEX_PRIVATE_FLAG) // C.FUTEX_WAIT_PRIVATE
	_FUTEX_WAKE_PRIVATE = (_FUTEX_WAKE | _FUTEX_PRIVATE_FLAG) // C.FUTEX_WAKE_PRIVATE
)

type sigset C.sigset_t
type siginfo C.struct_siginfo
type sigactiont C.struct_sigaction
type stackt C.stack_t

type timespec C.struct_timespec
type timeval C.struct_timeval
type itimerval C.struct_itimerval

// TODO(mauri870): sigcontext is not defined in cosmopolitan.h, not sure how to handle it.
type sigcontext struct{}
