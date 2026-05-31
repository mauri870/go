// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (amd64 || arm64)

package runtime

import (
	"internal/runtime/syscall/linux"
	"unsafe"
)

// rseqABI is the restartable sequences structure shared with the Linux kernel.
// The kernel atomically writes cpuID on every reschedule of this OS thread.
//
// The struct must be 32-byte aligned; we enforce this via the alignment
// argument to persistentalloc in rseqRegister.
//
// See linux/rseq.h.
type rseqABI struct {
	cpuIDStart uint32
	cpuID      uint32
	rseqCS     uint64
	flags      uint32
	_          [12]byte // pad to 32 bytes
}

// rseqRegister registers the rseq ABI for the current OS thread (mp).
// Called from mstart1 after minit.
func rseqRegister(mp *m) {
	r := (*rseqABI)(persistentalloc(unsafe.Sizeof(rseqABI{}), 32, &memstats.other_sys))
	r1, _, errno := linux.Syscall6(
		linux.SYS_RSEQ,
		uintptr(unsafe.Pointer(r)),
		uintptr(unsafe.Sizeof(rseqABI{})),
		0,          // flags: 0 = register
		_RSEQ_SIG,  // architecture-specific abort signature
		0, 0,
	)
	if errno == 0 && r1 == 0 {
		mp.rseqState = uintptr(unsafe.Pointer(r))
		return
	}

	// Registration failed. If errno is EINVAL, another registration already owns
	// this thread — most likely glibc (2.35+), which calls SYS_rseq in
	// __libc_start_main and pthread_create with its own struct rseq pointer
	// (size=AT_RSEQ_FEATURE_SIZE, align=AT_RSEQ_ALIGN) before the Go runtime
	// starts. The kernel rejects re-registration with a different pointer or size.
	//
	// Coexisting with glibc's registration (reading glibc's struct rseq via the
	// __rseq_offset TLS symbol) is left for a follow-up CL. For now, getcpuid()
	// returns -1 in CGO-enabled builds linked against glibc >= 2.35.
	// Pure Go builds (CGO_ENABLED=0) are not affected.
}

// rseqUnregister unregisters the rseq ABI for mp.
// Called from mexit before the OS thread exits.
func rseqUnregister(mp *m) {
	if mp.rseqState == 0 {
		return
	}
	linux.Syscall6(
		linux.SYS_RSEQ,
		mp.rseqState,
		uintptr(unsafe.Sizeof(rseqABI{})),
		1,         // RSEQ_FLAG_UNREGISTER
		_RSEQ_SIG,
		0, 0,
	)
	mp.rseqState = 0
}

// getcpuid returns the CPU number of the current OS thread.
//
// It first tries the rseq ABI (a single TLS load, ~1 ns). If rseq is not
// registered for this thread — most commonly because glibc pre-registered
// rseq with its own struct before mstart1 ran (CGO_ENABLED=1) — it falls
// back to getcpuidFallback, which is architecture-specific: on amd64 this
// uses the RDPID instruction (~3 cycles, no syscall) when the CPU supports
// it; on arm64 and other architectures it returns -1.
//
// A return value of -1 means no CPU ID is available; callers must fall back
// to a safe alternative (e.g. slot 0 or procPin).
//
//go:nosplit
func getcpuid() int32 {
	if getg().m.rseqState != 0 {
		return int32((*rseqABI)(unsafe.Pointer(getg().m.rseqState)).cpuID)
	}
	return getcpuidFallback()
}
