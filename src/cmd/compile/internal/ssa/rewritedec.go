// autogenerated from gen/dec.rules: do not edit!
// generated with: cd gen; go run *.go

package ssa

import "math"

var _ = math.MinInt8 // in case not otherwise used
func rewriteValuedec(v *Value) bool {
	switch v.Op {
	case OpComplexImag:
		return rewriteValuedec_OpComplexImag(v)
	case OpComplexReal:
		return rewriteValuedec_OpComplexReal(v)
	case OpIData:
		return rewriteValuedec_OpIData(v)
	case OpITab:
		return rewriteValuedec_OpITab(v)
	case OpLoad:
		return rewriteValuedec_OpLoad(v)
	case OpSliceCap:
		return rewriteValuedec_OpSliceCap(v)
	case OpSliceLen:
		return rewriteValuedec_OpSliceLen(v)
	case OpSlicePtr:
		return rewriteValuedec_OpSlicePtr(v)
	case OpStore:
		return rewriteValuedec_OpStore(v)
	case OpStringLen:
		return rewriteValuedec_OpStringLen(v)
	case OpStringPtr:
		return rewriteValuedec_OpStringPtr(v)
	}
	return false
}
func rewriteValuedec_OpComplexImag(v *Value) bool {
	// match: (ComplexImag (ComplexMake _ imag ))
	// cond:
	// result: imag
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpComplexMake {
			break
		}
		imag := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = imag.Type
		v.AddArg(imag)
		return true
	}
	return false
}
func rewriteValuedec_OpComplexReal(v *Value) bool {
	// match: (ComplexReal (ComplexMake real _  ))
	// cond:
	// result: real
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpComplexMake {
			break
		}
		real := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = real.Type
		v.AddArg(real)
		return true
	}
	return false
}
func rewriteValuedec_OpIData(v *Value) bool {
	// match: (IData (IMake _ data))
	// cond:
	// result: data
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpIMake {
			break
		}
		data := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = data.Type
		v.AddArg(data)
		return true
	}
	return false
}
func rewriteValuedec_OpITab(v *Value) bool {
	b := v.Block
	_ = b
	// match: (ITab (IMake itab _))
	// cond:
	// result: itab
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpIMake {
			break
		}
		itab := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = itab.Type
		v.AddArg(itab)
		return true
	}
	return false
}
func rewriteValuedec_OpLoad(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	// match: (Load <t> ptr mem)
	// cond: t.IsComplex() && t.Size() == 8
	// result: (ComplexMake     (Load <fe.TypeFloat32()> ptr mem)     (Load <fe.TypeFloat32()>       (OffPtr <fe.TypeFloat32().PtrTo()> [4] ptr)       mem)     )
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsComplex() && t.Size() == 8) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Pos, OpLoad, fe.TypeFloat32())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, fe.TypeFloat32())
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeFloat32().PtrTo())
		v2.AuxInt = 4
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsComplex() && t.Size() == 16
	// result: (ComplexMake     (Load <fe.TypeFloat64()> ptr mem)     (Load <fe.TypeFloat64()>       (OffPtr <fe.TypeFloat64().PtrTo()> [8] ptr)       mem)     )
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsComplex() && t.Size() == 16) {
			break
		}
		v.reset(OpComplexMake)
		v0 := b.NewValue0(v.Pos, OpLoad, fe.TypeFloat64())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, fe.TypeFloat64())
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeFloat64().PtrTo())
		v2.AuxInt = 8
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsString()
	// result: (StringMake     (Load <fe.TypeBytePtr()> ptr mem)     (Load <fe.TypeInt()>       (OffPtr <fe.TypeInt().PtrTo()> [config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsString()) {
			break
		}
		v.reset(OpStringMake)
		v0 := b.NewValue0(v.Pos, OpLoad, fe.TypeBytePtr())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, fe.TypeInt())
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsSlice()
	// result: (SliceMake     (Load <t.ElemType().PtrTo()> ptr mem)     (Load <fe.TypeInt()>       (OffPtr <fe.TypeInt().PtrTo()> [config.PtrSize] ptr)       mem)     (Load <fe.TypeInt()>       (OffPtr <fe.TypeInt().PtrTo()> [2*config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsSlice()) {
			break
		}
		v.reset(OpSliceMake)
		v0 := b.NewValue0(v.Pos, OpLoad, t.ElemType().PtrTo())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, fe.TypeInt())
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		v3 := b.NewValue0(v.Pos, OpLoad, fe.TypeInt())
		v4 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v4.AuxInt = 2 * config.PtrSize
		v4.AddArg(ptr)
		v3.AddArg(v4)
		v3.AddArg(mem)
		v.AddArg(v3)
		return true
	}
	// match: (Load <t> ptr mem)
	// cond: t.IsInterface()
	// result: (IMake     (Load <fe.TypeBytePtr()> ptr mem)     (Load <fe.TypeBytePtr()>       (OffPtr <fe.TypeBytePtr().PtrTo()> [config.PtrSize] ptr)       mem))
	for {
		t := v.Type
		ptr := v.Args[0]
		mem := v.Args[1]
		if !(t.IsInterface()) {
			break
		}
		v.reset(OpIMake)
		v0 := b.NewValue0(v.Pos, OpLoad, fe.TypeBytePtr())
		v0.AddArg(ptr)
		v0.AddArg(mem)
		v.AddArg(v0)
		v1 := b.NewValue0(v.Pos, OpLoad, fe.TypeBytePtr())
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeBytePtr().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(ptr)
		v1.AddArg(v2)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec_OpSliceCap(v *Value) bool {
	// match: (SliceCap (SliceMake _ _ cap))
	// cond:
	// result: cap
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		cap := v_0.Args[2]
		v.reset(OpCopy)
		v.Type = cap.Type
		v.AddArg(cap)
		return true
	}
	return false
}
func rewriteValuedec_OpSliceLen(v *Value) bool {
	// match: (SliceLen (SliceMake _ len _))
	// cond:
	// result: len
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		len := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = len.Type
		v.AddArg(len)
		return true
	}
	return false
}
func rewriteValuedec_OpSlicePtr(v *Value) bool {
	// match: (SlicePtr (SliceMake ptr _ _ ))
	// cond:
	// result: ptr
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpSliceMake {
			break
		}
		ptr := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = ptr.Type
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteValuedec_OpStore(v *Value) bool {
	b := v.Block
	_ = b
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	// match: (Store {t} dst (ComplexMake real imag) mem)
	// cond: t.(Type).Size() == 8
	// result: (Store {fe.TypeFloat32()}     (OffPtr <fe.TypeFloat32().PtrTo()> [4] dst)     imag     (Store {fe.TypeFloat32()} dst real mem))
	for {
		t := v.Aux
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpComplexMake {
			break
		}
		real := v_1.Args[0]
		imag := v_1.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 8) {
			break
		}
		v.reset(OpStore)
		v.Aux = fe.TypeFloat32()
		v0 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeFloat32().PtrTo())
		v0.AuxInt = 4
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(imag)
		v1 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v1.Aux = fe.TypeFloat32()
		v1.AddArg(dst)
		v1.AddArg(real)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store {t} dst (ComplexMake real imag) mem)
	// cond: t.(Type).Size() == 16
	// result: (Store {fe.TypeFloat64()}     (OffPtr <fe.TypeFloat64().PtrTo()> [8] dst)     imag     (Store {fe.TypeFloat64()} dst real mem))
	for {
		t := v.Aux
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpComplexMake {
			break
		}
		real := v_1.Args[0]
		imag := v_1.Args[1]
		mem := v.Args[2]
		if !(t.(Type).Size() == 16) {
			break
		}
		v.reset(OpStore)
		v.Aux = fe.TypeFloat64()
		v0 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeFloat64().PtrTo())
		v0.AuxInt = 8
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(imag)
		v1 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v1.Aux = fe.TypeFloat64()
		v1.AddArg(dst)
		v1.AddArg(real)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store dst (StringMake ptr len) mem)
	// cond:
	// result: (Store {fe.TypeInt()}     (OffPtr <fe.TypeInt().PtrTo()> [config.PtrSize] dst)     len     (Store {fe.TypeBytePtr()} dst ptr mem))
	for {
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpStringMake {
			break
		}
		ptr := v_1.Args[0]
		len := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = fe.TypeInt()
		v0 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v0.AuxInt = config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(len)
		v1 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v1.Aux = fe.TypeBytePtr()
		v1.AddArg(dst)
		v1.AddArg(ptr)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	// match: (Store dst (SliceMake ptr len cap) mem)
	// cond:
	// result: (Store {fe.TypeInt()}     (OffPtr <fe.TypeInt().PtrTo()> [2*config.PtrSize] dst)     cap     (Store {fe.TypeInt()}       (OffPtr <fe.TypeInt().PtrTo()> [config.PtrSize] dst)       len       (Store {fe.TypeBytePtr()} dst ptr mem)))
	for {
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpSliceMake {
			break
		}
		ptr := v_1.Args[0]
		len := v_1.Args[1]
		cap := v_1.Args[2]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = fe.TypeInt()
		v0 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v0.AuxInt = 2 * config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(cap)
		v1 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v1.Aux = fe.TypeInt()
		v2 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeInt().PtrTo())
		v2.AuxInt = config.PtrSize
		v2.AddArg(dst)
		v1.AddArg(v2)
		v1.AddArg(len)
		v3 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v3.Aux = fe.TypeBytePtr()
		v3.AddArg(dst)
		v3.AddArg(ptr)
		v3.AddArg(mem)
		v1.AddArg(v3)
		v.AddArg(v1)
		return true
	}
	// match: (Store dst (IMake itab data) mem)
	// cond:
	// result: (Store {fe.TypeBytePtr()}     (OffPtr <fe.TypeBytePtr().PtrTo()> [config.PtrSize] dst)     data     (Store {fe.TypeUintptr()} dst itab mem))
	for {
		dst := v.Args[0]
		v_1 := v.Args[1]
		if v_1.Op != OpIMake {
			break
		}
		itab := v_1.Args[0]
		data := v_1.Args[1]
		mem := v.Args[2]
		v.reset(OpStore)
		v.Aux = fe.TypeBytePtr()
		v0 := b.NewValue0(v.Pos, OpOffPtr, fe.TypeBytePtr().PtrTo())
		v0.AuxInt = config.PtrSize
		v0.AddArg(dst)
		v.AddArg(v0)
		v.AddArg(data)
		v1 := b.NewValue0(v.Pos, OpStore, TypeMem)
		v1.Aux = fe.TypeUintptr()
		v1.AddArg(dst)
		v1.AddArg(itab)
		v1.AddArg(mem)
		v.AddArg(v1)
		return true
	}
	return false
}
func rewriteValuedec_OpStringLen(v *Value) bool {
	// match: (StringLen (StringMake _ len))
	// cond:
	// result: len
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		len := v_0.Args[1]
		v.reset(OpCopy)
		v.Type = len.Type
		v.AddArg(len)
		return true
	}
	return false
}
func rewriteValuedec_OpStringPtr(v *Value) bool {
	// match: (StringPtr (StringMake ptr _))
	// cond:
	// result: ptr
	for {
		v_0 := v.Args[0]
		if v_0.Op != OpStringMake {
			break
		}
		ptr := v_0.Args[0]
		v.reset(OpCopy)
		v.Type = ptr.Type
		v.AddArg(ptr)
		return true
	}
	return false
}
func rewriteBlockdec(b *Block) bool {
	config := b.Func.Config
	_ = config
	fe := b.Func.fe
	_ = fe
	switch b.Kind {
	}
	return false
}
