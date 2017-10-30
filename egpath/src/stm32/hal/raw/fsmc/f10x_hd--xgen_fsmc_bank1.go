// +build f10x_hd

package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank1_Periph struct {
	BTCR [4]BTCR
}

func (p *FSMC_Bank1_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank1 = (*FSMC_Bank1_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank1_R_BASE)))

type BTCR struct {
	BCR BCR
	BTR BTR
}

type BCR_Bits uint32

func (b BCR_Bits) Field(mask BCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BCR_Bits) J(v int) BCR_Bits {
	return BCR_Bits(bits.Make32(v, uint32(mask)))
}

type BCR struct{ mmio.U32 }

func (r *BCR) Bits(mask BCR_Bits) BCR_Bits { return BCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BCR) StoreBits(mask, b BCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BCR) SetBits(mask BCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BCR) ClearBits(mask BCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BCR) Load() BCR_Bits              { return BCR_Bits(r.U32.Load()) }
func (r *BCR) Store(b BCR_Bits)            { r.U32.Store(uint32(b)) }

func (r *BCR) AtomicStoreBits(mask, b BCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BCR) AtomicSetBits(mask BCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BCR) AtomicClearBits(mask BCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BCR_Mask struct{ mmio.UM32 }

func (rm BCR_Mask) Load() BCR_Bits   { return BCR_Bits(rm.UM32.Load()) }
func (rm BCR_Mask) Store(b BCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank1_Periph) MBKEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(MBKEN)}}
}

func (p *FSMC_Bank1_Periph) MUXEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(MUXEN)}}
}

func (p *FSMC_Bank1_Periph) MTYP(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(MTYP)}}
}

func (p *FSMC_Bank1_Periph) MWID(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(MWID)}}
}

func (p *FSMC_Bank1_Periph) FACCEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(FACCEN)}}
}

func (p *FSMC_Bank1_Periph) BURSTEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(BURSTEN)}}
}

func (p *FSMC_Bank1_Periph) WAITPOL(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(WAITPOL)}}
}

func (p *FSMC_Bank1_Periph) WRAPMOD(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(WRAPMOD)}}
}

func (p *FSMC_Bank1_Periph) WAITCFG(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(WAITCFG)}}
}

func (p *FSMC_Bank1_Periph) WREN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(WREN)}}
}

func (p *FSMC_Bank1_Periph) WAITEN(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(WAITEN)}}
}

func (p *FSMC_Bank1_Periph) EXTMOD(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(EXTMOD)}}
}

func (p *FSMC_Bank1_Periph) ASYNCWAIT(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(ASYNCWAIT)}}
}

func (p *FSMC_Bank1_Periph) CBURSTRW(n int) BCR_Mask {
	return BCR_Mask{mmio.UM32{&p.BTCR[n].BCR.U32, uint32(CBURSTRW)}}
}

type BTR_Bits uint32

func (b BTR_Bits) Field(mask BTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BTR_Bits) J(v int) BTR_Bits {
	return BTR_Bits(bits.Make32(v, uint32(mask)))
}

type BTR struct{ mmio.U32 }

func (r *BTR) Bits(mask BTR_Bits) BTR_Bits { return BTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BTR) StoreBits(mask, b BTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BTR) SetBits(mask BTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BTR) ClearBits(mask BTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BTR) Load() BTR_Bits              { return BTR_Bits(r.U32.Load()) }
func (r *BTR) Store(b BTR_Bits)            { r.U32.Store(uint32(b)) }

func (r *BTR) AtomicStoreBits(mask, b BTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BTR) AtomicSetBits(mask BTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BTR) AtomicClearBits(mask BTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BTR_Mask struct{ mmio.UM32 }

func (rm BTR_Mask) Load() BTR_Bits   { return BTR_Bits(rm.UM32.Load()) }
func (rm BTR_Mask) Store(b BTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank1_Periph) ADDSET(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(ADDSET)}}
}

func (p *FSMC_Bank1_Periph) ADDHLD(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(ADDHLD)}}
}

func (p *FSMC_Bank1_Periph) DATAST(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(DATAST)}}
}

func (p *FSMC_Bank1_Periph) BUSTURN(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(BUSTURN)}}
}

func (p *FSMC_Bank1_Periph) CLKDIV(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(CLKDIV)}}
}

func (p *FSMC_Bank1_Periph) DATLAT(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(DATLAT)}}
}

func (p *FSMC_Bank1_Periph) ACCMOD(n int) BTR_Mask {
	return BTR_Mask{mmio.UM32{&p.BTCR[n].BTR.U32, uint32(ACCMOD)}}
}
