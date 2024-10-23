// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux

#include "textflag.h"

// TODO(minux): this is only valid for ARMv6+
// bool armcas(int32 *val, int32 old, int32 new)
// Atomically:
//	if(*val == old){
//		*val = new;
//		return 1;
//	}else
//		return 0;
TEXT	·Cas(SB),NOSPLIT,$0
	JMP	·armcas(SB)

// Non-linux OSes support only single processor machines before ARMv7.
// So we don't need memory barriers if goarm < 7. And we fail loud at
// startup (runtime.checkgoarm) if it is a multi-processor but goarm < 7.

TEXT	·Load(SB),NOSPLIT|NOFRAME,$0-8
	MOVW	addr+0(FP), R0
	MOVW	(R0), R1

	MOVB	runtime·goarm(SB), R11
	CMP	$7, R11
	BLT	2(PC)
	DMB	MB_ISH

	MOVW	R1, ret+4(FP)
	RET

TEXT	·Store(SB),NOSPLIT,$0-8
	MOVW	addr+0(FP), R1
	MOVW	v+4(FP), R2

	MOVB	runtime·goarm(SB), R8
	CMP	$7, R8
	BLT	2(PC)
	DMB	MB_ISH

	MOVW	R2, (R1)

	CMP	$7, R8
	BLT	2(PC)
	DMB	MB_ISH
	RET

TEXT	·Load8(SB),NOSPLIT|NOFRAME,$0-5
	MOVW	addr+0(FP), R0
	MOVB	(R0), R1

	MOVB	runtime·goarm(SB), R11
	CMP	$7, R11
	BLT	2(PC)
	DMB	MB_ISH

	MOVB	R1, ret+4(FP)
	RET

TEXT	·Store8(SB),NOSPLIT,$0-5
	MOVW	addr+0(FP), R1
	MOVB	v+4(FP), R2

	MOVB	runtime·goarm(SB), R8
	CMP	$7, R8
	BLT	2(PC)
	DMB	MB_ISH

	MOVB	R2, (R1)

	CMP	$7, R8
	BLT	2(PC)
	DMB	MB_ISH
	RET

TEXT ·And8(SB),NOSPLIT,$-4-5
	NO_LOCAL_POINTERS
	MOVW	addr+0(FP), R1

// Uses STREXB/LDREXB that is armv6k or later.
// For simplicity we only enable this on armv7.
#ifndef GOARM_7
	MOVB	internal∕cpu·ARM+const_offsetARMHasV7Atomics(SB), R11
	CMP	$1, R11
	BEQ	2(PC)
	JMP	·goAnd8(SB)
#endif
	JMP	armAnd8<>(SB)


TEXT ·Or8(SB),NOSPLIT,$-4-5
	NO_LOCAL_POINTERS
	MOVW	addr+0(FP), R1

// Uses STREXB/LDREXB that is armv6k or later.
// For simplicity we only enable this on armv7.
#ifndef GOARM_7
	MOVB	internal∕cpu·ARM+const_offsetARMHasV7Atomics(SB), R11
	CMP	$1, R11
	BEQ	2(PC)
	JMP	·goOr8(SB)
#endif
	JMP	armOr8<>(SB)

