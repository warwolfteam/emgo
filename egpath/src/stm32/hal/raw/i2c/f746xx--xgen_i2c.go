// +build f746xx

package i2c

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type I2C_Periph struct {
	CR1      CR1
	CR2      CR2
	OAR1     OAR1
	OAR2     OAR2
	TIMINGR  TIMINGR
	TIMEOUTR TIMEOUTR
	ISR      ISR
	ICR      ICR
	PECR     PECR
	RXDR     RXDR
	TXDR     TXDR
}

func (p *I2C_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var I2C1 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C1_BASE)))

//emgo:const
var I2C2 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C2_BASE)))

//emgo:const
var I2C3 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C3_BASE)))

//emgo:const
var I2C4 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C4_BASE)))

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

func (p *I2C_Periph) PE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PE)}}
}

func (p *I2C_Periph) TXIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TXIE)}}
}

func (p *I2C_Periph) RXIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RXIE)}}
}

func (p *I2C_Periph) ADDRIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ADDRIE)}}
}

func (p *I2C_Periph) NACKIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(NACKIE)}}
}

func (p *I2C_Periph) STOPIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(STOPIE)}}
}

func (p *I2C_Periph) TCIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TCIE)}}
}

func (p *I2C_Periph) ERRIE() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ERRIE)}}
}

func (p *I2C_Periph) DNF() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(DNF)}}
}

func (p *I2C_Periph) ANFOFF() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ANFOFF)}}
}

func (p *I2C_Periph) TXDMAEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(TXDMAEN)}}
}

func (p *I2C_Periph) RXDMAEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(RXDMAEN)}}
}

func (p *I2C_Periph) SBC() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SBC)}}
}

func (p *I2C_Periph) NOSTRETCH() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(NOSTRETCH)}}
}

func (p *I2C_Periph) GCEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(GCEN)}}
}

func (p *I2C_Periph) SMBHEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SMBHEN)}}
}

func (p *I2C_Periph) SMBDEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(SMBDEN)}}
}

func (p *I2C_Periph) ALERTEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(ALERTEN)}}
}

func (p *I2C_Periph) PECEN() CR1_Mask {
	return CR1_Mask{mmio.UM32{&p.CR1.U32, uint32(PECEN)}}
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

func (p *I2C_Periph) SADD() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(SADD)}}
}

func (p *I2C_Periph) RD_WRN() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RD_WRN)}}
}

func (p *I2C_Periph) ADD10() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(ADD10)}}
}

func (p *I2C_Periph) HEAD10R() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(HEAD10R)}}
}

func (p *I2C_Periph) START() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(START)}}
}

func (p *I2C_Periph) STOP() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(STOP)}}
}

func (p *I2C_Periph) NACK() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(NACK)}}
}

func (p *I2C_Periph) NBYTES() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(NBYTES)}}
}

func (p *I2C_Periph) RELOAD() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(RELOAD)}}
}

func (p *I2C_Periph) AUTOEND() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(AUTOEND)}}
}

func (p *I2C_Periph) PECBYTE() CR2_Mask {
	return CR2_Mask{mmio.UM32{&p.CR2.U32, uint32(PECBYTE)}}
}

type OAR1_Bits uint32

func (b OAR1_Bits) Field(mask OAR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR1_Bits) J(v int) OAR1_Bits {
	return OAR1_Bits(bits.Make32(v, uint32(mask)))
}

type OAR1 struct{ mmio.U32 }

func (r *OAR1) Bits(mask OAR1_Bits) OAR1_Bits { return OAR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *OAR1) StoreBits(mask, b OAR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OAR1) SetBits(mask OAR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *OAR1) ClearBits(mask OAR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *OAR1) Load() OAR1_Bits               { return OAR1_Bits(r.U32.Load()) }
func (r *OAR1) Store(b OAR1_Bits)             { r.U32.Store(uint32(b)) }

func (r *OAR1) AtomicStoreBits(mask, b OAR1_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OAR1) AtomicSetBits(mask OAR1_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OAR1) AtomicClearBits(mask OAR1_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OAR1_Mask struct{ mmio.UM32 }

func (rm OAR1_Mask) Load() OAR1_Bits   { return OAR1_Bits(rm.UM32.Load()) }
func (rm OAR1_Mask) Store(b OAR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) OA1() OAR1_Mask {
	return OAR1_Mask{mmio.UM32{&p.OAR1.U32, uint32(OA1)}}
}

func (p *I2C_Periph) OA1MODE() OAR1_Mask {
	return OAR1_Mask{mmio.UM32{&p.OAR1.U32, uint32(OA1MODE)}}
}

func (p *I2C_Periph) OA1EN() OAR1_Mask {
	return OAR1_Mask{mmio.UM32{&p.OAR1.U32, uint32(OA1EN)}}
}

type OAR2_Bits uint32

func (b OAR2_Bits) Field(mask OAR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR2_Bits) J(v int) OAR2_Bits {
	return OAR2_Bits(bits.Make32(v, uint32(mask)))
}

type OAR2 struct{ mmio.U32 }

func (r *OAR2) Bits(mask OAR2_Bits) OAR2_Bits { return OAR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *OAR2) StoreBits(mask, b OAR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OAR2) SetBits(mask OAR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *OAR2) ClearBits(mask OAR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *OAR2) Load() OAR2_Bits               { return OAR2_Bits(r.U32.Load()) }
func (r *OAR2) Store(b OAR2_Bits)             { r.U32.Store(uint32(b)) }

func (r *OAR2) AtomicStoreBits(mask, b OAR2_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OAR2) AtomicSetBits(mask OAR2_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OAR2) AtomicClearBits(mask OAR2_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OAR2_Mask struct{ mmio.UM32 }

func (rm OAR2_Mask) Load() OAR2_Bits   { return OAR2_Bits(rm.UM32.Load()) }
func (rm OAR2_Mask) Store(b OAR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) OA2() OAR2_Mask {
	return OAR2_Mask{mmio.UM32{&p.OAR2.U32, uint32(OA2)}}
}

func (p *I2C_Periph) OA2MSK() OAR2_Mask {
	return OAR2_Mask{mmio.UM32{&p.OAR2.U32, uint32(OA2MSK)}}
}

func (p *I2C_Periph) OA2EN() OAR2_Mask {
	return OAR2_Mask{mmio.UM32{&p.OAR2.U32, uint32(OA2EN)}}
}

type TIMINGR_Bits uint32

func (b TIMINGR_Bits) Field(mask TIMINGR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMINGR_Bits) J(v int) TIMINGR_Bits {
	return TIMINGR_Bits(bits.Make32(v, uint32(mask)))
}

type TIMINGR struct{ mmio.U32 }

func (r *TIMINGR) Bits(mask TIMINGR_Bits) TIMINGR_Bits { return TIMINGR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TIMINGR) StoreBits(mask, b TIMINGR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMINGR) SetBits(mask TIMINGR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *TIMINGR) ClearBits(mask TIMINGR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *TIMINGR) Load() TIMINGR_Bits                  { return TIMINGR_Bits(r.U32.Load()) }
func (r *TIMINGR) Store(b TIMINGR_Bits)                { r.U32.Store(uint32(b)) }

func (r *TIMINGR) AtomicStoreBits(mask, b TIMINGR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *TIMINGR) AtomicSetBits(mask TIMINGR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TIMINGR) AtomicClearBits(mask TIMINGR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type TIMINGR_Mask struct{ mmio.UM32 }

func (rm TIMINGR_Mask) Load() TIMINGR_Bits   { return TIMINGR_Bits(rm.UM32.Load()) }
func (rm TIMINGR_Mask) Store(b TIMINGR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) SCLL() TIMINGR_Mask {
	return TIMINGR_Mask{mmio.UM32{&p.TIMINGR.U32, uint32(SCLL)}}
}

func (p *I2C_Periph) SCLH() TIMINGR_Mask {
	return TIMINGR_Mask{mmio.UM32{&p.TIMINGR.U32, uint32(SCLH)}}
}

func (p *I2C_Periph) SDADEL() TIMINGR_Mask {
	return TIMINGR_Mask{mmio.UM32{&p.TIMINGR.U32, uint32(SDADEL)}}
}

func (p *I2C_Periph) SCLDEL() TIMINGR_Mask {
	return TIMINGR_Mask{mmio.UM32{&p.TIMINGR.U32, uint32(SCLDEL)}}
}

func (p *I2C_Periph) PRESC() TIMINGR_Mask {
	return TIMINGR_Mask{mmio.UM32{&p.TIMINGR.U32, uint32(PRESC)}}
}

type TIMEOUTR_Bits uint32

func (b TIMEOUTR_Bits) Field(mask TIMEOUTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMEOUTR_Bits) J(v int) TIMEOUTR_Bits {
	return TIMEOUTR_Bits(bits.Make32(v, uint32(mask)))
}

type TIMEOUTR struct{ mmio.U32 }

func (r *TIMEOUTR) Bits(mask TIMEOUTR_Bits) TIMEOUTR_Bits {
	return TIMEOUTR_Bits(r.U32.Bits(uint32(mask)))
}
func (r *TIMEOUTR) StoreBits(mask, b TIMEOUTR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TIMEOUTR) SetBits(mask TIMEOUTR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *TIMEOUTR) ClearBits(mask TIMEOUTR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *TIMEOUTR) Load() TIMEOUTR_Bits             { return TIMEOUTR_Bits(r.U32.Load()) }
func (r *TIMEOUTR) Store(b TIMEOUTR_Bits)           { r.U32.Store(uint32(b)) }

func (r *TIMEOUTR) AtomicStoreBits(mask, b TIMEOUTR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *TIMEOUTR) AtomicSetBits(mask TIMEOUTR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TIMEOUTR) AtomicClearBits(mask TIMEOUTR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type TIMEOUTR_Mask struct{ mmio.UM32 }

func (rm TIMEOUTR_Mask) Load() TIMEOUTR_Bits   { return TIMEOUTR_Bits(rm.UM32.Load()) }
func (rm TIMEOUTR_Mask) Store(b TIMEOUTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TIMEOUTA() TIMEOUTR_Mask {
	return TIMEOUTR_Mask{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMEOUTA)}}
}

func (p *I2C_Periph) TIDLE() TIMEOUTR_Mask {
	return TIMEOUTR_Mask{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIDLE)}}
}

func (p *I2C_Periph) TIMOUTEN() TIMEOUTR_Mask {
	return TIMEOUTR_Mask{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMOUTEN)}}
}

func (p *I2C_Periph) TIMEOUTB() TIMEOUTR_Mask {
	return TIMEOUTR_Mask{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMEOUTB)}}
}

func (p *I2C_Periph) TEXTEN() TIMEOUTR_Mask {
	return TIMEOUTR_Mask{mmio.UM32{&p.TIMEOUTR.U32, uint32(TEXTEN)}}
}

type ISR_Bits uint32

func (b ISR_Bits) Field(mask ISR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR_Bits) J(v int) ISR_Bits {
	return ISR_Bits(bits.Make32(v, uint32(mask)))
}

type ISR struct{ mmio.U32 }

func (r *ISR) Bits(mask ISR_Bits) ISR_Bits { return ISR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ISR) StoreBits(mask, b ISR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ISR) SetBits(mask ISR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ISR) ClearBits(mask ISR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ISR) Load() ISR_Bits              { return ISR_Bits(r.U32.Load()) }
func (r *ISR) Store(b ISR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ISR) AtomicStoreBits(mask, b ISR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ISR) AtomicSetBits(mask ISR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ISR) AtomicClearBits(mask ISR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ISR_Mask struct{ mmio.UM32 }

func (rm ISR_Mask) Load() ISR_Bits   { return ISR_Bits(rm.UM32.Load()) }
func (rm ISR_Mask) Store(b ISR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TXE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TXE)}}
}

func (p *I2C_Periph) TXIS() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TXIS)}}
}

func (p *I2C_Periph) RXNE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RXNE)}}
}

func (p *I2C_Periph) ADDR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ADDR)}}
}

func (p *I2C_Periph) NACKF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(NACKF)}}
}

func (p *I2C_Periph) STOPF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(STOPF)}}
}

func (p *I2C_Periph) TC() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TC)}}
}

func (p *I2C_Periph) TCR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TCR)}}
}

func (p *I2C_Periph) BERR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(BERR)}}
}

func (p *I2C_Periph) ARLO() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ARLO)}}
}

func (p *I2C_Periph) OVR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(OVR)}}
}

func (p *I2C_Periph) PECERR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(PECERR)}}
}

func (p *I2C_Periph) TIMEOUT() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TIMEOUT)}}
}

func (p *I2C_Periph) ALERT() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ALERT)}}
}

func (p *I2C_Periph) BUSY() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(BUSY)}}
}

func (p *I2C_Periph) DIR() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(DIR)}}
}

func (p *I2C_Periph) ADDCODE() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ADDCODE)}}
}

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ICR) AtomicStoreBits(mask, b ICR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ICR) AtomicSetBits(mask ICR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ICR) AtomicClearBits(mask ICR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) ADDRCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ADDRCF)}}
}

func (p *I2C_Periph) NACKCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(NACKCF)}}
}

func (p *I2C_Periph) STOPCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(STOPCF)}}
}

func (p *I2C_Periph) BERRCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(BERRCF)}}
}

func (p *I2C_Periph) ARLOCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ARLOCF)}}
}

func (p *I2C_Periph) OVRCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(OVRCF)}}
}

func (p *I2C_Periph) PECCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(PECCF)}}
}

func (p *I2C_Periph) TIMOUTCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(TIMOUTCF)}}
}

func (p *I2C_Periph) ALERTCF() ICR_Mask {
	return ICR_Mask{mmio.UM32{&p.ICR.U32, uint32(ALERTCF)}}
}

type PECR_Bits uint32

func (b PECR_Bits) Field(mask PECR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PECR_Bits) J(v int) PECR_Bits {
	return PECR_Bits(bits.Make32(v, uint32(mask)))
}

type PECR struct{ mmio.U32 }

func (r *PECR) Bits(mask PECR_Bits) PECR_Bits { return PECR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PECR) StoreBits(mask, b PECR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PECR) SetBits(mask PECR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PECR) ClearBits(mask PECR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PECR) Load() PECR_Bits               { return PECR_Bits(r.U32.Load()) }
func (r *PECR) Store(b PECR_Bits)             { r.U32.Store(uint32(b)) }

func (r *PECR) AtomicStoreBits(mask, b PECR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PECR) AtomicSetBits(mask PECR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PECR) AtomicClearBits(mask PECR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PECR_Mask struct{ mmio.UM32 }

func (rm PECR_Mask) Load() PECR_Bits   { return PECR_Bits(rm.UM32.Load()) }
func (rm PECR_Mask) Store(b PECR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) PEC() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(PEC)}}
}

type RXDR_Bits uint32

func (b RXDR_Bits) Field(mask RXDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXDR_Bits) J(v int) RXDR_Bits {
	return RXDR_Bits(bits.Make32(v, uint32(mask)))
}

type RXDR struct{ mmio.U32 }

func (r *RXDR) Bits(mask RXDR_Bits) RXDR_Bits { return RXDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *RXDR) StoreBits(mask, b RXDR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RXDR) SetBits(mask RXDR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RXDR) ClearBits(mask RXDR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RXDR) Load() RXDR_Bits               { return RXDR_Bits(r.U32.Load()) }
func (r *RXDR) Store(b RXDR_Bits)             { r.U32.Store(uint32(b)) }

func (r *RXDR) AtomicStoreBits(mask, b RXDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RXDR) AtomicSetBits(mask RXDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RXDR) AtomicClearBits(mask RXDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type RXDR_Mask struct{ mmio.UM32 }

func (rm RXDR_Mask) Load() RXDR_Bits   { return RXDR_Bits(rm.UM32.Load()) }
func (rm RXDR_Mask) Store(b RXDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) RXDATA() RXDR_Mask {
	return RXDR_Mask{mmio.UM32{&p.RXDR.U32, uint32(RXDATA)}}
}

type TXDR_Bits uint32

func (b TXDR_Bits) Field(mask TXDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXDR_Bits) J(v int) TXDR_Bits {
	return TXDR_Bits(bits.Make32(v, uint32(mask)))
}

type TXDR struct{ mmio.U32 }

func (r *TXDR) Bits(mask TXDR_Bits) TXDR_Bits { return TXDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TXDR) StoreBits(mask, b TXDR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TXDR) SetBits(mask TXDR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TXDR) ClearBits(mask TXDR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TXDR) Load() TXDR_Bits               { return TXDR_Bits(r.U32.Load()) }
func (r *TXDR) Store(b TXDR_Bits)             { r.U32.Store(uint32(b)) }

func (r *TXDR) AtomicStoreBits(mask, b TXDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TXDR) AtomicSetBits(mask TXDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TXDR) AtomicClearBits(mask TXDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TXDR_Mask struct{ mmio.UM32 }

func (rm TXDR_Mask) Load() TXDR_Bits   { return TXDR_Bits(rm.UM32.Load()) }
func (rm TXDR_Mask) Store(b TXDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TXDATA() TXDR_Mask {
	return TXDR_Mask{mmio.UM32{&p.TXDR.U32, uint32(TXDATA)}}
}
