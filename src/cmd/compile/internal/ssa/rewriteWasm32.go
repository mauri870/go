// Code generated from _gen/Wasm32.rules using 'go generate'; DO NOT EDIT.

package ssa

import "internal/buildcfg"
import "math"
import "cmd/compile/internal/types"

func rewriteValueWasm32(v *Value) bool {
	switch v.Op {
	case OpAbs:
		v.Op = OpWasm32F64Abs
		return true
	case OpAdd16:
		v.Op = OpWasm32I64Add
		return true
	case OpAdd32:
		v.Op = OpWasm32I64Add
		return true
	case OpAdd32F:
		v.Op = OpWasm32F32Add
		return true
	case OpAdd64:
		v.Op = OpWasm32I64Add
		return true
	case OpAdd64F:
		v.Op = OpWasm32F64Add
		return true
	case OpAdd8:
		v.Op = OpWasm32I64Add
		return true
	case OpAddPtr:
		v.Op = OpWasm32I64Add
		return true
	case OpAddr:
		return rewriteValueWasm32_OpAddr(v)
	case OpAnd16:
		v.Op = OpWasm32I64And
		return true
	case OpAnd32:
		v.Op = OpWasm32I64And
		return true
	case OpAnd64:
		v.Op = OpWasm32I64And
		return true
	case OpAnd8:
		v.Op = OpWasm32I64And
		return true
	case OpAndB:
		v.Op = OpWasm32I64And
		return true
	case OpBitLen64:
		return rewriteValueWasm32_OpBitLen64(v)
	case OpCeil:
		v.Op = OpWasm32F64Ceil
		return true
	case OpClosureCall:
		v.Op = OpWasm32LoweredClosureCall
		return true
	case OpCom16:
		return rewriteValueWasm32_OpCom16(v)
	case OpCom32:
		return rewriteValueWasm32_OpCom32(v)
	case OpCom64:
		return rewriteValueWasm32_OpCom64(v)
	case OpCom8:
		return rewriteValueWasm32_OpCom8(v)
	case OpCondSelect:
		v.Op = OpWasm32Select
		return true
	case OpConst16:
		return rewriteValueWasm32_OpConst16(v)
	case OpConst32:
		return rewriteValueWasm32_OpConst32(v)
	case OpConst32F:
		v.Op = OpWasm32F32Const
		return true
	case OpConst64:
		v.Op = OpWasm32I64Const
		return true
	case OpConst64F:
		v.Op = OpWasm32F64Const
		return true
	case OpConst8:
		return rewriteValueWasm32_OpConst8(v)
	case OpConstBool:
		return rewriteValueWasm32_OpConstBool(v)
	case OpConstNil:
		return rewriteValueWasm32_OpConstNil(v)
	case OpConvert:
		v.Op = OpWasm32LoweredConvert
		return true
	case OpCopysign:
		v.Op = OpWasm32F64Copysign
		return true
	case OpCtz16:
		return rewriteValueWasm32_OpCtz16(v)
	case OpCtz16NonZero:
		v.Op = OpWasm32I64Ctz
		return true
	case OpCtz32:
		return rewriteValueWasm32_OpCtz32(v)
	case OpCtz32NonZero:
		v.Op = OpWasm32I64Ctz
		return true
	case OpCtz64:
		v.Op = OpWasm32I64Ctz
		return true
	case OpCtz64NonZero:
		v.Op = OpWasm32I64Ctz
		return true
	case OpCtz8:
		return rewriteValueWasm32_OpCtz8(v)
	case OpCtz8NonZero:
		v.Op = OpWasm32I64Ctz
		return true
	case OpCvt32Fto32:
		v.Op = OpWasm32I64TruncSatF32S
		return true
	case OpCvt32Fto32U:
		v.Op = OpWasm32I64TruncSatF32U
		return true
	case OpCvt32Fto64:
		v.Op = OpWasm32I64TruncSatF32S
		return true
	case OpCvt32Fto64F:
		v.Op = OpWasm32F64PromoteF32
		return true
	case OpCvt32Fto64U:
		v.Op = OpWasm32I64TruncSatF32U
		return true
	case OpCvt32Uto32F:
		return rewriteValueWasm32_OpCvt32Uto32F(v)
	case OpCvt32Uto64F:
		return rewriteValueWasm32_OpCvt32Uto64F(v)
	case OpCvt32to32F:
		return rewriteValueWasm32_OpCvt32to32F(v)
	case OpCvt32to64F:
		return rewriteValueWasm32_OpCvt32to64F(v)
	case OpCvt64Fto32:
		v.Op = OpWasm32I64TruncSatF64S
		return true
	case OpCvt64Fto32F:
		v.Op = OpWasm32F32DemoteF64
		return true
	case OpCvt64Fto32U:
		v.Op = OpWasm32I64TruncSatF64U
		return true
	case OpCvt64Fto64:
		v.Op = OpWasm32I64TruncSatF64S
		return true
	case OpCvt64Fto64U:
		v.Op = OpWasm32I64TruncSatF64U
		return true
	case OpCvt64Uto32F:
		v.Op = OpWasm32F32ConvertI64U
		return true
	case OpCvt64Uto64F:
		v.Op = OpWasm32F64ConvertI64U
		return true
	case OpCvt64to32F:
		v.Op = OpWasm32F32ConvertI64S
		return true
	case OpCvt64to64F:
		v.Op = OpWasm32F64ConvertI64S
		return true
	case OpCvtBoolToUint8:
		v.Op = OpCopy
		return true
	case OpDiv16:
		return rewriteValueWasm32_OpDiv16(v)
	case OpDiv16u:
		return rewriteValueWasm32_OpDiv16u(v)
	case OpDiv32:
		return rewriteValueWasm32_OpDiv32(v)
	case OpDiv32F:
		v.Op = OpWasm32F32Div
		return true
	case OpDiv32u:
		return rewriteValueWasm32_OpDiv32u(v)
	case OpDiv64:
		return rewriteValueWasm32_OpDiv64(v)
	case OpDiv64F:
		v.Op = OpWasm32F64Div
		return true
	case OpDiv64u:
		v.Op = OpWasm32I64DivU
		return true
	case OpDiv8:
		return rewriteValueWasm32_OpDiv8(v)
	case OpDiv8u:
		return rewriteValueWasm32_OpDiv8u(v)
	case OpEq16:
		return rewriteValueWasm32_OpEq16(v)
	case OpEq32:
		return rewriteValueWasm32_OpEq32(v)
	case OpEq32F:
		v.Op = OpWasm32F32Eq
		return true
	case OpEq64:
		v.Op = OpWasm32I64Eq
		return true
	case OpEq64F:
		v.Op = OpWasm32F64Eq
		return true
	case OpEq8:
		return rewriteValueWasm32_OpEq8(v)
	case OpEqB:
		v.Op = OpWasm32I64Eq
		return true
	case OpEqPtr:
		v.Op = OpWasm32I64Eq
		return true
	case OpFloor:
		v.Op = OpWasm32F64Floor
		return true
	case OpGetCallerPC:
		v.Op = OpWasm32LoweredGetCallerPC
		return true
	case OpGetCallerSP:
		v.Op = OpWasm32LoweredGetCallerSP
		return true
	case OpGetClosurePtr:
		v.Op = OpWasm32LoweredGetClosurePtr
		return true
	case OpInterCall:
		v.Op = OpWasm32LoweredInterCall
		return true
	case OpIsInBounds:
		v.Op = OpWasm32I64LtU
		return true
	case OpIsNonNil:
		return rewriteValueWasm32_OpIsNonNil(v)
	case OpIsSliceInBounds:
		v.Op = OpWasm32I64LeU
		return true
	case OpLeq16:
		return rewriteValueWasm32_OpLeq16(v)
	case OpLeq16U:
		return rewriteValueWasm32_OpLeq16U(v)
	case OpLeq32:
		return rewriteValueWasm32_OpLeq32(v)
	case OpLeq32F:
		v.Op = OpWasm32F32Le
		return true
	case OpLeq32U:
		return rewriteValueWasm32_OpLeq32U(v)
	case OpLeq64:
		v.Op = OpWasm32I64LeS
		return true
	case OpLeq64F:
		v.Op = OpWasm32F64Le
		return true
	case OpLeq64U:
		v.Op = OpWasm32I64LeU
		return true
	case OpLeq8:
		return rewriteValueWasm32_OpLeq8(v)
	case OpLeq8U:
		return rewriteValueWasm32_OpLeq8U(v)
	case OpLess16:
		return rewriteValueWasm32_OpLess16(v)
	case OpLess16U:
		return rewriteValueWasm32_OpLess16U(v)
	case OpLess32:
		return rewriteValueWasm32_OpLess32(v)
	case OpLess32F:
		v.Op = OpWasm32F32Lt
		return true
	case OpLess32U:
		return rewriteValueWasm32_OpLess32U(v)
	case OpLess64:
		v.Op = OpWasm32I64LtS
		return true
	case OpLess64F:
		v.Op = OpWasm32F64Lt
		return true
	case OpLess64U:
		v.Op = OpWasm32I64LtU
		return true
	case OpLess8:
		return rewriteValueWasm32_OpLess8(v)
	case OpLess8U:
		return rewriteValueWasm32_OpLess8U(v)
	case OpLoad:
		return rewriteValueWasm32_OpLoad(v)
	case OpLocalAddr:
		return rewriteValueWasm32_OpLocalAddr(v)
	case OpLsh16x16:
		return rewriteValueWasm32_OpLsh16x16(v)
	case OpLsh16x32:
		return rewriteValueWasm32_OpLsh16x32(v)
	case OpLsh16x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh16x8:
		return rewriteValueWasm32_OpLsh16x8(v)
	case OpLsh32x16:
		return rewriteValueWasm32_OpLsh32x16(v)
	case OpLsh32x32:
		return rewriteValueWasm32_OpLsh32x32(v)
	case OpLsh32x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh32x8:
		return rewriteValueWasm32_OpLsh32x8(v)
	case OpLsh64x16:
		return rewriteValueWasm32_OpLsh64x16(v)
	case OpLsh64x32:
		return rewriteValueWasm32_OpLsh64x32(v)
	case OpLsh64x64:
		return rewriteValueWasm32_OpLsh64x64(v)
	case OpLsh64x8:
		return rewriteValueWasm32_OpLsh64x8(v)
	case OpLsh8x16:
		return rewriteValueWasm32_OpLsh8x16(v)
	case OpLsh8x32:
		return rewriteValueWasm32_OpLsh8x32(v)
	case OpLsh8x64:
		v.Op = OpLsh64x64
		return true
	case OpLsh8x8:
		return rewriteValueWasm32_OpLsh8x8(v)
	case OpMod16:
		return rewriteValueWasm32_OpMod16(v)
	case OpMod16u:
		return rewriteValueWasm32_OpMod16u(v)
	case OpMod32:
		return rewriteValueWasm32_OpMod32(v)
	case OpMod32u:
		return rewriteValueWasm32_OpMod32u(v)
	case OpMod64:
		return rewriteValueWasm32_OpMod64(v)
	case OpMod64u:
		v.Op = OpWasm32I64RemU
		return true
	case OpMod8:
		return rewriteValueWasm32_OpMod8(v)
	case OpMod8u:
		return rewriteValueWasm32_OpMod8u(v)
	case OpMove:
		return rewriteValueWasm32_OpMove(v)
	case OpMul16:
		v.Op = OpWasm32I64Mul
		return true
	case OpMul32:
		v.Op = OpWasm32I64Mul
		return true
	case OpMul32F:
		v.Op = OpWasm32F32Mul
		return true
	case OpMul64:
		v.Op = OpWasm32I64Mul
		return true
	case OpMul64F:
		v.Op = OpWasm32F64Mul
		return true
	case OpMul8:
		v.Op = OpWasm32I64Mul
		return true
	case OpNeg16:
		return rewriteValueWasm32_OpNeg16(v)
	case OpNeg32:
		return rewriteValueWasm32_OpNeg32(v)
	case OpNeg32F:
		v.Op = OpWasm32F32Neg
		return true
	case OpNeg64:
		return rewriteValueWasm32_OpNeg64(v)
	case OpNeg64F:
		v.Op = OpWasm32F64Neg
		return true
	case OpNeg8:
		return rewriteValueWasm32_OpNeg8(v)
	case OpNeq16:
		return rewriteValueWasm32_OpNeq16(v)
	case OpNeq32:
		return rewriteValueWasm32_OpNeq32(v)
	case OpNeq32F:
		v.Op = OpWasm32F32Ne
		return true
	case OpNeq64:
		v.Op = OpWasm32I64Ne
		return true
	case OpNeq64F:
		v.Op = OpWasm32F64Ne
		return true
	case OpNeq8:
		return rewriteValueWasm32_OpNeq8(v)
	case OpNeqB:
		v.Op = OpWasm32I64Ne
		return true
	case OpNeqPtr:
		v.Op = OpWasm32I64Ne
		return true
	case OpNilCheck:
		v.Op = OpWasm32LoweredNilCheck
		return true
	case OpNot:
		v.Op = OpWasm32I64Eqz
		return true
	case OpOffPtr:
		v.Op = OpWasm32I64AddConst
		return true
	case OpOr16:
		v.Op = OpWasm32I64Or
		return true
	case OpOr32:
		v.Op = OpWasm32I64Or
		return true
	case OpOr64:
		v.Op = OpWasm32I64Or
		return true
	case OpOr8:
		v.Op = OpWasm32I64Or
		return true
	case OpOrB:
		v.Op = OpWasm32I64Or
		return true
	case OpPopCount16:
		return rewriteValueWasm32_OpPopCount16(v)
	case OpPopCount32:
		return rewriteValueWasm32_OpPopCount32(v)
	case OpPopCount64:
		v.Op = OpWasm32I64Popcnt
		return true
	case OpPopCount8:
		return rewriteValueWasm32_OpPopCount8(v)
	case OpRotateLeft16:
		return rewriteValueWasm32_OpRotateLeft16(v)
	case OpRotateLeft32:
		v.Op = OpWasm32I32Rotl
		return true
	case OpRotateLeft64:
		v.Op = OpWasm32I64Rotl
		return true
	case OpRotateLeft8:
		return rewriteValueWasm32_OpRotateLeft8(v)
	case OpRound32F:
		v.Op = OpCopy
		return true
	case OpRound64F:
		v.Op = OpCopy
		return true
	case OpRoundToEven:
		v.Op = OpWasm32F64Nearest
		return true
	case OpRsh16Ux16:
		return rewriteValueWasm32_OpRsh16Ux16(v)
	case OpRsh16Ux32:
		return rewriteValueWasm32_OpRsh16Ux32(v)
	case OpRsh16Ux64:
		return rewriteValueWasm32_OpRsh16Ux64(v)
	case OpRsh16Ux8:
		return rewriteValueWasm32_OpRsh16Ux8(v)
	case OpRsh16x16:
		return rewriteValueWasm32_OpRsh16x16(v)
	case OpRsh16x32:
		return rewriteValueWasm32_OpRsh16x32(v)
	case OpRsh16x64:
		return rewriteValueWasm32_OpRsh16x64(v)
	case OpRsh16x8:
		return rewriteValueWasm32_OpRsh16x8(v)
	case OpRsh32Ux16:
		return rewriteValueWasm32_OpRsh32Ux16(v)
	case OpRsh32Ux32:
		return rewriteValueWasm32_OpRsh32Ux32(v)
	case OpRsh32Ux64:
		return rewriteValueWasm32_OpRsh32Ux64(v)
	case OpRsh32Ux8:
		return rewriteValueWasm32_OpRsh32Ux8(v)
	case OpRsh32x16:
		return rewriteValueWasm32_OpRsh32x16(v)
	case OpRsh32x32:
		return rewriteValueWasm32_OpRsh32x32(v)
	case OpRsh32x64:
		return rewriteValueWasm32_OpRsh32x64(v)
	case OpRsh32x8:
		return rewriteValueWasm32_OpRsh32x8(v)
	case OpRsh64Ux16:
		return rewriteValueWasm32_OpRsh64Ux16(v)
	case OpRsh64Ux32:
		return rewriteValueWasm32_OpRsh64Ux32(v)
	case OpRsh64Ux64:
		return rewriteValueWasm32_OpRsh64Ux64(v)
	case OpRsh64Ux8:
		return rewriteValueWasm32_OpRsh64Ux8(v)
	case OpRsh64x16:
		return rewriteValueWasm32_OpRsh64x16(v)
	case OpRsh64x32:
		return rewriteValueWasm32_OpRsh64x32(v)
	case OpRsh64x64:
		return rewriteValueWasm32_OpRsh64x64(v)
	case OpRsh64x8:
		return rewriteValueWasm32_OpRsh64x8(v)
	case OpRsh8Ux16:
		return rewriteValueWasm32_OpRsh8Ux16(v)
	case OpRsh8Ux32:
		return rewriteValueWasm32_OpRsh8Ux32(v)
	case OpRsh8Ux64:
		return rewriteValueWasm32_OpRsh8Ux64(v)
	case OpRsh8Ux8:
		return rewriteValueWasm32_OpRsh8Ux8(v)
	case OpRsh8x16:
		return rewriteValueWasm32_OpRsh8x16(v)
	case OpRsh8x32:
		return rewriteValueWasm32_OpRsh8x32(v)
	case OpRsh8x64:
		return rewriteValueWasm32_OpRsh8x64(v)
	case OpRsh8x8:
		return rewriteValueWasm32_OpRsh8x8(v)
	case OpSignExt16to32:
		return rewriteValueWasm32_OpSignExt16to32(v)
	case OpSignExt16to64:
		return rewriteValueWasm32_OpSignExt16to64(v)
	case OpSignExt32to64:
		return rewriteValueWasm32_OpSignExt32to64(v)
	case OpSignExt8to16:
		return rewriteValueWasm32_OpSignExt8to16(v)
	case OpSignExt8to32:
		return rewriteValueWasm32_OpSignExt8to32(v)
	case OpSignExt8to64:
		return rewriteValueWasm32_OpSignExt8to64(v)
	case OpSlicemask:
		return rewriteValueWasm32_OpSlicemask(v)
	case OpSqrt:
		v.Op = OpWasm32F64Sqrt
		return true
	case OpSqrt32:
		v.Op = OpWasm32F32Sqrt
		return true
	case OpStaticCall:
		v.Op = OpWasm32LoweredStaticCall
		return true
	case OpStore:
		return rewriteValueWasm32_OpStore(v)
	case OpSub16:
		v.Op = OpWasm32I64Sub
		return true
	case OpSub32:
		v.Op = OpWasm32I64Sub
		return true
	case OpSub32F:
		v.Op = OpWasm32F32Sub
		return true
	case OpSub64:
		v.Op = OpWasm32I64Sub
		return true
	case OpSub64F:
		v.Op = OpWasm32F64Sub
		return true
	case OpSub8:
		v.Op = OpWasm32I64Sub
		return true
	case OpSubPtr:
		v.Op = OpWasm32I64Sub
		return true
	case OpTailCall:
		v.Op = OpWasm32LoweredTailCall
		return true
	case OpTrunc:
		v.Op = OpWasm32F64Trunc
		return true
	case OpTrunc16to8:
		v.Op = OpCopy
		return true
	case OpTrunc32to16:
		v.Op = OpCopy
		return true
	case OpTrunc32to8:
		v.Op = OpCopy
		return true
	case OpTrunc64to16:
		v.Op = OpCopy
		return true
	case OpTrunc64to32:
		v.Op = OpCopy
		return true
	case OpTrunc64to8:
		v.Op = OpCopy
		return true
	case OpWB:
		v.Op = OpWasm32LoweredWB
		return true
	case OpWasm32F64Add:
		return rewriteValueWasm32_OpWasm32F64Add(v)
	case OpWasm32F64Mul:
		return rewriteValueWasm32_OpWasm32F64Mul(v)
	case OpWasm32I64Add:
		return rewriteValueWasm32_OpWasm32I64Add(v)
	case OpWasm32I64AddConst:
		return rewriteValueWasm32_OpWasm32I64AddConst(v)
	case OpWasm32I64And:
		return rewriteValueWasm32_OpWasm32I64And(v)
	case OpWasm32I64Eq:
		return rewriteValueWasm32_OpWasm32I64Eq(v)
	case OpWasm32I64Eqz:
		return rewriteValueWasm32_OpWasm32I64Eqz(v)
	case OpWasm32I64LeU:
		return rewriteValueWasm32_OpWasm32I64LeU(v)
	case OpWasm32I64Load:
		return rewriteValueWasm32_OpWasm32I64Load(v)
	case OpWasm32I64Load16S:
		return rewriteValueWasm32_OpWasm32I64Load16S(v)
	case OpWasm32I64Load16U:
		return rewriteValueWasm32_OpWasm32I64Load16U(v)
	case OpWasm32I64Load32S:
		return rewriteValueWasm32_OpWasm32I64Load32S(v)
	case OpWasm32I64Load32U:
		return rewriteValueWasm32_OpWasm32I64Load32U(v)
	case OpWasm32I64Load8S:
		return rewriteValueWasm32_OpWasm32I64Load8S(v)
	case OpWasm32I64Load8U:
		return rewriteValueWasm32_OpWasm32I64Load8U(v)
	case OpWasm32I64LtU:
		return rewriteValueWasm32_OpWasm32I64LtU(v)
	case OpWasm32I64Mul:
		return rewriteValueWasm32_OpWasm32I64Mul(v)
	case OpWasm32I64Ne:
		return rewriteValueWasm32_OpWasm32I64Ne(v)
	case OpWasm32I64Or:
		return rewriteValueWasm32_OpWasm32I64Or(v)
	case OpWasm32I64Shl:
		return rewriteValueWasm32_OpWasm32I64Shl(v)
	case OpWasm32I64ShrS:
		return rewriteValueWasm32_OpWasm32I64ShrS(v)
	case OpWasm32I64ShrU:
		return rewriteValueWasm32_OpWasm32I64ShrU(v)
	case OpWasm32I64Store:
		return rewriteValueWasm32_OpWasm32I64Store(v)
	case OpWasm32I64Store16:
		return rewriteValueWasm32_OpWasm32I64Store16(v)
	case OpWasm32I64Store32:
		return rewriteValueWasm32_OpWasm32I64Store32(v)
	case OpWasm32I64Store8:
		return rewriteValueWasm32_OpWasm32I64Store8(v)
	case OpWasm32I64Xor:
		return rewriteValueWasm32_OpWasm32I64Xor(v)
	case OpXor16:
		v.Op = OpWasm32I64Xor
		return true
	case OpXor32:
		v.Op = OpWasm32I64Xor
		return true
	case OpXor64:
		v.Op = OpWasm32I64Xor
		return true
	case OpXor8:
		v.Op = OpWasm32I64Xor
		return true
	case OpZero:
		return rewriteValueWasm32_OpZero(v)
	case OpZeroExt16to32:
		return rewriteValueWasm32_OpZeroExt16to32(v)
	case OpZeroExt16to64:
		return rewriteValueWasm32_OpZeroExt16to64(v)
	case OpZeroExt32to64:
		return rewriteValueWasm32_OpZeroExt32to64(v)
	case OpZeroExt8to16:
		return rewriteValueWasm32_OpZeroExt8to16(v)
	case OpZeroExt8to32:
		return rewriteValueWasm32_OpZeroExt8to32(v)
	case OpZeroExt8to64:
		return rewriteValueWasm32_OpZeroExt8to64(v)
	}
	return false
}
func rewriteValueWasm32_OpAddr(v *Value) bool {
	v_0 := v.Args[0]
	// match: (Addr {sym} base)
	// result: (LoweredAddr {sym} [0] base)
	for {
		sym := auxToSym(v.Aux)
		base := v_0
		v.reset(OpWasm32LoweredAddr)
		v.AuxInt = int32ToAuxInt(0)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
}
func rewriteValueWasm32_OpBitLen64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (BitLen64 x)
	// result: (I64Sub (I64Const [64]) (I64Clz x))
	for {
		x := v_0
		v.reset(OpWasm32I64Sub)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Clz, typ.Int64)
		v1.AddArg(x)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpCom16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com16 x)
	// result: (I64Xor x (I64Const [-1]))
	for {
		x := v_0
		v.reset(OpWasm32I64Xor)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpCom32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com32 x)
	// result: (I64Xor x (I64Const [-1]))
	for {
		x := v_0
		v.reset(OpWasm32I64Xor)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpCom64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com64 x)
	// result: (I64Xor x (I64Const [-1]))
	for {
		x := v_0
		v.reset(OpWasm32I64Xor)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpCom8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Com8 x)
	// result: (I64Xor x (I64Const [-1]))
	for {
		x := v_0
		v.reset(OpWasm32I64Xor)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(-1)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpConst16(v *Value) bool {
	// match: (Const16 [c])
	// result: (I64Const [int64(c)])
	for {
		c := auxIntToInt16(v.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(c))
		return true
	}
}
func rewriteValueWasm32_OpConst32(v *Value) bool {
	// match: (Const32 [c])
	// result: (I64Const [int64(c)])
	for {
		c := auxIntToInt32(v.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(c))
		return true
	}
}
func rewriteValueWasm32_OpConst8(v *Value) bool {
	// match: (Const8 [c])
	// result: (I64Const [int64(c)])
	for {
		c := auxIntToInt8(v.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(c))
		return true
	}
}
func rewriteValueWasm32_OpConstBool(v *Value) bool {
	// match: (ConstBool [c])
	// result: (I64Const [b2i(c)])
	for {
		c := auxIntToBool(v.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(b2i(c))
		return true
	}
}
func rewriteValueWasm32_OpConstNil(v *Value) bool {
	// match: (ConstNil)
	// result: (I64Const [0])
	for {
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
}
func rewriteValueWasm32_OpCtz16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz16 x)
	// result: (I64Ctz (I64Or x (I64Const [0x10000])))
	for {
		x := v_0
		v.reset(OpWasm32I64Ctz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Or, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0x10000)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCtz32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz32 x)
	// result: (I64Ctz (I64Or x (I64Const [0x100000000])))
	for {
		x := v_0
		v.reset(OpWasm32I64Ctz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Or, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0x100000000)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCtz8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Ctz8 x)
	// result: (I64Ctz (I64Or x (I64Const [0x100])))
	for {
		x := v_0
		v.reset(OpWasm32I64Ctz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Or, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0x100)
		v0.AddArg2(x, v1)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCvt32Uto32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Uto32F x)
	// result: (F32ConvertI64U (ZeroExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasm32F32ConvertI64U)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCvt32Uto64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32Uto64F x)
	// result: (F64ConvertI64U (ZeroExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasm32F64ConvertI64U)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCvt32to32F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to32F x)
	// result: (F32ConvertI64S (SignExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasm32F32ConvertI64S)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpCvt32to64F(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Cvt32to64F x)
	// result: (F64ConvertI64S (SignExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasm32F64ConvertI64S)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpDiv16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16 [false] x y)
	// result: (I64DivS (SignExt16to64 x) (SignExt16to64 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm32_OpDiv16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div16u x y)
	// result: (I64DivU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpDiv32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32 [false] x y)
	// result: (I64DivS (SignExt32to64 x) (SignExt32to64 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm32_OpDiv32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div32u x y)
	// result: (I64DivU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpDiv64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Div64 [false] x y)
	// result: (I64DivS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm32_OpDiv8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8 x y)
	// result: (I64DivS (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpDiv8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Div8u x y)
	// result: (I64DivU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64DivU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpEq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq16 x y)
	// result: (I64Eq (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpEq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq32 x y)
	// result: (I64Eq (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpEq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Eq8 x y)
	// result: (I64Eq (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Eq)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpIsNonNil(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (IsNonNil p)
	// result: (I64Eqz (I64Eqz p))
	for {
		p := v_0
		v.reset(OpWasm32I64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Eqz, typ.Bool)
		v0.AddArg(p)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpLeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16 x y)
	// result: (I64LeS (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLeq16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq16U x y)
	// result: (I64LeU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32 x y)
	// result: (I64LeS (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLeq32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq32U x y)
	// result: (I64LeU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8 x y)
	// result: (I64LeS (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLeq8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Leq8U x y)
	// result: (I64LeU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LeU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16 x y)
	// result: (I64LtS (SignExt16to64 x) (SignExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less16U x y)
	// result: (I64LtU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32 x y)
	// result: (I64LtS (SignExt32to64 x) (SignExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less32U x y)
	// result: (I64LtU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8 x y)
	// result: (I64LtS (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLess8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Less8U x y)
	// result: (I64LtU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64LtU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpLoad(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Load <t> ptr mem)
	// cond: is32BitFloat(t)
	// result: (F32Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpWasm32F32Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: is64BitFloat(t)
	// result: (F64Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpWasm32F64Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 8
	// result: (I64Load ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 8) {
			break
		}
		v.reset(OpWasm32I64Load)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 4 && !t.IsSigned()
	// result: (I64Load32U ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 4 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load32U)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 4 && t.IsSigned()
	// result: (I64Load32S ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 4 && t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load32S)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && !t.IsSigned()
	// result: (I64Load16U ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 2 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load16U)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 2 && t.IsSigned()
	// result: (I64Load16S ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 2 && t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load16S)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && !t.IsSigned()
	// result: (I64Load8U ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 1 && !t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load8U)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.Size() == 1 && t.IsSigned()
	// result: (I64Load8S ptr mem)
	for {
		t := v.Type
		ptr := v_0
		mem := v_1
		if !(t.Size() == 1 && t.IsSigned()) {
			break
		}
		v.reset(OpWasm32I64Load8S)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpLocalAddr(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (LocalAddr <t> {sym} base mem)
	// cond: t.Elem().HasPointers()
	// result: (LoweredAddr {sym} (SPanchored base mem))
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		mem := v_1
		if !(t.Elem().HasPointers()) {
			break
		}
		v.reset(OpWasm32LoweredAddr)
		v.Aux = symToAux(sym)
		v0 := b.NewValue0(v.Pos, OpSPanchored, typ.Uintptr)
		v0.AddArg2(base, mem)
		v.AddArg(v0)
		return true
	}
	// match: (LocalAddr <t> {sym} base _)
	// cond: !t.Elem().HasPointers()
	// result: (LoweredAddr {sym} base)
	for {
		t := v.Type
		sym := auxToSym(v.Aux)
		base := v_0
		if !(!t.Elem().HasPointers()) {
			break
		}
		v.reset(OpWasm32LoweredAddr)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	return false
}
func rewriteValueWasm32_OpLsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x16 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x32 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh16x8 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x16 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x32 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh32x8 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x16 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x32 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64Shl x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasm32I64Shl)
		v.AddArg2(x, y)
		return true
	}
	// match: (Lsh64x64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64Shl x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasm32I64Shl)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Lsh64x64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64Const [0])
	for {
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Lsh64x64 x y)
	// result: (Select (I64Shl x y) (I64Const [0]) (I64LtU y (I64Const [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32Select)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpWasm32I64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueWasm32_OpLsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh64x8 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x16 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x32 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpLsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Lsh8x8 [c] x y)
	// result: (Lsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpLsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpMod16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16 [false] x y)
	// result: (I64RemS (SignExt16to64 x) (SignExt16to64 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm32_OpMod16u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod16u x y)
	// result: (I64RemU (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpMod32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32 [false] x y)
	// result: (I64RemS (SignExt32to64 x) (SignExt32to64 y))
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
	return false
}
func rewriteValueWasm32_OpMod32u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod32u x y)
	// result: (I64RemU (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpMod64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Mod64 [false] x y)
	// result: (I64RemS x y)
	for {
		if auxIntToBool(v.AuxInt) != false {
			break
		}
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemS)
		v.AddArg2(x, y)
		return true
	}
	return false
}
func rewriteValueWasm32_OpMod8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8 x y)
	// result: (I64RemS (SignExt8to64 x) (SignExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemS)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpMod8u(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Mod8u x y)
	// result: (I64RemU (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64RemU)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpMove(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Move [0] _ _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_2
		v.copyOf(mem)
		return true
	}
	// match: (Move [1] dst src mem)
	// result: (I64Store8 dst (I64Load8U src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store8)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load8U, typ.UInt8)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [2] dst src mem)
	// result: (I64Store16 dst (I64Load16U src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store16)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load16U, typ.UInt16)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [4] dst src mem)
	// result: (I64Store32 dst (I64Load32U src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store32)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load32U, typ.UInt32)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [8] dst src mem)
	// result: (I64Store dst (I64Load src mem) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load, typ.UInt64)
		v0.AddArg2(src, mem)
		v.AddArg3(dst, v0, mem)
		return true
	}
	// match: (Move [16] dst src mem)
	// result: (I64Store [8] dst (I64Load [8] src mem) (I64Store dst (I64Load src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(8)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load, typ.UInt64)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [3] dst src mem)
	// result: (I64Store8 [2] dst (I64Load8U [2] src mem) (I64Store16 dst (I64Load16U src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store8)
		v.AuxInt = int64ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load8U, typ.UInt8)
		v0.AuxInt = int64ToAuxInt(2)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store16, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load16U, typ.UInt16)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [5] dst src mem)
	// result: (I64Store8 [4] dst (I64Load8U [4] src mem) (I64Store32 dst (I64Load32U src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 5 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store8)
		v.AuxInt = int64ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load8U, typ.UInt8)
		v0.AuxInt = int64ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load32U, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [6] dst src mem)
	// result: (I64Store16 [4] dst (I64Load16U [4] src mem) (I64Store32 dst (I64Load32U src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store16)
		v.AuxInt = int64ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load16U, typ.UInt16)
		v0.AuxInt = int64ToAuxInt(4)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load32U, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [7] dst src mem)
	// result: (I64Store32 [3] dst (I64Load32U [3] src mem) (I64Store32 dst (I64Load32U src mem) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 7 {
			break
		}
		dst := v_0
		src := v_1
		mem := v_2
		v.reset(OpWasm32I64Store32)
		v.AuxInt = int64ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load32U, typ.UInt32)
		v0.AuxInt = int64ToAuxInt(3)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load32U, typ.UInt32)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: s > 8 && s < 16
	// result: (I64Store [s-8] dst (I64Load [s-8] src mem) (I64Store dst (I64Load src mem) mem))
	for {
		s := auxIntToInt64(v.AuxInt)
		dst := v_0
		src := v_1
		mem := v_2
		if !(s > 8 && s < 16) {
			break
		}
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(s - 8)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Load, typ.UInt64)
		v0.AuxInt = int64ToAuxInt(s - 8)
		v0.AddArg2(src, mem)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Load, typ.UInt64)
		v2.AddArg2(src, mem)
		v1.AddArg3(dst, v2, mem)
		v.AddArg3(dst, v0, v1)
		return true
	}
	// match: (Move [s] dst src mem)
	// cond: logLargeCopy(v, s)
	// result: (LoweredMove [s] dst src mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		dst := v_0
		src := v_1
		mem := v_2
		if !(logLargeCopy(v, s)) {
			break
		}
		v.reset(OpWasm32LoweredMove)
		v.AuxInt = int64ToAuxInt(s)
		v.AddArg3(dst, src, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpNeg16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg16 x)
	// result: (I64Sub (I64Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasm32I64Sub)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm32_OpNeg32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg32 x)
	// result: (I64Sub (I64Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasm32I64Sub)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm32_OpNeg64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg64 x)
	// result: (I64Sub (I64Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasm32I64Sub)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm32_OpNeg8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neg8 x)
	// result: (I64Sub (I64Const [0]) x)
	for {
		x := v_0
		v.reset(OpWasm32I64Sub)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg2(v0, x)
		return true
	}
}
func rewriteValueWasm32_OpNeq16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq16 x y)
	// result: (I64Ne (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpNeq32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq32 x y)
	// result: (I64Ne (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpNeq8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Neq8 x y)
	// result: (I64Ne (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64Ne)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpPopCount16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount16 x)
	// result: (I64Popcnt (ZeroExt16to64 x))
	for {
		x := v_0
		v.reset(OpWasm32I64Popcnt)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpPopCount32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount32 x)
	// result: (I64Popcnt (ZeroExt32to64 x))
	for {
		x := v_0
		v.reset(OpWasm32I64Popcnt)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpPopCount8(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (PopCount8 x)
	// result: (I64Popcnt (ZeroExt8to64 x))
	for {
		x := v_0
		v.reset(OpWasm32I64Popcnt)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
}
func rewriteValueWasm32_OpRotateLeft16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft16 <t> x (I64Const [c]))
	// result: (Or16 (Lsh16x64 <t> x (I64Const [c&15])) (Rsh16Ux64 <t> x (I64Const [-c&15])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr16)
		v0 := b.NewValue0(v.Pos, OpLsh16x64, t)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(c & 15)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh16Ux64, t)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(-c & 15)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueWasm32_OpRotateLeft8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (RotateLeft8 <t> x (I64Const [c]))
	// result: (Or8 (Lsh8x64 <t> x (I64Const [c&7])) (Rsh8Ux64 <t> x (I64Const [-c&7])))
	for {
		t := v.Type
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		v.reset(OpOr8)
		v0 := b.NewValue0(v.Pos, OpLsh8x64, t)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(c & 7)
		v0.AddArg2(x, v1)
		v2 := b.NewValue0(v.Pos, OpRsh8Ux64, t)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(-c & 7)
		v2.AddArg2(x, v3)
		v.AddArg2(v0, v2)
		return true
	}
	return false
}
func rewriteValueWasm32_OpRsh16Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux16 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt16to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh16Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux32 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt16to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh16Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt16to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh16Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16Ux8 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt16to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh16x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x16 [c] x y)
	// result: (Rsh64x64 [c] (SignExt16to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh16x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x32 [c] x y)
	// result: (Rsh64x64 [c] (SignExt16to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh16x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt16to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh16x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh16x8 [c] x y)
	// result: (Rsh64x64 [c] (SignExt16to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt16to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux16 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt32to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux32 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt32to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt32to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh32Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32Ux8 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt32to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x16 [c] x y)
	// result: (Rsh64x64 [c] (SignExt32to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x32 [c] x y)
	// result: (Rsh64x64 [c] (SignExt32to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh32x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt32to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh32x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh32x8 [c] x y)
	// result: (Rsh64x64 [c] (SignExt32to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt32to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh64Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux16 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux32 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64ShrU x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasm32I64ShrU)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh64Ux64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64ShrU x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasm32I64ShrU)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64Ux64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64Const [0])
	for {
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (Rsh64Ux64 x y)
	// result: (Select (I64ShrU x y) (I64Const [0]) (I64LtU y (I64Const [64])))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32Select)
		v0 := b.NewValue0(v.Pos, OpWasm32I64ShrU, typ.Int64)
		v0.AddArg2(x, y)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0)
		v2 := b.NewValue0(v.Pos, OpWasm32I64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v.AddArg3(v0, v1, v2)
		return true
	}
}
func rewriteValueWasm32_OpRsh64Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64Ux8 [c] x y)
	// result: (Rsh64Ux64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x16 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x32 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x64 x y)
	// cond: shiftIsBounded(v)
	// result: (I64ShrS x y)
	for {
		x := v_0
		y := v_1
		if !(shiftIsBounded(v)) {
			break
		}
		v.reset(OpWasm32I64ShrS)
		v.AddArg2(x, y)
		return true
	}
	// match: (Rsh64x64 x (I64Const [c]))
	// cond: uint64(c) < 64
	// result: (I64ShrS x (I64Const [c]))
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) < 64) {
			break
		}
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(c)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x64 x (I64Const [c]))
	// cond: uint64(c) >= 64
	// result: (I64ShrS x (I64Const [63]))
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		c := auxIntToInt64(v_1.AuxInt)
		if !(uint64(c) >= 64) {
			break
		}
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(63)
		v.AddArg2(x, v0)
		return true
	}
	// match: (Rsh64x64 x y)
	// result: (I64ShrS x (Select <typ.Int64> y (I64Const [63]) (I64LtU y (I64Const [64]))))
	for {
		x := v_0
		y := v_1
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32Select, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(63)
		v2 := b.NewValue0(v.Pos, OpWasm32I64LtU, typ.Bool)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v3.AuxInt = int64ToAuxInt(64)
		v2.AddArg2(y, v3)
		v0.AddArg3(y, v1, v2)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh64x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh64x8 [c] x y)
	// result: (Rsh64x64 [c] x (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(y)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpRsh8Ux16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux16 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt8to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh8Ux32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux32 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt8to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh8Ux64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux64 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt8to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh8Ux8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8Ux8 [c] x y)
	// result: (Rsh64Ux64 [c] (ZeroExt8to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64Ux64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh8x16(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x16 [c] x y)
	// result: (Rsh64x64 [c] (SignExt8to64 x) (ZeroExt16to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt16to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh8x32(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x32 [c] x y)
	// result: (Rsh64x64 [c] (SignExt8to64 x) (ZeroExt32to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt32to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpRsh8x64(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x64 [c] x y)
	// result: (Rsh64x64 [c] (SignExt8to64 x) y)
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v.AddArg2(v0, y)
		return true
	}
}
func rewriteValueWasm32_OpRsh8x8(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Rsh8x8 [c] x y)
	// result: (Rsh64x64 [c] (SignExt8to64 x) (ZeroExt8to64 y))
	for {
		c := auxIntToBool(v.AuxInt)
		x := v_0
		y := v_1
		v.reset(OpRsh64x64)
		v.AuxInt = boolToAuxInt(c)
		v0 := b.NewValue0(v.Pos, OpSignExt8to64, typ.Int64)
		v0.AddArg(x)
		v1 := b.NewValue0(v.Pos, OpZeroExt8to64, typ.UInt64)
		v1.AddArg(y)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt16to32 x:(I64Load16S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load16S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt16to32 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend16S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend16S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt16to32 x)
	// result: (I64ShrS (I64Shl x (I64Const [48])) (I64Const [48]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(48)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt16to64 x:(I64Load16S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load16S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt16to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend16S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend16S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt16to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [48])) (I64Const [48]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(48)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt32to64 x:(I64Load32S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load32S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt32to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend32S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend32S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt32to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [32])) (I64Const [32]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(32)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to16 x:(I64Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to16 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to16 x)
	// result: (I64ShrS (I64Shl x (I64Const [56])) (I64Const [56]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(56)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to32 x:(I64Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to32 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to32 x)
	// result: (I64ShrS (I64Shl x (I64Const [56])) (I64Const [56]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(56)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSignExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (SignExt8to64 x:(I64Load8S _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8S {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (SignExt8to64 x)
	// cond: buildcfg.GOWASM.SignExt
	// result: (I64Extend8S x)
	for {
		x := v_0
		if !(buildcfg.GOWASM.SignExt) {
			break
		}
		v.reset(OpWasm32I64Extend8S)
		v.AddArg(x)
		return true
	}
	// match: (SignExt8to64 x)
	// result: (I64ShrS (I64Shl x (I64Const [56])) (I64Const [56]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Shl, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(56)
		v0.AddArg2(x, v1)
		v.AddArg2(v0, v1)
		return true
	}
}
func rewriteValueWasm32_OpSlicemask(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Slicemask x)
	// result: (I64ShrS (I64Sub (I64Const [0]) x) (I64Const [63]))
	for {
		x := v_0
		v.reset(OpWasm32I64ShrS)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Sub, typ.Int64)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v1.AuxInt = int64ToAuxInt(0)
		v0.AddArg2(v1, x)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v2.AuxInt = int64ToAuxInt(63)
		v.AddArg2(v0, v2)
		return true
	}
}
func rewriteValueWasm32_OpStore(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (Store {t} ptr val mem)
	// cond: is64BitFloat(t)
	// result: (F64Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(is64BitFloat(t)) {
			break
		}
		v.reset(OpWasm32F64Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: is32BitFloat(t)
	// result: (F32Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(is32BitFloat(t)) {
			break
		}
		v.reset(OpWasm32F32Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 8
	// result: (I64Store ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 8) {
			break
		}
		v.reset(OpWasm32I64Store)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 4
	// result: (I64Store32 ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 4) {
			break
		}
		v.reset(OpWasm32I64Store32)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 2
	// result: (I64Store16 ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 2) {
			break
		}
		v.reset(OpWasm32I64Store16)
		v.AddArg3(ptr, val, mem)
		return true
	}
	// match: (Store {t} ptr val mem)
	// cond: t.Size() == 1
	// result: (I64Store8 ptr val mem)
	for {
		t := auxToType(v.Aux)
		ptr := v_0
		val := v_1
		mem := v_2
		if !(t.Size() == 1) {
			break
		}
		v.reset(OpWasm32I64Store8)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32F64Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F64Add (F64Const [x]) (F64Const [y]))
	// result: (F64Const [x + y])
	for {
		if v_0.Op != OpWasm32F64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		if v_1.Op != OpWasm32F64Const {
			break
		}
		y := auxIntToFloat64(v_1.AuxInt)
		v.reset(OpWasm32F64Const)
		v.AuxInt = float64ToAuxInt(x + y)
		return true
	}
	// match: (F64Add (F64Const [x]) y)
	// cond: y.Op != OpWasm32F64Const
	// result: (F64Add y (F64Const [x]))
	for {
		if v_0.Op != OpWasm32F64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32F64Const) {
			break
		}
		v.reset(OpWasm32F64Add)
		v0 := b.NewValue0(v.Pos, OpWasm32F64Const, typ.Float64)
		v0.AuxInt = float64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32F64Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (F64Mul (F64Const [x]) (F64Const [y]))
	// cond: !math.IsNaN(x * y)
	// result: (F64Const [x * y])
	for {
		if v_0.Op != OpWasm32F64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		if v_1.Op != OpWasm32F64Const {
			break
		}
		y := auxIntToFloat64(v_1.AuxInt)
		if !(!math.IsNaN(x * y)) {
			break
		}
		v.reset(OpWasm32F64Const)
		v.AuxInt = float64ToAuxInt(x * y)
		return true
	}
	// match: (F64Mul (F64Const [x]) y)
	// cond: y.Op != OpWasm32F64Const
	// result: (F64Mul y (F64Const [x]))
	for {
		if v_0.Op != OpWasm32F64Const {
			break
		}
		x := auxIntToFloat64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32F64Const) {
			break
		}
		v.reset(OpWasm32F64Mul)
		v0 := b.NewValue0(v.Pos, OpWasm32F64Const, typ.Float64)
		v0.AuxInt = float64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Add(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Add (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x + y])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x + y)
		return true
	}
	// match: (I64Add (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Add y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Add)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I64Add x (I64Const <t> [y]))
	// cond: !t.IsPtr()
	// result: (I64AddConst [y] x)
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const {
			break
		}
		t := v_1.Type
		y := auxIntToInt64(v_1.AuxInt)
		if !(!t.IsPtr()) {
			break
		}
		v.reset(OpWasm32I64AddConst)
		v.AuxInt = int64ToAuxInt(y)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64AddConst(v *Value) bool {
	v_0 := v.Args[0]
	// match: (I64AddConst [0] x)
	// result: x
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		x := v_0
		v.copyOf(x)
		return true
	}
	// match: (I64AddConst [off] (LoweredAddr {sym} [off2] base))
	// cond: isU32Bit(off+int64(off2))
	// result: (LoweredAddr {sym} [int32(off)+off2] base)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32LoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		base := v_0.Args[0]
		if !(isU32Bit(off + int64(off2))) {
			break
		}
		v.reset(OpWasm32LoweredAddr)
		v.AuxInt = int32ToAuxInt(int32(off) + off2)
		v.Aux = symToAux(sym)
		v.AddArg(base)
		return true
	}
	// match: (I64AddConst [off] x:(SP))
	// cond: isU32Bit(off)
	// result: (LoweredAddr [int32(off)] x)
	for {
		off := auxIntToInt64(v.AuxInt)
		x := v_0
		if x.Op != OpSP || !(isU32Bit(off)) {
			break
		}
		v.reset(OpWasm32LoweredAddr)
		v.AuxInt = int32ToAuxInt(int32(off))
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64And(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64And (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x & y])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x & y)
		return true
	}
	// match: (I64And (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64And y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Eq(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Eq (I64Const [x]) (I64Const [y]))
	// cond: x == y
	// result: (I64Const [1])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(1)
		return true
	}
	// match: (I64Eq (I64Const [x]) (I64Const [y]))
	// cond: x != y
	// result: (I64Const [0])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (I64Eq (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Eq y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Eq)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I64Eq x (I64Const [0]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasm32I64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Eqz(v *Value) bool {
	v_0 := v.Args[0]
	// match: (I64Eqz (I64Eqz (I64Eqz x)))
	// result: (I64Eqz x)
	for {
		if v_0.Op != OpWasm32I64Eqz {
			break
		}
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpWasm32I64Eqz {
			break
		}
		x := v_0_0.Args[0]
		v.reset(OpWasm32I64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64LeU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64LeU x (I64Const [0]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasm32I64Eqz)
		v.AddArg(x)
		return true
	}
	// match: (I64LeU (I64Const [1]) x)
	// result: (I64Eqz (I64Eqz x))
	for {
		if v_0.Op != OpWasm32I64Const || auxIntToInt64(v_0.AuxInt) != 1 {
			break
		}
		x := v_1
		v.reset(OpWasm32I64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (I64Load [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I64Load [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(off+int64(off2))
	// result: (I64Const [int64(read64(sym, off+int64(off2), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32LoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(off+int64(off2))) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(read64(sym, off+int64(off2), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load16S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load16S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load16S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load16S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load16U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (I64Load16U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load16U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load16U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I64Load16U [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(off+int64(off2))
	// result: (I64Const [int64(read16(sym, off+int64(off2), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32LoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(off+int64(off2))) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(read16(sym, off+int64(off2), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load32S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load32S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load32S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load32S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load32U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	config := b.Func.Config
	// match: (I64Load32U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load32U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load32U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I64Load32U [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(off+int64(off2))
	// result: (I64Const [int64(read32(sym, off+int64(off2), config.ctxt.Arch.ByteOrder))])
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32LoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(off+int64(off2))) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(read32(sym, off+int64(off2), config.ctxt.Arch.ByteOrder)))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load8S(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load8S [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load8S [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load8S)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Load8U(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Load8U [off] (I64AddConst [off2] ptr) mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Load8U [off+off2] ptr mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		mem := v_1
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Load8U)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg2(ptr, mem)
		return true
	}
	// match: (I64Load8U [off] (LoweredAddr {sym} [off2] (SB)) _)
	// cond: symIsRO(sym) && isU32Bit(off+int64(off2))
	// result: (I64Const [int64(read8(sym, off+int64(off2)))])
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32LoweredAddr {
			break
		}
		off2 := auxIntToInt32(v_0.AuxInt)
		sym := auxToSym(v_0.Aux)
		v_0_0 := v_0.Args[0]
		if v_0_0.Op != OpSB || !(symIsRO(sym) && isU32Bit(off+int64(off2))) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(read8(sym, off+int64(off2))))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64LtU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64LtU (I64Const [0]) x)
	// result: (I64Eqz (I64Eqz x))
	for {
		if v_0.Op != OpWasm32I64Const || auxIntToInt64(v_0.AuxInt) != 0 {
			break
		}
		x := v_1
		v.reset(OpWasm32I64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	// match: (I64LtU x (I64Const [1]))
	// result: (I64Eqz x)
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const || auxIntToInt64(v_1.AuxInt) != 1 {
			break
		}
		v.reset(OpWasm32I64Eqz)
		v.AddArg(x)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Mul(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Mul (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x * y])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x * y)
		return true
	}
	// match: (I64Mul (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Mul y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Mul)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Ne(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Ne (I64Const [x]) (I64Const [y]))
	// cond: x == y
	// result: (I64Const [0])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x == y) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(0)
		return true
	}
	// match: (I64Ne (I64Const [x]) (I64Const [y]))
	// cond: x != y
	// result: (I64Const [1])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		if !(x != y) {
			break
		}
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(1)
		return true
	}
	// match: (I64Ne (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Ne y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Ne)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	// match: (I64Ne x (I64Const [0]))
	// result: (I64Eqz (I64Eqz x))
	for {
		x := v_0
		if v_1.Op != OpWasm32I64Const || auxIntToInt64(v_1.AuxInt) != 0 {
			break
		}
		v.reset(OpWasm32I64Eqz)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Eqz, typ.Bool)
		v0.AddArg(x)
		v.AddArg(v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Or(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Or (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x | y])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x | y)
		return true
	}
	// match: (I64Or (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Or y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Or)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Shl(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Shl (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x << uint64(y)])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x << uint64(y))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64ShrS(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64ShrS (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x >> uint64(y)])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x >> uint64(y))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64ShrU(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64ShrU (I64Const [x]) (I64Const [y]))
	// result: (I64Const [int64(uint64(x) >> uint64(y))])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(int64(uint64(x) >> uint64(y)))
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Store(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Store [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Store16(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store16 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Store16 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Store16)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Store32(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store32 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Store32 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Store32)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Store8(v *Value) bool {
	v_2 := v.Args[2]
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	// match: (I64Store8 [off] (I64AddConst [off2] ptr) val mem)
	// cond: isU32Bit(off+off2)
	// result: (I64Store8 [off+off2] ptr val mem)
	for {
		off := auxIntToInt64(v.AuxInt)
		if v_0.Op != OpWasm32I64AddConst {
			break
		}
		off2 := auxIntToInt64(v_0.AuxInt)
		ptr := v_0.Args[0]
		val := v_1
		mem := v_2
		if !(isU32Bit(off + off2)) {
			break
		}
		v.reset(OpWasm32I64Store8)
		v.AuxInt = int64ToAuxInt(off + off2)
		v.AddArg3(ptr, val, mem)
		return true
	}
	return false
}
func rewriteValueWasm32_OpWasm32I64Xor(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (I64Xor (I64Const [x]) (I64Const [y]))
	// result: (I64Const [x ^ y])
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		if v_1.Op != OpWasm32I64Const {
			break
		}
		y := auxIntToInt64(v_1.AuxInt)
		v.reset(OpWasm32I64Const)
		v.AuxInt = int64ToAuxInt(x ^ y)
		return true
	}
	// match: (I64Xor (I64Const [x]) y)
	// cond: y.Op != OpWasm32I64Const
	// result: (I64Xor y (I64Const [x]))
	for {
		if v_0.Op != OpWasm32I64Const {
			break
		}
		x := auxIntToInt64(v_0.AuxInt)
		y := v_1
		if !(y.Op != OpWasm32I64Const) {
			break
		}
		v.reset(OpWasm32I64Xor)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(x)
		v.AddArg2(y, v0)
		return true
	}
	return false
}
func rewriteValueWasm32_OpZero(v *Value) bool {
	v_1 := v.Args[1]
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (Zero [0] _ mem)
	// result: mem
	for {
		if auxIntToInt64(v.AuxInt) != 0 {
			break
		}
		mem := v_1
		v.copyOf(mem)
		return true
	}
	// match: (Zero [1] destptr mem)
	// result: (I64Store8 destptr (I64Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 1 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store8)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [2] destptr mem)
	// result: (I64Store16 destptr (I64Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 2 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store16)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [4] destptr mem)
	// result: (I64Store32 destptr (I64Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 4 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store32)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [8] destptr mem)
	// result: (I64Store destptr (I64Const [0]) mem)
	for {
		if auxIntToInt64(v.AuxInt) != 8 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v.AddArg3(destptr, v0, mem)
		return true
	}
	// match: (Zero [3] destptr mem)
	// result: (I64Store8 [2] destptr (I64Const [0]) (I64Store16 destptr (I64Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 3 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store8)
		v.AuxInt = int64ToAuxInt(2)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store16, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [5] destptr mem)
	// result: (I64Store8 [4] destptr (I64Const [0]) (I64Store32 destptr (I64Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 5 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store8)
		v.AuxInt = int64ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [6] destptr mem)
	// result: (I64Store16 [4] destptr (I64Const [0]) (I64Store32 destptr (I64Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 6 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store16)
		v.AuxInt = int64ToAuxInt(4)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [7] destptr mem)
	// result: (I64Store32 [3] destptr (I64Const [0]) (I64Store32 destptr (I64Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 7 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store32)
		v.AuxInt = int64ToAuxInt(3)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store32, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [s] destptr mem)
	// cond: s%8 != 0 && s > 8 && s < 32
	// result: (Zero [s-s%8] (OffPtr <destptr.Type> destptr [s%8]) (I64Store destptr (I64Const [0]) mem))
	for {
		s := auxIntToInt64(v.AuxInt)
		destptr := v_0
		mem := v_1
		if !(s%8 != 0 && s > 8 && s < 32) {
			break
		}
		v.reset(OpZero)
		v.AuxInt = int64ToAuxInt(s - s%8)
		v0 := b.NewValue0(v.Pos, OpOffPtr, destptr.Type)
		v0.AuxInt = int64ToAuxInt(s % 8)
		v0.AddArg(destptr)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v2.AuxInt = int64ToAuxInt(0)
		v1.AddArg3(destptr, v2, mem)
		v.AddArg2(v0, v1)
		return true
	}
	// match: (Zero [16] destptr mem)
	// result: (I64Store [8] destptr (I64Const [0]) (I64Store destptr (I64Const [0]) mem))
	for {
		if auxIntToInt64(v.AuxInt) != 16 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(8)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v1.AddArg3(destptr, v0, mem)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [24] destptr mem)
	// result: (I64Store [16] destptr (I64Const [0]) (I64Store [8] destptr (I64Const [0]) (I64Store destptr (I64Const [0]) mem)))
	for {
		if auxIntToInt64(v.AuxInt) != 24 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(16)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v1.AuxInt = int64ToAuxInt(8)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v2.AddArg3(destptr, v0, mem)
		v1.AddArg3(destptr, v0, v2)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [32] destptr mem)
	// result: (I64Store [24] destptr (I64Const [0]) (I64Store [16] destptr (I64Const [0]) (I64Store [8] destptr (I64Const [0]) (I64Store destptr (I64Const [0]) mem))))
	for {
		if auxIntToInt64(v.AuxInt) != 32 {
			break
		}
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32I64Store)
		v.AuxInt = int64ToAuxInt(24)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0)
		v1 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v1.AuxInt = int64ToAuxInt(16)
		v2 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v2.AuxInt = int64ToAuxInt(8)
		v3 := b.NewValue0(v.Pos, OpWasm32I64Store, types.TypeMem)
		v3.AddArg3(destptr, v0, mem)
		v2.AddArg3(destptr, v0, v3)
		v1.AddArg3(destptr, v0, v2)
		v.AddArg3(destptr, v0, v1)
		return true
	}
	// match: (Zero [s] destptr mem)
	// result: (LoweredZero [s] destptr mem)
	for {
		s := auxIntToInt64(v.AuxInt)
		destptr := v_0
		mem := v_1
		v.reset(OpWasm32LoweredZero)
		v.AuxInt = int64ToAuxInt(s)
		v.AddArg2(destptr, mem)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt16to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt16to32 x:(I64Load16U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load16U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt16to32 x)
	// result: (I64And x (I64Const [0xffff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt16to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt16to64 x:(I64Load16U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load16U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt16to64 x)
	// result: (I64And x (I64Const [0xffff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt32to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt32to64 x:(I64Load32U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load32U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt32to64 x)
	// result: (I64And x (I64Const [0xffffffff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xffffffff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt8to16(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to16 x:(I64Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to16 x)
	// result: (I64And x (I64Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt8to32(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to32 x:(I64Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to32 x)
	// result: (I64And x (I64Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteValueWasm32_OpZeroExt8to64(v *Value) bool {
	v_0 := v.Args[0]
	b := v.Block
	typ := &b.Func.Config.Types
	// match: (ZeroExt8to64 x:(I64Load8U _ _))
	// result: x
	for {
		x := v_0
		if x.Op != OpWasm32I64Load8U {
			break
		}
		v.copyOf(x)
		return true
	}
	// match: (ZeroExt8to64 x)
	// result: (I64And x (I64Const [0xff]))
	for {
		x := v_0
		v.reset(OpWasm32I64And)
		v0 := b.NewValue0(v.Pos, OpWasm32I64Const, typ.Int64)
		v0.AuxInt = int64ToAuxInt(0xff)
		v.AddArg2(x, v0)
		return true
	}
}
func rewriteBlockWasm32(b *Block) bool {
	return false
}
