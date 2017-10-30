package fmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type FMC_Bank1E_Periph struct {
	BWTR [7]BWTR
}

func (p *FMC_Bank1E_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FMC_Bank1E = (*FMC_Bank1E_Periph)(unsafe.Pointer(uintptr(mmap.FMC_Bank1E_R_BASE)))

type BWTR_Bits uint32

func (b BWTR_Bits) Field(mask BWTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BWTR_Bits) J(v int) BWTR_Bits {
	return BWTR_Bits(bits.Make32(v, uint32(mask)))
}

type BWTR struct{ mmio.U32 }

func (r *BWTR) Bits(mask BWTR_Bits) BWTR_Bits { return BWTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BWTR) StoreBits(mask, b BWTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BWTR) SetBits(mask BWTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BWTR) ClearBits(mask BWTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BWTR) Load() BWTR_Bits               { return BWTR_Bits(r.U32.Load()) }
func (r *BWTR) Store(b BWTR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BWTR) AtomicStoreBits(mask, b BWTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BWTR) AtomicSetBits(mask BWTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BWTR) AtomicClearBits(mask BWTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BWTR_Mask struct{ mmio.UM32 }

func (rm BWTR_Mask) Load() BWTR_Bits   { return BWTR_Bits(rm.UM32.Load()) }
func (rm BWTR_Mask) Store(b BWTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank1E_Periph) EADDSET(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EADDSET)}}
}

func (p *FMC_Bank1E_Periph) EADDHLD(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EADDHLD)}}
}

func (p *FMC_Bank1E_Periph) EDATAST(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EDATAST)}}
}

func (p *FMC_Bank1E_Periph) EBUSTURN(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EBUSTURN)}}
}

func (p *FMC_Bank1E_Periph) EACCMOD(n int) BWTR_Mask {
	return BWTR_Mask{mmio.UM32{&p.BWTR[n].U32, uint32(EACCMOD)}}
}
