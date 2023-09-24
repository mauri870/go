// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(mauri870): why this line is needed here?
// Perhaps the detection of `_cosmo.go` in filenames is not working?
//go:build cosmo

package runtime

import (
	"unsafe"
	_ "unsafe"
) // for go:cgo_export_static and go:cgo_export_dynamic

// Export the main function.

//go:cgo_export_static main.main
//go:cgo_export_dynamic main.main

type mOS struct {
	waitsemacount uint32
}

const (
	_ESRCH       = 3
	_EWOULDBLOCK = _EAGAIN
	_ENOTSUP     = 91

	// From OpenBSD's sys/time.h
	_CLOCK_REALTIME  = 0
	_CLOCK_VIRTUAL   = 1
	_CLOCK_PROF      = 2
	_CLOCK_MONOTONIC = 3
)

// type sigset uint32

// var sigset_all = ^sigset(0)

// From OpenBSD's <sys/sysctl.h>
const (
	_CTL_KERN   = 1
	_KERN_OSREV = 3

	_CTL_HW        = 6
	_HW_NCPU       = 3
	_HW_PAGESIZE   = 7
	_HW_NCPUONLINE = 25
)

func sysctlInt(mib []uint32) (int32, bool) {
	var out int32
	nout := unsafe.Sizeof(out)
	ret := sysctl(&mib[0], uint32(len(mib)), (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
	if ret < 0 {
		return 0, false
	}
	return out, true
}

func sysctlUint64(mib []uint32) (uint64, bool) {
	var out uint64
	nout := unsafe.Sizeof(out)
	ret := sysctl(&mib[0], uint32(len(mib)), (*byte)(unsafe.Pointer(&out)), &nout, nil, 0)
	if ret < 0 {
		return 0, false
	}
	return out, true
}

//go:linkname internal_cpu_sysctlUint64 internal/cpu.sysctlUint64
func internal_cpu_sysctlUint64(mib []uint32) (uint64, bool) {
	return sysctlUint64(mib)
}

func getncpu() int32 {
	// Try hw.ncpuonline first because hw.ncpu would report a number twice as
	// high as the actual CPUs running on OpenBSD 6.4 with hyperthreading
	// disabled (hw.smt=0). See https://golang.org/issue/30127
	if n, ok := sysctlInt([]uint32{_CTL_HW, _HW_NCPUONLINE}); ok {
		return int32(n)
	}
	if n, ok := sysctlInt([]uint32{_CTL_HW, _HW_NCPU}); ok {
		return int32(n)
	}
	return 1
}

func getPageSize() uintptr {
	if ps, ok := sysctlInt([]uint32{_CTL_HW, _HW_PAGESIZE}); ok {
		return uintptr(ps)
	}
	return 0
}

func getOSRev() int {
	if osrev, ok := sysctlInt([]uint32{_CTL_KERN, _KERN_OSREV}); ok {
		return int(osrev)
	}
	return 0
}

func osinit() {
	ncpu = getncpu()
	physPageSize = getPageSize()
	haveMapStack = getOSRev() >= 201805 // OpenBSD 6.3
}

var urandom_dev = []byte("/dev/urandom\x00")

//go:nosplit
func getRandomData(r []byte) {
	fd := open(&urandom_dev[0], 0 /* O_RDONLY */, 0)
	n := read(fd, unsafe.Pointer(&r[0]), int32(len(r)))
	closefd(fd)
	extendRandom(r, int(n))
}

func goenvs() {
	goenvs_unix()
}

// Called to initialize a new m (including the bootstrap m).
// Called on the parent thread (main thread in case of bootstrap), can allocate memory.
func mpreinit(mp *m) {
	gsignalSize := int32(32 * 1024)
	if GOARCH == "mips64" {
		gsignalSize = int32(64 * 1024)
	}
	mp.gsignal = malg(gsignalSize)
	mp.gsignal.m = mp
}

// Called to initialize a new m (including the bootstrap m).
// Called on the new thread, can not allocate memory.
func minit() {
	getg().m.procid = uint64(getthrid())
	minitSignals()
}

// Called from dropm to undo the effect of an minit.
//
//go:nosplit
func unminit() {
	unminitSignals()
	getg().m.procid = 0
}

// Called from exitm, but not from drop, to undo the effect of thread-owned
// resources in minit, semacreate, or elsewhere. Do not take locks after calling this.
func mdestroy(mp *m) {
}

func sigtramp()

//go:nosplit
//go:nowritebarrierrec
func setsig(i uint32, fn uintptr) {
	// TODO(mauri870): sigactiont is defined in defs_cosmo.go but it seems to not be enough.
	// Can it be that cosmopolitan.h is missing a `typedef struct sigaction sigaction_t`?
	//
	// src/runtime/os_cosmo.go:162:5: sa.sa_flags undefined (type sigactiont has no field or method sa_flags)
	// src/runtime/os_cosmo.go:163:5: sa.sa_mask undefined (type sigactiont has no field or method sa_mask)
	// src/runtime/os_cosmo.go:163:22: undefined: sigset_all
	// src/runtime/os_cosmo.go:167:5: sa.sa_sigaction undefined (type sigactiont has no field or method sa_sigaction)

	// var sa sigactiont
	// sa.sa_flags = _SA_SIGINFO | _SA_ONSTACK | _SA_RESTART
	// sa.sa_mask = uint32(sigset_all)
	// if fn == abi.FuncPCABIInternal(sighandler) { // abi.FuncPCABIInternal(sighandler) matches the callers in signal_unix.go
	// 	fn = abi.FuncPCABI0(sigtramp)
	// }
	// sa.sa_sigaction = fn
	// sigaction(i, &sa, nil)
}

//go:nosplit
//go:nowritebarrierrec
func setsigstack(i uint32) {
	throw("setsigstack")
}

//go:nosplit
//go:nowritebarrierrec
func getsig(i uint32) uintptr {
	// TODO(mauri870): sigactiont is defined in defs_cosmo.go but it seems to not be enough
	//
	// src/runtime/os_cosmo.go:162:5: sa.sa_flags undefined (type sigactiont has no field or method sa_flags)
	// src/runtime/os_cosmo.go:163:5: sa.sa_mask undefined (type sigactiont has no field or method sa_mask)
	// src/runtime/os_cosmo.go:163:22: undefined: sigset_all
	// src/runtime/os_cosmo.go:167:5: sa.sa_sigaction undefined (type sigactiont has no field or method sa_sigaction)

	// var sa sigactiont
	// sigaction(i, nil, &sa)
	// return sa.sa_sigaction
	return 0
}

// setSignalstackSP sets the ss_sp field of a stackt.
//
//go:nosplit
func setSignalstackSP(s *stackt, sp uintptr) {
	// TODO(mauri870): this should be s.Sp?
	// s.ss_sp = sp
}

//go:nosplit
//go:nowritebarrierrec
func sigaddset(mask *sigset, i int) {
	// TODO(mauri870): this was copied from netbsd.
	// On netbsd __bits is [4]uint32, while for cosmos it is [2]uint64.
	// Check if the code below and in sigdelset needs adjustments.
	// netbsd code:
	// mask.X__bits[(i-1)/32] |= 1 << ((uint32(i) - 1) & 31)
	mask.X__bits[(i-1)/64] |= 1 << ((uint64(i) - 1) & 63)
}

//go:nosplit
func sigfillset(mask *uint64) {
	*mask = ^uint64(0)
}

func sigdelset(mask *sigset, i int) {
	mask.X__bits[(i-1)/64] &^= 1 << ((uint64(i) - 1) & 63)
}

//go:nosplit
func (c *sigctxt) fixsigcode(sig uint32) {
}

func setProcessCPUProfiler(hz int32) {
	setProcessCPUProfilerTimer(hz)
}

func setThreadCPUProfiler(hz int32) {
	setThreadCPUProfilerHz(hz)
}

//go:nosplit
func validSIGPROF(mp *m, c *sigctxt) bool {
	return true
}

var haveMapStack = false

func osStackRemap(s *mspan, flags int32) {
	if !haveMapStack {
		// OpenBSD prior to 6.3 did not have MAP_STACK and so
		// the following mmap will fail. But it also didn't
		// require MAP_STACK (obviously), so there's no need
		// to do the mmap.
		return
	}
	a, err := mmap(unsafe.Pointer(s.base()), s.npages*pageSize, _PROT_READ|_PROT_WRITE, _MAP_PRIVATE|_MAP_ANON|_MAP_FIXED|flags, -1, 0)
	if err != 0 || uintptr(a) != s.base() {
		print("runtime: remapping stack memory ", hex(s.base()), " ", s.npages*pageSize, " a=", a, " err=", err, "\n")
		throw("remapping stack memory failed")
	}
}

//go:nosplit
func raise(sig uint32) {
	thrkill(getthrid(), int(sig))
}

func signalM(mp *m, sig int) {
	thrkill(int32(mp.procid), sig)
}

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
const sigPerThreadSyscall = 1 << 31

//go:nosplit
func runPerThreadSyscall() {
	throw("runPerThreadSyscall only valid on linux")
}

//go:nosplit
func osyield() {
	sleep(0)
}

//go:nosplit
func osyield_no_g() {
	osyield()
}

// Atomically,
//
//	if(*addr == val) sleep
//
// Might be woken up spuriously; that's allowed.
// Don't sleep longer than ns; ns < 0 means forever.
//
//go:nosplit
func futexsleep(addr *uint32, val uint32, ns int64) {
	// Some Linux kernels have a bug where futex of
	// FUTEX_WAIT returns an internal error code
	// as an errno. Libpthread ignores the return value
	// here, and so can we: as it says a few lines up,
	// spurious wakeups are allowed.
	if ns < 0 {
		futex(unsafe.Pointer(addr), _FUTEX_WAIT_PRIVATE, val, nil, nil, 0)
		return
	}

	var ts timespec
	ts.setNsec(ns)
	futex(unsafe.Pointer(addr), _FUTEX_WAIT_PRIVATE, val, unsafe.Pointer(&ts), nil, 0)
}

// If any procs are sleeping on addr, wake up at most cnt.
//
//go:nosplit
func futexwakeup(addr *uint32, cnt uint32) {
	ret := futex(unsafe.Pointer(addr), _FUTEX_WAKE_PRIVATE, cnt, nil, nil, 0)
	if ret >= 0 {
		return
	}

	// I don't know that futex wakeup can return
	// EAGAIN or EINTR, but if it does, it would be
	// safe to loop and call futex again.
	systemstack(func() {
		print("futexwakeup addr=", addr, " returned ", ret, "\n")
	})

	*(*int32)(unsafe.Pointer(uintptr(0x1006))) = 0x1006
}

// TODO(mauri870): Needs implementation
func sleep(ms int32) int32

// TODO(mauri870): Copied from os_linux.go. Needs implementation
func futex(addr unsafe.Pointer, op int32, val uint32, ts, addr2 unsafe.Pointer, val3 uint32) int32

// TODO(mauri870): missing implementation
func newosproc(mp *m) {}
