// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race

#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

#define RARG0 DI
#define RARG1 SI
#define RARG2 DX
#define RARG3 CX

TEXT	sync∕atomic·AndInt64(SB), NOSPLIT, $0-24
	GO_ARGS

	MOVQ 	addr+0(FP), R12
	MOVQ 	mask+8(FP), R13
	MOVQ 	(R12), R15
	ANDQ 	R15, R13

	MOVQ 	R12, 16(SP)
	MOVQ 	R15, 24(SP)
	MOVQ 	R13, 32(SP)

	MOVQ	g_racectx(R14), RARG0	// goroutine context
	MOVQ	8(SP), RARG1	// caller pc
	MOVQ	(SP), RARG2	// pc
	LEAQ	16(SP), RARG3	// arguments
	MOVQ	$__tsan_go_atomic64_compare_exchange(SB), AX
	CALL	racecall(SB)	// does not return

	MOVQ	R15, ret+16(FP)

	RET

TEXT	sync∕atomic·AndInt32(SB), NOSPLIT, $0-20
	GO_ARGS

	MOVQ 	addr+0(FP), R12
	MOVL 	mask+8(FP), R13
	MOVL 	(R12), R15
	ANDL 	R15, R13

	MOVQ 	R12, 16(SP)
	MOVL 	R15, 24(SP)
	MOVL 	R13, 28(SP)

	MOVQ	g_racectx(R14), RARG0	// goroutine context
	MOVQ	8(SP), RARG1	// caller pc
	MOVQ	(SP), RARG2	// pc
	LEAQ	16(SP), RARG3	// arguments
	MOVQ	$__tsan_go_atomic32_compare_exchange(SB), AX
	CALL	racecall(SB)	// does not return

	MOVQ	R15, ret+16(FP)
	RET

// Or
TEXT	sync∕atomic·OrInt32(SB), NOSPLIT, $0-20
	GO_ARGS

	MOVQ 	addr+0(FP), R12
	MOVL 	mask+8(FP), R13
	MOVL 	(R12), R15
	ORL 	R15, R13

	MOVQ 	R12, 16(SP)
	MOVL 	R15, 24(SP)
	MOVL 	R13, 28(SP)

	MOVQ	g_racectx(R14), RARG0	// goroutine context
	MOVQ	8(SP), RARG1	// caller pc
	MOVQ	(SP), RARG2	// pc
	LEAQ	16(SP), RARG3	// arguments
	MOVQ	$__tsan_go_atomic32_compare_exchange(SB), AX
	CALL	racecall(SB)	// does not return

	MOVQ	R15, ret+16(FP)
	RET

TEXT	sync∕atomic·OrInt64(SB), NOSPLIT, $0-24
	GO_ARGS

	MOVQ 	addr+0(FP), R12
	MOVQ 	mask+8(FP), R13
	MOVQ 	(R12), R15
	ORQ 	R15, R13

	MOVQ 	R12, 16(SP)
	MOVQ 	R15, 24(SP)
	MOVQ 	R13, 32(SP)

	MOVQ	g_racectx(R14), RARG0	// goroutine context
	MOVQ	8(SP), RARG1	// caller pc
	MOVQ	(SP), RARG2	// pc
	LEAQ	16(SP), RARG3	// arguments
	MOVQ	$__tsan_go_atomic64_compare_exchange(SB), AX
	CALL	racecall(SB)	// does not return

	MOVQ	R15, ret+16(FP)
	RET
