package ob

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type OB_Periph struct {
	RDP     RDP
	USER    USER
	WRP01   WRP01
	WRP23   WRP23
	WRP45   WRP45
	WRP67   WRP67
	WRP89   WRP89
	WRP1011 WRP1011
	_       [24]uint32
	WRP1213 WRP1213
	WRP1415 WRP1415
}

func (p *OB_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OB = (*OB_Periph)(unsafe.Pointer(uintptr(mmap.OB_BASE)))

type RDP_Bits uint32

func (b RDP_Bits) Field(mask RDP_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RDP_Bits) J(v int) RDP_Bits {
	return RDP_Bits(bits.Make32(v, uint32(mask)))
}

type RDP struct{ mmio.U32 }

func (r *RDP) Bits(mask RDP_Bits) RDP_Bits { return RDP_Bits(r.U32.Bits(uint32(mask))) }
func (r *RDP) StoreBits(mask, b RDP_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDP) SetBits(mask RDP_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *RDP) ClearBits(mask RDP_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDP) Load() RDP_Bits              { return RDP_Bits(r.U32.Load()) }
func (r *RDP) Store(b RDP_Bits)            { r.U32.Store(uint32(b)) }

func (r *RDP) AtomicStoreBits(mask, b RDP_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDP) AtomicSetBits(mask RDP_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDP) AtomicClearBits(mask RDP_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type RDP_Mask struct{ mmio.UM32 }

func (rm RDP_Mask) Load() RDP_Bits   { return RDP_Bits(rm.UM32.Load()) }
func (rm RDP_Mask) Store(b RDP_Bits) { rm.UM32.Store(uint32(b)) }

type USER_Bits uint32

func (b USER_Bits) Field(mask USER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask USER_Bits) J(v int) USER_Bits {
	return USER_Bits(bits.Make32(v, uint32(mask)))
}

type USER struct{ mmio.U32 }

func (r *USER) Bits(mask USER_Bits) USER_Bits { return USER_Bits(r.U32.Bits(uint32(mask))) }
func (r *USER) StoreBits(mask, b USER_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *USER) SetBits(mask USER_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *USER) ClearBits(mask USER_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *USER) Load() USER_Bits               { return USER_Bits(r.U32.Load()) }
func (r *USER) Store(b USER_Bits)             { r.U32.Store(uint32(b)) }

func (r *USER) AtomicStoreBits(mask, b USER_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *USER) AtomicSetBits(mask USER_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *USER) AtomicClearBits(mask USER_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type USER_Mask struct{ mmio.UM32 }

func (rm USER_Mask) Load() USER_Bits   { return USER_Bits(rm.UM32.Load()) }
func (rm USER_Mask) Store(b USER_Bits) { rm.UM32.Store(uint32(b)) }

type WRP01_Bits uint32

func (b WRP01_Bits) Field(mask WRP01_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP01_Bits) J(v int) WRP01_Bits {
	return WRP01_Bits(bits.Make32(v, uint32(mask)))
}

type WRP01 struct{ mmio.U32 }

func (r *WRP01) Bits(mask WRP01_Bits) WRP01_Bits { return WRP01_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP01) StoreBits(mask, b WRP01_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP01) SetBits(mask WRP01_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRP01) ClearBits(mask WRP01_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRP01) Load() WRP01_Bits                { return WRP01_Bits(r.U32.Load()) }
func (r *WRP01) Store(b WRP01_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRP01) AtomicStoreBits(mask, b WRP01_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRP01) AtomicSetBits(mask WRP01_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP01) AtomicClearBits(mask WRP01_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRP01_Mask struct{ mmio.UM32 }

func (rm WRP01_Mask) Load() WRP01_Bits   { return WRP01_Bits(rm.UM32.Load()) }
func (rm WRP01_Mask) Store(b WRP01_Bits) { rm.UM32.Store(uint32(b)) }

type WRP23_Bits uint32

func (b WRP23_Bits) Field(mask WRP23_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP23_Bits) J(v int) WRP23_Bits {
	return WRP23_Bits(bits.Make32(v, uint32(mask)))
}

type WRP23 struct{ mmio.U32 }

func (r *WRP23) Bits(mask WRP23_Bits) WRP23_Bits { return WRP23_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP23) StoreBits(mask, b WRP23_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP23) SetBits(mask WRP23_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRP23) ClearBits(mask WRP23_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRP23) Load() WRP23_Bits                { return WRP23_Bits(r.U32.Load()) }
func (r *WRP23) Store(b WRP23_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRP23) AtomicStoreBits(mask, b WRP23_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRP23) AtomicSetBits(mask WRP23_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP23) AtomicClearBits(mask WRP23_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRP23_Mask struct{ mmio.UM32 }

func (rm WRP23_Mask) Load() WRP23_Bits   { return WRP23_Bits(rm.UM32.Load()) }
func (rm WRP23_Mask) Store(b WRP23_Bits) { rm.UM32.Store(uint32(b)) }

type WRP45_Bits uint32

func (b WRP45_Bits) Field(mask WRP45_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP45_Bits) J(v int) WRP45_Bits {
	return WRP45_Bits(bits.Make32(v, uint32(mask)))
}

type WRP45 struct{ mmio.U32 }

func (r *WRP45) Bits(mask WRP45_Bits) WRP45_Bits { return WRP45_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP45) StoreBits(mask, b WRP45_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP45) SetBits(mask WRP45_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRP45) ClearBits(mask WRP45_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRP45) Load() WRP45_Bits                { return WRP45_Bits(r.U32.Load()) }
func (r *WRP45) Store(b WRP45_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRP45) AtomicStoreBits(mask, b WRP45_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRP45) AtomicSetBits(mask WRP45_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP45) AtomicClearBits(mask WRP45_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRP45_Mask struct{ mmio.UM32 }

func (rm WRP45_Mask) Load() WRP45_Bits   { return WRP45_Bits(rm.UM32.Load()) }
func (rm WRP45_Mask) Store(b WRP45_Bits) { rm.UM32.Store(uint32(b)) }

type WRP67_Bits uint32

func (b WRP67_Bits) Field(mask WRP67_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP67_Bits) J(v int) WRP67_Bits {
	return WRP67_Bits(bits.Make32(v, uint32(mask)))
}

type WRP67 struct{ mmio.U32 }

func (r *WRP67) Bits(mask WRP67_Bits) WRP67_Bits { return WRP67_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP67) StoreBits(mask, b WRP67_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP67) SetBits(mask WRP67_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRP67) ClearBits(mask WRP67_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRP67) Load() WRP67_Bits                { return WRP67_Bits(r.U32.Load()) }
func (r *WRP67) Store(b WRP67_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRP67) AtomicStoreBits(mask, b WRP67_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRP67) AtomicSetBits(mask WRP67_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP67) AtomicClearBits(mask WRP67_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRP67_Mask struct{ mmio.UM32 }

func (rm WRP67_Mask) Load() WRP67_Bits   { return WRP67_Bits(rm.UM32.Load()) }
func (rm WRP67_Mask) Store(b WRP67_Bits) { rm.UM32.Store(uint32(b)) }

type WRP89_Bits uint32

func (b WRP89_Bits) Field(mask WRP89_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP89_Bits) J(v int) WRP89_Bits {
	return WRP89_Bits(bits.Make32(v, uint32(mask)))
}

type WRP89 struct{ mmio.U32 }

func (r *WRP89) Bits(mask WRP89_Bits) WRP89_Bits { return WRP89_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP89) StoreBits(mask, b WRP89_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP89) SetBits(mask WRP89_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRP89) ClearBits(mask WRP89_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRP89) Load() WRP89_Bits                { return WRP89_Bits(r.U32.Load()) }
func (r *WRP89) Store(b WRP89_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRP89) AtomicStoreBits(mask, b WRP89_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRP89) AtomicSetBits(mask WRP89_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP89) AtomicClearBits(mask WRP89_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRP89_Mask struct{ mmio.UM32 }

func (rm WRP89_Mask) Load() WRP89_Bits   { return WRP89_Bits(rm.UM32.Load()) }
func (rm WRP89_Mask) Store(b WRP89_Bits) { rm.UM32.Store(uint32(b)) }

type WRP1011_Bits uint32

func (b WRP1011_Bits) Field(mask WRP1011_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1011_Bits) J(v int) WRP1011_Bits {
	return WRP1011_Bits(bits.Make32(v, uint32(mask)))
}

type WRP1011 struct{ mmio.U32 }

func (r *WRP1011) Bits(mask WRP1011_Bits) WRP1011_Bits { return WRP1011_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP1011) StoreBits(mask, b WRP1011_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP1011) SetBits(mask WRP1011_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *WRP1011) ClearBits(mask WRP1011_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *WRP1011) Load() WRP1011_Bits                  { return WRP1011_Bits(r.U32.Load()) }
func (r *WRP1011) Store(b WRP1011_Bits)                { r.U32.Store(uint32(b)) }

func (r *WRP1011) AtomicStoreBits(mask, b WRP1011_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *WRP1011) AtomicSetBits(mask WRP1011_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP1011) AtomicClearBits(mask WRP1011_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP1011_Mask struct{ mmio.UM32 }

func (rm WRP1011_Mask) Load() WRP1011_Bits   { return WRP1011_Bits(rm.UM32.Load()) }
func (rm WRP1011_Mask) Store(b WRP1011_Bits) { rm.UM32.Store(uint32(b)) }

type WRP1213_Bits uint32

func (b WRP1213_Bits) Field(mask WRP1213_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1213_Bits) J(v int) WRP1213_Bits {
	return WRP1213_Bits(bits.Make32(v, uint32(mask)))
}

type WRP1213 struct{ mmio.U32 }

func (r *WRP1213) Bits(mask WRP1213_Bits) WRP1213_Bits { return WRP1213_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP1213) StoreBits(mask, b WRP1213_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP1213) SetBits(mask WRP1213_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *WRP1213) ClearBits(mask WRP1213_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *WRP1213) Load() WRP1213_Bits                  { return WRP1213_Bits(r.U32.Load()) }
func (r *WRP1213) Store(b WRP1213_Bits)                { r.U32.Store(uint32(b)) }

func (r *WRP1213) AtomicStoreBits(mask, b WRP1213_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *WRP1213) AtomicSetBits(mask WRP1213_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP1213) AtomicClearBits(mask WRP1213_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP1213_Mask struct{ mmio.UM32 }

func (rm WRP1213_Mask) Load() WRP1213_Bits   { return WRP1213_Bits(rm.UM32.Load()) }
func (rm WRP1213_Mask) Store(b WRP1213_Bits) { rm.UM32.Store(uint32(b)) }

type WRP1415_Bits uint32

func (b WRP1415_Bits) Field(mask WRP1415_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1415_Bits) J(v int) WRP1415_Bits {
	return WRP1415_Bits(bits.Make32(v, uint32(mask)))
}

type WRP1415 struct{ mmio.U32 }

func (r *WRP1415) Bits(mask WRP1415_Bits) WRP1415_Bits { return WRP1415_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRP1415) StoreBits(mask, b WRP1415_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRP1415) SetBits(mask WRP1415_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *WRP1415) ClearBits(mask WRP1415_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *WRP1415) Load() WRP1415_Bits                  { return WRP1415_Bits(r.U32.Load()) }
func (r *WRP1415) Store(b WRP1415_Bits)                { r.U32.Store(uint32(b)) }

func (r *WRP1415) AtomicStoreBits(mask, b WRP1415_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *WRP1415) AtomicSetBits(mask WRP1415_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRP1415) AtomicClearBits(mask WRP1415_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type WRP1415_Mask struct{ mmio.UM32 }

func (rm WRP1415_Mask) Load() WRP1415_Bits   { return WRP1415_Bits(rm.UM32.Load()) }
func (rm WRP1415_Mask) Store(b WRP1415_Bits) { rm.UM32.Store(uint32(b)) }
