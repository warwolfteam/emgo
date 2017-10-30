// +build f10x_hd

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank4_Periph struct {
	PCR4  PCR4
	SR4   SR4
	PMEM4 PMEM4
	PATT4 PATT4
	PIO4  PIO4
}

func (p *FSMC_Bank4_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank4 = (*FSMC_Bank4_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank4_R_BASE)))

type PCR4_Bits uint32

func (b PCR4_Bits) Field(mask PCR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR4_Bits) J(v int) PCR4_Bits {
	return PCR4_Bits(bits.Make32(v, uint32(mask)))
}

type PCR4 struct{ mmio.U32 }

func (r *PCR4) Bits(mask PCR4_Bits) PCR4_Bits { return PCR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PCR4) StoreBits(mask, b PCR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PCR4) SetBits(mask PCR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PCR4) ClearBits(mask PCR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PCR4) Load() PCR4_Bits               { return PCR4_Bits(r.U32.Load()) }
func (r *PCR4) Store(b PCR4_Bits)             { r.U32.Store(uint32(b)) }

func (r *PCR4) AtomicStoreBits(mask, b PCR4_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PCR4) AtomicSetBits(mask PCR4_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PCR4) AtomicClearBits(mask PCR4_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PCR4_Mask struct{ mmio.UM32 }

func (rm PCR4_Mask) Load() PCR4_Bits   { return PCR4_Bits(rm.UM32.Load()) }
func (rm PCR4_Mask) Store(b PCR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) PWAITEN() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank4_Periph) PBKEN() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank4_Periph) PTYP() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank4_Periph) PWID() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(PWID)}}
}

func (p *FSMC_Bank4_Periph) ECCEN() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank4_Periph) TCLR() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank4_Periph) TAR() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(TAR)}}
}

func (p *FSMC_Bank4_Periph) ECCPS() PCR4_Mask {
	return PCR4_Mask{mmio.UM32{&p.PCR4.U32, uint32(ECCPS)}}
}

type SR4_Bits uint32

func (b SR4_Bits) Field(mask SR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR4_Bits) J(v int) SR4_Bits {
	return SR4_Bits(bits.Make32(v, uint32(mask)))
}

type SR4 struct{ mmio.U32 }

func (r *SR4) Bits(mask SR4_Bits) SR4_Bits { return SR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR4) StoreBits(mask, b SR4_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR4) SetBits(mask SR4_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SR4) ClearBits(mask SR4_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SR4) Load() SR4_Bits              { return SR4_Bits(r.U32.Load()) }
func (r *SR4) Store(b SR4_Bits)            { r.U32.Store(uint32(b)) }

func (r *SR4) AtomicStoreBits(mask, b SR4_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *SR4) AtomicSetBits(mask SR4_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SR4) AtomicClearBits(mask SR4_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type SR4_Mask struct{ mmio.UM32 }

func (rm SR4_Mask) Load() SR4_Bits   { return SR4_Bits(rm.UM32.Load()) }
func (rm SR4_Mask) Store(b SR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) IRS() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(IRS)}}
}

func (p *FSMC_Bank4_Periph) ILS() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(ILS)}}
}

func (p *FSMC_Bank4_Periph) IFS() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(IFS)}}
}

func (p *FSMC_Bank4_Periph) IREN() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(IREN)}}
}

func (p *FSMC_Bank4_Periph) ILEN() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank4_Periph) IFEN() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank4_Periph) FEMPT() SR4_Mask {
	return SR4_Mask{mmio.UM32{&p.SR4.U32, uint32(FEMPT)}}
}

type PMEM4_Bits uint32

func (b PMEM4_Bits) Field(mask PMEM4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM4_Bits) J(v int) PMEM4_Bits {
	return PMEM4_Bits(bits.Make32(v, uint32(mask)))
}

type PMEM4 struct{ mmio.U32 }

func (r *PMEM4) Bits(mask PMEM4_Bits) PMEM4_Bits { return PMEM4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PMEM4) StoreBits(mask, b PMEM4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PMEM4) SetBits(mask PMEM4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PMEM4) ClearBits(mask PMEM4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PMEM4) Load() PMEM4_Bits                { return PMEM4_Bits(r.U32.Load()) }
func (r *PMEM4) Store(b PMEM4_Bits)              { r.U32.Store(uint32(b)) }

func (r *PMEM4) AtomicStoreBits(mask, b PMEM4_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PMEM4) AtomicSetBits(mask PMEM4_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PMEM4) AtomicClearBits(mask PMEM4_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PMEM4_Mask struct{ mmio.UM32 }

func (rm PMEM4_Mask) Load() PMEM4_Bits   { return PMEM4_Bits(rm.UM32.Load()) }
func (rm PMEM4_Mask) Store(b PMEM4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) MEMSET4() PMEM4_Mask {
	return PMEM4_Mask{mmio.UM32{&p.PMEM4.U32, uint32(MEMSET4)}}
}

func (p *FSMC_Bank4_Periph) MEMWAIT4() PMEM4_Mask {
	return PMEM4_Mask{mmio.UM32{&p.PMEM4.U32, uint32(MEMWAIT4)}}
}

func (p *FSMC_Bank4_Periph) MEMHOLD4() PMEM4_Mask {
	return PMEM4_Mask{mmio.UM32{&p.PMEM4.U32, uint32(MEMHOLD4)}}
}

func (p *FSMC_Bank4_Periph) MEMHIZ4() PMEM4_Mask {
	return PMEM4_Mask{mmio.UM32{&p.PMEM4.U32, uint32(MEMHIZ4)}}
}

type PATT4_Bits uint32

func (b PATT4_Bits) Field(mask PATT4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT4_Bits) J(v int) PATT4_Bits {
	return PATT4_Bits(bits.Make32(v, uint32(mask)))
}

type PATT4 struct{ mmio.U32 }

func (r *PATT4) Bits(mask PATT4_Bits) PATT4_Bits { return PATT4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PATT4) StoreBits(mask, b PATT4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PATT4) SetBits(mask PATT4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PATT4) ClearBits(mask PATT4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PATT4) Load() PATT4_Bits                { return PATT4_Bits(r.U32.Load()) }
func (r *PATT4) Store(b PATT4_Bits)              { r.U32.Store(uint32(b)) }

func (r *PATT4) AtomicStoreBits(mask, b PATT4_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PATT4) AtomicSetBits(mask PATT4_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PATT4) AtomicClearBits(mask PATT4_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PATT4_Mask struct{ mmio.UM32 }

func (rm PATT4_Mask) Load() PATT4_Bits   { return PATT4_Bits(rm.UM32.Load()) }
func (rm PATT4_Mask) Store(b PATT4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) ATTSET4() PATT4_Mask {
	return PATT4_Mask{mmio.UM32{&p.PATT4.U32, uint32(ATTSET4)}}
}

func (p *FSMC_Bank4_Periph) ATTWAIT4() PATT4_Mask {
	return PATT4_Mask{mmio.UM32{&p.PATT4.U32, uint32(ATTWAIT4)}}
}

func (p *FSMC_Bank4_Periph) ATTHOLD4() PATT4_Mask {
	return PATT4_Mask{mmio.UM32{&p.PATT4.U32, uint32(ATTHOLD4)}}
}

func (p *FSMC_Bank4_Periph) ATTHIZ4() PATT4_Mask {
	return PATT4_Mask{mmio.UM32{&p.PATT4.U32, uint32(ATTHIZ4)}}
}

type PIO4_Bits uint32

func (b PIO4_Bits) Field(mask PIO4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PIO4_Bits) J(v int) PIO4_Bits {
	return PIO4_Bits(bits.Make32(v, uint32(mask)))
}

type PIO4 struct{ mmio.U32 }

func (r *PIO4) Bits(mask PIO4_Bits) PIO4_Bits { return PIO4_Bits(r.U32.Bits(uint32(mask))) }
func (r *PIO4) StoreBits(mask, b PIO4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PIO4) SetBits(mask PIO4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PIO4) ClearBits(mask PIO4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PIO4) Load() PIO4_Bits               { return PIO4_Bits(r.U32.Load()) }
func (r *PIO4) Store(b PIO4_Bits)             { r.U32.Store(uint32(b)) }

func (r *PIO4) AtomicStoreBits(mask, b PIO4_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PIO4) AtomicSetBits(mask PIO4_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PIO4) AtomicClearBits(mask PIO4_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PIO4_Mask struct{ mmio.UM32 }

func (rm PIO4_Mask) Load() PIO4_Bits   { return PIO4_Bits(rm.UM32.Load()) }
func (rm PIO4_Mask) Store(b PIO4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) IOSET4() PIO4_Mask {
	return PIO4_Mask{mmio.UM32{&p.PIO4.U32, uint32(IOSET4)}}
}

func (p *FSMC_Bank4_Periph) IOWAIT4() PIO4_Mask {
	return PIO4_Mask{mmio.UM32{&p.PIO4.U32, uint32(IOWAIT4)}}
}

func (p *FSMC_Bank4_Periph) IOHOLD4() PIO4_Mask {
	return PIO4_Mask{mmio.UM32{&p.PIO4.U32, uint32(IOHOLD4)}}
}

func (p *FSMC_Bank4_Periph) IOHIZ4() PIO4_Mask {
	return PIO4_Mask{mmio.UM32{&p.PIO4.U32, uint32(IOHIZ4)}}
}
