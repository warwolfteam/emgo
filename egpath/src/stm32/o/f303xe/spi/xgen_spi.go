package spi

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type SPI_Periph struct {
	CR1     CR1
	CR2     CR2
	SR      SR
	DR      DR
	CRCPR   CRCPR
	RXCRCR  RXCRCR
	TXCRCR  TXCRCR
	I2SCFGR I2SCFGR
	I2SPR   I2SPR
}

func (p *SPI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var I2S2ext = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.I2S2ext_BASE)))

//emgo:const
var SPI2 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI2_BASE)))

//emgo:const
var SPI3 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI3_BASE)))

//emgo:const
var I2S3ext = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.I2S3ext_BASE)))

//emgo:const
var SPI1 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI1_BASE)))

//emgo:const
var SPI4 = (*SPI_Periph)(unsafe.Pointer(uintptr(mmap.SPI4_BASE)))

type CR1_Bits uint32

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

func (r *CR1) AtomicStoreBits(mask, b CR1_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CR1) AtomicSetBits(mask CR1_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR1) AtomicClearBits(mask CR1_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) CPHA() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CPHA)}}
}

func (p *SPI_Periph) CPOL() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CPOL)}}
}

func (p *SPI_Periph) MSTR() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(MSTR)}}
}

func (p *SPI_Periph) BR() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(BR)}}
}

func (p *SPI_Periph) SPE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SPE)}}
}

func (p *SPI_Periph) LSBFIRST() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(LSBFIRST)}}
}

func (p *SPI_Periph) SSI() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SSI)}}
}

func (p *SPI_Periph) SSM() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SSM)}}
}

func (p *SPI_Periph) RXONLY() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RXONLY)}}
}

func (p *SPI_Periph) CRCL() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CRCL)}}
}

func (p *SPI_Periph) CRCNEXT() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CRCNEXT)}}
}

func (p *SPI_Periph) CRCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(CRCEN)}}
}

func (p *SPI_Periph) BIDIOE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(BIDIOE)}}
}

func (p *SPI_Periph) BIDIMODE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(BIDIMODE)}}
}

type CR2_Bits uint32

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

func (r *CR2) AtomicStoreBits(mask, b CR2_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CR2) AtomicSetBits(mask CR2_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR2) AtomicClearBits(mask CR2_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXDMAEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RXDMAEN)}}
}

func (p *SPI_Periph) TXDMAEN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(TXDMAEN)}}
}

func (p *SPI_Periph) SSOE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(SSOE)}}
}

func (p *SPI_Periph) NSSP() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(NSSP)}}
}

func (p *SPI_Periph) FRF() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(FRF)}}
}

func (p *SPI_Periph) ERRIE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ERRIE)}}
}

func (p *SPI_Periph) RXNEIE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RXNEIE)}}
}

func (p *SPI_Periph) TXEIE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(TXEIE)}}
}

func (p *SPI_Periph) DS() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(DS)}}
}

func (p *SPI_Periph) FRXTH() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(FRXTH)}}
}

func (p *SPI_Periph) LDMARX() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LDMARX)}}
}

func (p *SPI_Periph) LDMATX() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(LDMATX)}}
}

type SR_Bits uint32

func (b SR_Bits) Field(mask SR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR_Bits) J(v int) SR_Bits {
	return SR_Bits(bits.Make32(v, uint32(mask)))
}

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

func (r *SR) AtomicStoreBits(mask, b SR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *SR) AtomicSetBits(mask SR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SR) AtomicClearBits(mask SR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXNE() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(RXNE)}}
}

func (p *SPI_Periph) TXE() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(TXE)}}
}

func (p *SPI_Periph) CHSIDE() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(CHSIDE)}}
}

func (p *SPI_Periph) UDR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(UDR)}}
}

func (p *SPI_Periph) CRCERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(CRCERR)}}
}

func (p *SPI_Periph) MODF() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(MODF)}}
}

func (p *SPI_Periph) OVR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OVR)}}
}

func (p *SPI_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

func (p *SPI_Periph) FRE() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FRE)}}
}

func (p *SPI_Periph) FRLVL() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FRLVL)}}
}

func (p *SPI_Periph) FTLVL() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(FTLVL)}}
}

type DR_Bits uint32

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

func (r *DR) AtomicStoreBits(mask, b DR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *DR) AtomicSetBits(mask DR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DR) AtomicClearBits(mask DR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }

type CRCPR_Bits uint32

func (b CRCPR_Bits) Field(mask CRCPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CRCPR_Bits) J(v int) CRCPR_Bits {
	return CRCPR_Bits(bits.Make32(v, uint32(mask)))
}

type CRCPR struct{ mmio.U32 }

func (r *CRCPR) Bits(mask CRCPR_Bits) CRCPR_Bits { return CRCPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CRCPR) StoreBits(mask, b CRCPR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CRCPR) SetBits(mask CRCPR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CRCPR) ClearBits(mask CRCPR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CRCPR) Load() CRCPR_Bits                { return CRCPR_Bits(r.U32.Load()) }
func (r *CRCPR) Store(b CRCPR_Bits)              { r.U32.Store(uint32(b)) }

func (r *CRCPR) AtomicStoreBits(mask, b CRCPR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CRCPR) AtomicSetBits(mask CRCPR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CRCPR) AtomicClearBits(mask CRCPR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CRCPR_Mask struct{ mmio.UM32 }

func (rm CRCPR_Mask) Load() CRCPR_Bits   { return CRCPR_Bits(rm.UM32.Load()) }
func (rm CRCPR_Mask) Store(b CRCPR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) CRCPOLY() CRCPR_Mask {
	return CRCPR_Mask{mmio.UM32{&p.CRCPR.U32, uint32(CRCPOLY)}}
}

type RXCRCR_Bits uint32

func (b RXCRCR_Bits) Field(mask RXCRCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXCRCR_Bits) J(v int) RXCRCR_Bits {
	return RXCRCR_Bits(bits.Make32(v, uint32(mask)))
}

type RXCRCR struct{ mmio.U32 }

func (r *RXCRCR) Bits(mask RXCRCR_Bits) RXCRCR_Bits { return RXCRCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RXCRCR) StoreBits(mask, b RXCRCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXCRCR) SetBits(mask RXCRCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *RXCRCR) ClearBits(mask RXCRCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *RXCRCR) Load() RXCRCR_Bits                 { return RXCRCR_Bits(r.U32.Load()) }
func (r *RXCRCR) Store(b RXCRCR_Bits)               { r.U32.Store(uint32(b)) }

func (r *RXCRCR) AtomicStoreBits(mask, b RXCRCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RXCRCR) AtomicSetBits(mask RXCRCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RXCRCR) AtomicClearBits(mask RXCRCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type RXCRCR_Mask struct{ mmio.UM32 }

func (rm RXCRCR_Mask) Load() RXCRCR_Bits   { return RXCRCR_Bits(rm.UM32.Load()) }
func (rm RXCRCR_Mask) Store(b RXCRCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) RXCRC() RXCRCR_Mask {
	return RXCRCR_Mask{mmio.UM32{&p.RXCRCR.U32, uint32(RXCRC)}}
}

type TXCRCR_Bits uint32

func (b TXCRCR_Bits) Field(mask TXCRCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXCRCR_Bits) J(v int) TXCRCR_Bits {
	return TXCRCR_Bits(bits.Make32(v, uint32(mask)))
}

type TXCRCR struct{ mmio.U32 }

func (r *TXCRCR) Bits(mask TXCRCR_Bits) TXCRCR_Bits { return TXCRCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TXCRCR) StoreBits(mask, b TXCRCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TXCRCR) SetBits(mask TXCRCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TXCRCR) ClearBits(mask TXCRCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TXCRCR) Load() TXCRCR_Bits                 { return TXCRCR_Bits(r.U32.Load()) }
func (r *TXCRCR) Store(b TXCRCR_Bits)               { r.U32.Store(uint32(b)) }

func (r *TXCRCR) AtomicStoreBits(mask, b TXCRCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TXCRCR) AtomicSetBits(mask TXCRCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TXCRCR) AtomicClearBits(mask TXCRCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TXCRCR_Mask struct{ mmio.UM32 }

func (rm TXCRCR_Mask) Load() TXCRCR_Bits   { return TXCRCR_Bits(rm.UM32.Load()) }
func (rm TXCRCR_Mask) Store(b TXCRCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) TXCRC() TXCRCR_Mask {
	return TXCRCR_Mask{mmio.UM32{&p.TXCRCR.U32, uint32(TXCRC)}}
}

type I2SCFGR_Bits uint32

func (b I2SCFGR_Bits) Field(mask I2SCFGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SCFGR_Bits) J(v int) I2SCFGR_Bits {
	return I2SCFGR_Bits(bits.Make32(v, uint32(mask)))
}

type I2SCFGR struct{ mmio.U32 }

func (r *I2SCFGR) Bits(mask I2SCFGR_Bits) I2SCFGR_Bits { return I2SCFGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *I2SCFGR) StoreBits(mask, b I2SCFGR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *I2SCFGR) SetBits(mask I2SCFGR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *I2SCFGR) ClearBits(mask I2SCFGR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *I2SCFGR) Load() I2SCFGR_Bits                  { return I2SCFGR_Bits(r.U32.Load()) }
func (r *I2SCFGR) Store(b I2SCFGR_Bits)                { r.U32.Store(uint32(b)) }

func (r *I2SCFGR) AtomicStoreBits(mask, b I2SCFGR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *I2SCFGR) AtomicSetBits(mask I2SCFGR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *I2SCFGR) AtomicClearBits(mask I2SCFGR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type I2SCFGR_Mask struct{ mmio.UM32 }

func (rm I2SCFGR_Mask) Load() I2SCFGR_Bits   { return I2SCFGR_Bits(rm.UM32.Load()) }
func (rm I2SCFGR_Mask) Store(b I2SCFGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) CHLEN() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(CHLEN)}}
}

func (p *SPI_Periph) DATLEN() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(DATLEN)}}
}

func (p *SPI_Periph) CKPOL() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(CKPOL)}}
}

func (p *SPI_Periph) I2SSTD() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SSTD)}}
}

func (p *SPI_Periph) PCMSYNC() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(PCMSYNC)}}
}

func (p *SPI_Periph) I2SCFG() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SCFG)}}
}

func (p *SPI_Periph) I2SE() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SE)}}
}

func (p *SPI_Periph) I2SMOD() I2SCFGR_Mask {
	return I2SCFGR_Mask{mmio.UM32{&p.I2SCFGR.U32, uint32(I2SMOD)}}
}

type I2SPR_Bits uint32

func (b I2SPR_Bits) Field(mask I2SPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask I2SPR_Bits) J(v int) I2SPR_Bits {
	return I2SPR_Bits(bits.Make32(v, uint32(mask)))
}

type I2SPR struct{ mmio.U32 }

func (r *I2SPR) Bits(mask I2SPR_Bits) I2SPR_Bits { return I2SPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *I2SPR) StoreBits(mask, b I2SPR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *I2SPR) SetBits(mask I2SPR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *I2SPR) ClearBits(mask I2SPR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *I2SPR) Load() I2SPR_Bits                { return I2SPR_Bits(r.U32.Load()) }
func (r *I2SPR) Store(b I2SPR_Bits)              { r.U32.Store(uint32(b)) }

func (r *I2SPR) AtomicStoreBits(mask, b I2SPR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *I2SPR) AtomicSetBits(mask I2SPR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *I2SPR) AtomicClearBits(mask I2SPR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type I2SPR_Mask struct{ mmio.UM32 }

func (rm I2SPR_Mask) Load() I2SPR_Bits   { return I2SPR_Bits(rm.UM32.Load()) }
func (rm I2SPR_Mask) Store(b I2SPR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *SPI_Periph) I2SDIV() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM32{&p.I2SPR.U32, uint32(I2SDIV)}}
}

func (p *SPI_Periph) ODD() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM32{&p.I2SPR.U32, uint32(ODD)}}
}

func (p *SPI_Periph) MCKOE() I2SPR_Mask {
	return I2SPR_Mask{mmio.UM32{&p.I2SPR.U32, uint32(MCKOE)}}
}
