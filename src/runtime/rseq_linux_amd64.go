// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "internal/cpu"

// _RSEQ_SIG is the architecture-specific signature value that must appear
// in the instruction stream immediately before each rseq abort handler.
// On amd64 this is the encoding of: nop dword [rax+0x53053053]
// (a no-op that is distinct from any valid instruction prefix sequence).
const _RSEQ_SIG = 0x53053053

// getcpuidFallback is called by getcpuid when rseq is not registered.
// On amd64 it uses the RDPID instruction (CPUID leaf 7, ECX bit 22) to
// read IA32_TSC_AUX, which the kernel keeps set to cpu_id|(numa<<12).
// Returns -1 if RDPID is not supported by this CPU.
//
//go:nosplit
func getcpuidFallback() int32 {
	if cpu.X86.HasRDPID {
		return rdpid()
	}
	return -1
}

// rdpid executes the RDPID instruction and returns the current CPU ID.
// The instruction reads IA32_TSC_AUX; the kernel stores cpu_id in bits
// [11:0] and the NUMA node in bits [31:12].
//
//go:nosplit
func rdpid() int32
