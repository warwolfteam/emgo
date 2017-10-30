package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type FLASH_Periph struct {
	ACR     ACR
	PECR    PECR
	PDKEYR  PDKEYR
	PEKEYR  PEKEYR
	PRGKEYR PRGKEYR
	OPTKEYR OPTKEYR
	SR      SR
	OBR     OBR
	WRPR    WRPR
	_       [23]uint32
	WRPR1   WRPR1
	WRPR2   WRPR2
	WRPR3   WRPR3
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR_Bits uint32

func (b ACR_Bits) Field(mask ACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR_Bits) J(v int) ACR_Bits {
	return ACR_Bits(bits.Make32(v, uint32(mask)))
}

type ACR struct{ mmio.U32 }

func (r *ACR) Bits(mask ACR_Bits) ACR_Bits { return ACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ACR) StoreBits(mask, b ACR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ACR) SetBits(mask ACR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ACR) ClearBits(mask ACR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ACR) Load() ACR_Bits              { return ACR_Bits(r.U32.Load()) }
func (r *ACR) Store(b ACR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ACR) AtomicStoreBits(mask, b ACR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ACR) AtomicSetBits(mask ACR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ACR) AtomicClearBits(mask ACR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ACR_Mask struct{ mmio.UM32 }

func (rm ACR_Mask) Load() ACR_Bits   { return ACR_Bits(rm.UM32.Load()) }
func (rm ACR_Mask) Store(b ACR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func (p *FLASH_Periph) PRFTEN() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func (p *FLASH_Periph) ACC64() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(ACC64)}}
}

func (p *FLASH_Periph) SLEEP_PD() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(SLEEP_PD)}}
}

func (p *FLASH_Periph) RUN_PD() ACR_Mask {
	return ACR_Mask{mmio.UM32{&p.ACR.U32, uint32(RUN_PD)}}
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

func (p *FLASH_Periph) PELOCK() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(PELOCK)}}
}

func (p *FLASH_Periph) PRGLOCK() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(PRGLOCK)}}
}

func (p *FLASH_Periph) OPTLOCK() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(OPTLOCK)}}
}

func (p *FLASH_Periph) PROG() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(PROG)}}
}

func (p *FLASH_Periph) DATA() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(DATA)}}
}

func (p *FLASH_Periph) FTDW() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(FTDW)}}
}

func (p *FLASH_Periph) ERASE() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(ERASE)}}
}

func (p *FLASH_Periph) FPRG() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(FPRG)}}
}

func (p *FLASH_Periph) PARALLBANK() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(PARALLBANK)}}
}

func (p *FLASH_Periph) EOPIE() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) ERRIE() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) OBL_LAUNCH() PECR_Mask {
	return PECR_Mask{mmio.UM32{&p.PECR.U32, uint32(OBL_LAUNCH)}}
}

type PDKEYR_Bits uint32

func (b PDKEYR_Bits) Field(mask PDKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PDKEYR_Bits) J(v int) PDKEYR_Bits {
	return PDKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type PDKEYR struct{ mmio.U32 }

func (r *PDKEYR) Bits(mask PDKEYR_Bits) PDKEYR_Bits { return PDKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PDKEYR) StoreBits(mask, b PDKEYR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PDKEYR) SetBits(mask PDKEYR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *PDKEYR) ClearBits(mask PDKEYR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *PDKEYR) Load() PDKEYR_Bits                 { return PDKEYR_Bits(r.U32.Load()) }
func (r *PDKEYR) Store(b PDKEYR_Bits)               { r.U32.Store(uint32(b)) }

func (r *PDKEYR) AtomicStoreBits(mask, b PDKEYR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PDKEYR) AtomicSetBits(mask PDKEYR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PDKEYR) AtomicClearBits(mask PDKEYR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PDKEYR_Mask struct{ mmio.UM32 }

func (rm PDKEYR_Mask) Load() PDKEYR_Bits   { return PDKEYR_Bits(rm.UM32.Load()) }
func (rm PDKEYR_Mask) Store(b PDKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type PEKEYR_Bits uint32

func (b PEKEYR_Bits) Field(mask PEKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PEKEYR_Bits) J(v int) PEKEYR_Bits {
	return PEKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type PEKEYR struct{ mmio.U32 }

func (r *PEKEYR) Bits(mask PEKEYR_Bits) PEKEYR_Bits { return PEKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PEKEYR) StoreBits(mask, b PEKEYR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PEKEYR) SetBits(mask PEKEYR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *PEKEYR) ClearBits(mask PEKEYR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *PEKEYR) Load() PEKEYR_Bits                 { return PEKEYR_Bits(r.U32.Load()) }
func (r *PEKEYR) Store(b PEKEYR_Bits)               { r.U32.Store(uint32(b)) }

func (r *PEKEYR) AtomicStoreBits(mask, b PEKEYR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PEKEYR) AtomicSetBits(mask PEKEYR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PEKEYR) AtomicClearBits(mask PEKEYR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PEKEYR_Mask struct{ mmio.UM32 }

func (rm PEKEYR_Mask) Load() PEKEYR_Bits   { return PEKEYR_Bits(rm.UM32.Load()) }
func (rm PEKEYR_Mask) Store(b PEKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type PRGKEYR_Bits uint32

func (b PRGKEYR_Bits) Field(mask PRGKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRGKEYR_Bits) J(v int) PRGKEYR_Bits {
	return PRGKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type PRGKEYR struct{ mmio.U32 }

func (r *PRGKEYR) Bits(mask PRGKEYR_Bits) PRGKEYR_Bits { return PRGKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PRGKEYR) StoreBits(mask, b PRGKEYR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PRGKEYR) SetBits(mask PRGKEYR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *PRGKEYR) ClearBits(mask PRGKEYR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *PRGKEYR) Load() PRGKEYR_Bits                  { return PRGKEYR_Bits(r.U32.Load()) }
func (r *PRGKEYR) Store(b PRGKEYR_Bits)                { r.U32.Store(uint32(b)) }

func (r *PRGKEYR) AtomicStoreBits(mask, b PRGKEYR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *PRGKEYR) AtomicSetBits(mask PRGKEYR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PRGKEYR) AtomicClearBits(mask PRGKEYR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type PRGKEYR_Mask struct{ mmio.UM32 }

func (rm PRGKEYR_Mask) Load() PRGKEYR_Bits   { return PRGKEYR_Bits(rm.UM32.Load()) }
func (rm PRGKEYR_Mask) Store(b PRGKEYR_Bits) { rm.UM32.Store(uint32(b)) }

type OPTKEYR_Bits uint32

func (b OPTKEYR_Bits) Field(mask OPTKEYR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR_Bits) J(v int) OPTKEYR_Bits {
	return OPTKEYR_Bits(bits.Make32(v, uint32(mask)))
}

type OPTKEYR struct{ mmio.U32 }

func (r *OPTKEYR) Bits(mask OPTKEYR_Bits) OPTKEYR_Bits { return OPTKEYR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OPTKEYR) StoreBits(mask, b OPTKEYR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OPTKEYR) SetBits(mask OPTKEYR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OPTKEYR) ClearBits(mask OPTKEYR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OPTKEYR) Load() OPTKEYR_Bits                  { return OPTKEYR_Bits(r.U32.Load()) }
func (r *OPTKEYR) Store(b OPTKEYR_Bits)                { r.U32.Store(uint32(b)) }

func (r *OPTKEYR) AtomicStoreBits(mask, b OPTKEYR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *OPTKEYR) AtomicSetBits(mask OPTKEYR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OPTKEYR) AtomicClearBits(mask OPTKEYR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OPTKEYR_Mask struct{ mmio.UM32 }

func (rm OPTKEYR_Mask) Load() OPTKEYR_Bits   { return OPTKEYR_Bits(rm.UM32.Load()) }
func (rm OPTKEYR_Mask) Store(b OPTKEYR_Bits) { rm.UM32.Store(uint32(b)) }

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

func (p *FLASH_Periph) BSY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

func (p *FLASH_Periph) EOP() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func (p *FLASH_Periph) ENHV() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(ENHV)}}
}

func (p *FLASH_Periph) READY() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(READY)}}
}

func (p *FLASH_Periph) WRPERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) PGAERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func (p *FLASH_Periph) SIZERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(SIZERR)}}
}

func (p *FLASH_Periph) OPTVERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OPTVERR)}}
}

func (p *FLASH_Periph) OPTVERRUSR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(OPTVERRUSR)}}
}

func (p *FLASH_Periph) RDERR() SR_Mask {
	return SR_Mask{mmio.UM32{&p.SR.U32, uint32(RDERR)}}
}

type OBR_Bits uint32

func (b OBR_Bits) Field(mask OBR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OBR_Bits) J(v int) OBR_Bits {
	return OBR_Bits(bits.Make32(v, uint32(mask)))
}

type OBR struct{ mmio.U32 }

func (r *OBR) Bits(mask OBR_Bits) OBR_Bits { return OBR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OBR) StoreBits(mask, b OBR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OBR) SetBits(mask OBR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OBR) ClearBits(mask OBR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OBR) Load() OBR_Bits              { return OBR_Bits(r.U32.Load()) }
func (r *OBR) Store(b OBR_Bits)            { r.U32.Store(uint32(b)) }

func (r *OBR) AtomicStoreBits(mask, b OBR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OBR) AtomicSetBits(mask OBR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OBR) AtomicClearBits(mask OBR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OBR_Mask struct{ mmio.UM32 }

func (rm OBR_Mask) Load() OBR_Bits   { return OBR_Bits(rm.UM32.Load()) }
func (rm OBR_Mask) Store(b OBR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) RDPRT() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(RDPRT)}}
}

func (p *FLASH_Periph) SPRMOD() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(SPRMOD)}}
}

func (p *FLASH_Periph) BOR_LEV() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(BOR_LEV)}}
}

func (p *FLASH_Periph) IWDG_SW() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(IWDG_SW)}}
}

func (p *FLASH_Periph) nRST_STOP() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(nRST_STOP)}}
}

func (p *FLASH_Periph) nRST_STDBY() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(nRST_STDBY)}}
}

func (p *FLASH_Periph) BFB2() OBR_Mask {
	return OBR_Mask{mmio.UM32{&p.OBR.U32, uint32(BFB2)}}
}

type WRPR_Bits uint32

func (b WRPR_Bits) Field(mask WRPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR_Bits) J(v int) WRPR_Bits {
	return WRPR_Bits(bits.Make32(v, uint32(mask)))
}

type WRPR struct{ mmio.U32 }

func (r *WRPR) Bits(mask WRPR_Bits) WRPR_Bits { return WRPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRPR) StoreBits(mask, b WRPR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRPR) SetBits(mask WRPR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *WRPR) ClearBits(mask WRPR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *WRPR) Load() WRPR_Bits               { return WRPR_Bits(r.U32.Load()) }
func (r *WRPR) Store(b WRPR_Bits)             { r.U32.Store(uint32(b)) }

func (r *WRPR) AtomicStoreBits(mask, b WRPR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRPR) AtomicSetBits(mask WRPR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRPR) AtomicClearBits(mask WRPR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRPR_Mask struct{ mmio.UM32 }

func (rm WRPR_Mask) Load() WRPR_Bits   { return WRPR_Bits(rm.UM32.Load()) }
func (rm WRPR_Mask) Store(b WRPR_Bits) { rm.UM32.Store(uint32(b)) }

type WRPR1_Bits uint32

func (b WRPR1_Bits) Field(mask WRPR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR1_Bits) J(v int) WRPR1_Bits {
	return WRPR1_Bits(bits.Make32(v, uint32(mask)))
}

type WRPR1 struct{ mmio.U32 }

func (r *WRPR1) Bits(mask WRPR1_Bits) WRPR1_Bits { return WRPR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRPR1) StoreBits(mask, b WRPR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRPR1) SetBits(mask WRPR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRPR1) ClearBits(mask WRPR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRPR1) Load() WRPR1_Bits                { return WRPR1_Bits(r.U32.Load()) }
func (r *WRPR1) Store(b WRPR1_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRPR1) AtomicStoreBits(mask, b WRPR1_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRPR1) AtomicSetBits(mask WRPR1_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRPR1) AtomicClearBits(mask WRPR1_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRPR1_Mask struct{ mmio.UM32 }

func (rm WRPR1_Mask) Load() WRPR1_Bits   { return WRPR1_Bits(rm.UM32.Load()) }
func (rm WRPR1_Mask) Store(b WRPR1_Bits) { rm.UM32.Store(uint32(b)) }

type WRPR2_Bits uint32

func (b WRPR2_Bits) Field(mask WRPR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR2_Bits) J(v int) WRPR2_Bits {
	return WRPR2_Bits(bits.Make32(v, uint32(mask)))
}

type WRPR2 struct{ mmio.U32 }

func (r *WRPR2) Bits(mask WRPR2_Bits) WRPR2_Bits { return WRPR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRPR2) StoreBits(mask, b WRPR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRPR2) SetBits(mask WRPR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRPR2) ClearBits(mask WRPR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRPR2) Load() WRPR2_Bits                { return WRPR2_Bits(r.U32.Load()) }
func (r *WRPR2) Store(b WRPR2_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRPR2) AtomicStoreBits(mask, b WRPR2_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRPR2) AtomicSetBits(mask WRPR2_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRPR2) AtomicClearBits(mask WRPR2_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRPR2_Mask struct{ mmio.UM32 }

func (rm WRPR2_Mask) Load() WRPR2_Bits   { return WRPR2_Bits(rm.UM32.Load()) }
func (rm WRPR2_Mask) Store(b WRPR2_Bits) { rm.UM32.Store(uint32(b)) }

type WRPR3_Bits uint32

func (b WRPR3_Bits) Field(mask WRPR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR3_Bits) J(v int) WRPR3_Bits {
	return WRPR3_Bits(bits.Make32(v, uint32(mask)))
}

type WRPR3 struct{ mmio.U32 }

func (r *WRPR3) Bits(mask WRPR3_Bits) WRPR3_Bits { return WRPR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *WRPR3) StoreBits(mask, b WRPR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WRPR3) SetBits(mask WRPR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WRPR3) ClearBits(mask WRPR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WRPR3) Load() WRPR3_Bits                { return WRPR3_Bits(r.U32.Load()) }
func (r *WRPR3) Store(b WRPR3_Bits)              { r.U32.Store(uint32(b)) }

func (r *WRPR3) AtomicStoreBits(mask, b WRPR3_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WRPR3) AtomicSetBits(mask WRPR3_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WRPR3) AtomicClearBits(mask WRPR3_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WRPR3_Mask struct{ mmio.UM32 }

func (rm WRPR3_Mask) Load() WRPR3_Bits   { return WRPR3_Bits(rm.UM32.Load()) }
func (rm WRPR3_Mask) Store(b WRPR3_Bits) { rm.UM32.Store(uint32(b)) }
