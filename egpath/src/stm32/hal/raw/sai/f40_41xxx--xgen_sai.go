// +build f40_41xxx

package sai

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type SAI_Periph struct {
	GCR GCR
}

func (p *SAI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SAI1 = (*SAI_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_BASE)))

type GCR_Bits uint32

func (b GCR_Bits) Field(mask GCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask GCR_Bits) J(v int) GCR_Bits {
	return GCR_Bits(bits.Make32(v, uint32(mask)))
}

type GCR struct{ mmio.U32 }

func (r *GCR) Bits(mask GCR_Bits) GCR_Bits { return GCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *GCR) StoreBits(mask, b GCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GCR) SetBits(mask GCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *GCR) ClearBits(mask GCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *GCR) Load() GCR_Bits              { return GCR_Bits(r.U32.Load()) }
func (r *GCR) Store(b GCR_Bits)            { r.U32.Store(uint32(b)) }

func (r *GCR) AtomicStoreBits(mask, b GCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *GCR) AtomicSetBits(mask GCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *GCR) AtomicClearBits(mask GCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type GCR_Mask struct{ mmio.UM32 }

func (rm GCR_Mask) Load() GCR_Bits   { return GCR_Bits(rm.UM32.Load()) }
func (rm GCR_Mask) Store(b GCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SAI_Periph) SYNCIN() GCR_Mask {
	return GCR_Mask{mmio.UM32{&p.GCR.U32, uint32(SYNCIN)}}
}

func (p *SAI_Periph) SYNCOUT() GCR_Mask {
	return GCR_Mask{mmio.UM32{&p.GCR.U32, uint32(SYNCOUT)}}
}
