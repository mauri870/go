// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "textflag.h"

// func rdpid() int32
//
// Executes the RDPID instruction, which reads IA32_TSC_AUX into AX.
// The Linux kernel stores cpu_id in bits [11:0] and the NUMA node in
// bits [31:12]. We mask off the NUMA bits and return the CPU ID.
//
// RDPID is non-serializing and has ~3-cycle latency on modern Intel and
// AMD microarchitectures (Ice Lake+, Zen 2+). The caller must have checked
// cpu.X86.HasRDPID before calling this function.
TEXT runtime·rdpid(SB),NOSPLIT,$0-4
	RDPID AX
	ANDL  $0xfff, AX  // mask off NUMA node bits [31:12]
	MOVL  AX, ret+0(FP)
	RET
