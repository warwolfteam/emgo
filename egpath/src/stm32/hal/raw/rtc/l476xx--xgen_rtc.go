// +build l476xx

package rtc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type RTC_Periph struct {
	TR      TR
	DR      DR
	CR      CR
	ISR     ISR
	PRER    PRER
	WUTR    WUTR
	_       uint32
	ALRMR   [2]ALRMR
	WPR     WPR
	SSR     SSR
	SHIFTR  SHIFTR
	TSTR    TSTR
	TSDR    TSDR
	TSSSR   TSSSR
	CALR    CALR
	TAMPCR  TAMPCR
	ALRMSSR [2]ALRMSSR
	OR      OR
	BKPR    [32]BKPR
}

func (p *RTC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RTC = (*RTC_Periph)(unsafe.Pointer(uintptr(mmap.RTC_BASE)))

type TR_Bits uint32

func (b TR_Bits) Field(mask TR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TR_Bits) J(v int) TR_Bits {
	return TR_Bits(bits.Make32(v, uint32(mask)))
}

type TR struct{ mmio.U32 }

func (r *TR) Bits(mask TR_Bits) TR_Bits { return TR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TR) StoreBits(mask, b TR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TR) SetBits(mask TR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *TR) ClearBits(mask TR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *TR) Load() TR_Bits             { return TR_Bits(r.U32.Load()) }
func (r *TR) Store(b TR_Bits)           { r.U32.Store(uint32(b)) }

func (r *TR) AtomicStoreBits(mask, b TR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TR) AtomicSetBits(mask TR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TR) AtomicClearBits(mask TR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TR_Mask struct{ mmio.UM32 }

func (rm TR_Mask) Load() TR_Bits   { return TR_Bits(rm.UM32.Load()) }
func (rm TR_Mask) Store(b TR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PM() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(PM)}}
}

func (p *RTC_Periph) HT() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(HT)}}
}

func (p *RTC_Periph) HU() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(HU)}}
}

func (p *RTC_Periph) MNT() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(MNT)}}
}

func (p *RTC_Periph) MNU() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(MNU)}}
}

func (p *RTC_Periph) ST() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(ST)}}
}

func (p *RTC_Periph) SU() TR_Mask {
	return TR_Mask{mmio.UM32{&p.TR.U32, uint32(SU)}}
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

func (p *RTC_Periph) YT() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(YT)}}
}

func (p *RTC_Periph) YU() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(YU)}}
}

func (p *RTC_Periph) WDU() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(WDU)}}
}

func (p *RTC_Periph) MT() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(MT)}}
}

func (p *RTC_Periph) MU() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(MU)}}
}

func (p *RTC_Periph) DT() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(DT)}}
}

func (p *RTC_Periph) DU() DR_Mask {
	return DR_Mask{mmio.UM32{&p.DR.U32, uint32(DU)}}
}

type CR_Bits uint32

func (b CR_Bits) Field(mask CR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR_Bits) J(v int) CR_Bits {
	return CR_Bits(bits.Make32(v, uint32(mask)))
}

type CR struct{ mmio.U32 }

func (r *CR) Bits(mask CR_Bits) CR_Bits { return CR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR) StoreBits(mask, b CR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR) SetBits(mask CR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *CR) ClearBits(mask CR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *CR) Load() CR_Bits             { return CR_Bits(r.U32.Load()) }
func (r *CR) Store(b CR_Bits)           { r.U32.Store(uint32(b)) }

func (r *CR) AtomicStoreBits(mask, b CR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CR) AtomicSetBits(mask CR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CR) AtomicClearBits(mask CR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CR_Mask struct{ mmio.UM32 }

func (rm CR_Mask) Load() CR_Bits   { return CR_Bits(rm.UM32.Load()) }
func (rm CR_Mask) Store(b CR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) ITSE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ITSE)}}
}

func (p *RTC_Periph) COE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(COE)}}
}

func (p *RTC_Periph) OSEL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(OSEL)}}
}

func (p *RTC_Periph) POL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(POL)}}
}

func (p *RTC_Periph) COSEL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(COSEL)}}
}

func (p *RTC_Periph) BCK() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BCK)}}
}

func (p *RTC_Periph) SUB1H() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(SUB1H)}}
}

func (p *RTC_Periph) ADD1H() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ADD1H)}}
}

func (p *RTC_Periph) TSIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TSIE)}}
}

func (p *RTC_Periph) WUTIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WUTIE)}}
}

func (p *RTC_Periph) ALRBIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ALRBIE)}}
}

func (p *RTC_Periph) ALRAIE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ALRAIE)}}
}

func (p *RTC_Periph) TSE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TSE)}}
}

func (p *RTC_Periph) WUTE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WUTE)}}
}

func (p *RTC_Periph) ALRBE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ALRBE)}}
}

func (p *RTC_Periph) ALRAE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(ALRAE)}}
}

func (p *RTC_Periph) FMT() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(FMT)}}
}

func (p *RTC_Periph) BYPSHAD() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(BYPSHAD)}}
}

func (p *RTC_Periph) REFCKON() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(REFCKON)}}
}

func (p *RTC_Periph) TSEDGE() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(TSEDGE)}}
}

func (p *RTC_Periph) WUCKSEL() CR_Mask {
	return CR_Mask{mmio.UM32{&p.CR.U32, uint32(WUCKSEL)}}
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

func (p *RTC_Periph) ITSF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ITSF)}}
}

func (p *RTC_Periph) RECALPF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RECALPF)}}
}

func (p *RTC_Periph) TAMP3F() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TAMP3F)}}
}

func (p *RTC_Periph) TAMP2F() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TAMP2F)}}
}

func (p *RTC_Periph) TAMP1F() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TAMP1F)}}
}

func (p *RTC_Periph) TSOVF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TSOVF)}}
}

func (p *RTC_Periph) TSF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(TSF)}}
}

func (p *RTC_Periph) WUTF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(WUTF)}}
}

func (p *RTC_Periph) ALRBF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ALRBF)}}
}

func (p *RTC_Periph) ALRAF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ALRAF)}}
}

func (p *RTC_Periph) INIT() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(INIT)}}
}

func (p *RTC_Periph) INITF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(INITF)}}
}

func (p *RTC_Periph) RSF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(RSF)}}
}

func (p *RTC_Periph) INITS() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(INITS)}}
}

func (p *RTC_Periph) SHPF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(SHPF)}}
}

func (p *RTC_Periph) WUTWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(WUTWF)}}
}

func (p *RTC_Periph) ALRBWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ALRBWF)}}
}

func (p *RTC_Periph) ALRAWF() ISR_Mask {
	return ISR_Mask{mmio.UM32{&p.ISR.U32, uint32(ALRAWF)}}
}

type PRER_Bits uint32

func (b PRER_Bits) Field(mask PRER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PRER_Bits) J(v int) PRER_Bits {
	return PRER_Bits(bits.Make32(v, uint32(mask)))
}

type PRER struct{ mmio.U32 }

func (r *PRER) Bits(mask PRER_Bits) PRER_Bits { return PRER_Bits(r.U32.Bits(uint32(mask))) }
func (r *PRER) StoreBits(mask, b PRER_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PRER) SetBits(mask PRER_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PRER) ClearBits(mask PRER_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PRER) Load() PRER_Bits               { return PRER_Bits(r.U32.Load()) }
func (r *PRER) Store(b PRER_Bits)             { r.U32.Store(uint32(b)) }

func (r *PRER) AtomicStoreBits(mask, b PRER_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PRER) AtomicSetBits(mask PRER_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PRER) AtomicClearBits(mask PRER_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PRER_Mask struct{ mmio.UM32 }

func (rm PRER_Mask) Load() PRER_Bits   { return PRER_Bits(rm.UM32.Load()) }
func (rm PRER_Mask) Store(b PRER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) PREDIV_A() PRER_Mask {
	return PRER_Mask{mmio.UM32{&p.PRER.U32, uint32(PREDIV_A)}}
}

func (p *RTC_Periph) PREDIV_S() PRER_Mask {
	return PRER_Mask{mmio.UM32{&p.PRER.U32, uint32(PREDIV_S)}}
}

type WUTR_Bits uint32

func (b WUTR_Bits) Field(mask WUTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WUTR_Bits) J(v int) WUTR_Bits {
	return WUTR_Bits(bits.Make32(v, uint32(mask)))
}

type WUTR struct{ mmio.U32 }

func (r *WUTR) Bits(mask WUTR_Bits) WUTR_Bits { return WUTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WUTR) StoreBits(mask, b WUTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WUTR) SetBits(mask WUTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *WUTR) ClearBits(mask WUTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *WUTR) Load() WUTR_Bits               { return WUTR_Bits(r.U32.Load()) }
func (r *WUTR) Store(b WUTR_Bits)             { r.U32.Store(uint32(b)) }

func (r *WUTR) AtomicStoreBits(mask, b WUTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WUTR) AtomicSetBits(mask WUTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WUTR) AtomicClearBits(mask WUTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WUTR_Mask struct{ mmio.UM32 }

func (rm WUTR_Mask) Load() WUTR_Bits   { return WUTR_Bits(rm.UM32.Load()) }
func (rm WUTR_Mask) Store(b WUTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) WUT() WUTR_Mask {
	return WUTR_Mask{mmio.UM32{&p.WUTR.U32, uint32(WUT)}}
}

type ALRMR_Bits uint32

func (b ALRMR_Bits) Field(mask ALRMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMR_Bits) J(v int) ALRMR_Bits {
	return ALRMR_Bits(bits.Make32(v, uint32(mask)))
}

type ALRMR struct{ mmio.U32 }

func (r *ALRMR) Bits(mask ALRMR_Bits) ALRMR_Bits { return ALRMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ALRMR) StoreBits(mask, b ALRMR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ALRMR) SetBits(mask ALRMR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ALRMR) ClearBits(mask ALRMR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ALRMR) Load() ALRMR_Bits                { return ALRMR_Bits(r.U32.Load()) }
func (r *ALRMR) Store(b ALRMR_Bits)              { r.U32.Store(uint32(b)) }

func (r *ALRMR) AtomicStoreBits(mask, b ALRMR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ALRMR) AtomicSetBits(mask ALRMR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ALRMR) AtomicClearBits(mask ALRMR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ALRMR_Mask struct{ mmio.UM32 }

func (rm ALRMR_Mask) Load() ALRMR_Bits   { return ALRMR_Bits(rm.UM32.Load()) }
func (rm ALRMR_Mask) Store(b ALRMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) AMSK4(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK4)}}
}

func (p *RTC_Periph) AWDSEL(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AWDSEL)}}
}

func (p *RTC_Periph) ADT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(ADT)}}
}

func (p *RTC_Periph) ADU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(ADU)}}
}

func (p *RTC_Periph) AMSK3(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK3)}}
}

func (p *RTC_Periph) APM(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(APM)}}
}

func (p *RTC_Periph) AHT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AHT)}}
}

func (p *RTC_Periph) AHU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AHU)}}
}

func (p *RTC_Periph) AMSK2(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK2)}}
}

func (p *RTC_Periph) AMNT(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMNT)}}
}

func (p *RTC_Periph) AMNU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMNU)}}
}

func (p *RTC_Periph) AMSK1(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AMSK1)}}
}

func (p *RTC_Periph) AST(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(AST)}}
}

func (p *RTC_Periph) ASU(n int) ALRMR_Mask {
	return ALRMR_Mask{mmio.UM32{&p.ALRMR[n].U32, uint32(ASU)}}
}

type WPR_Bits uint32

func (b WPR_Bits) Field(mask WPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WPR_Bits) J(v int) WPR_Bits {
	return WPR_Bits(bits.Make32(v, uint32(mask)))
}

type WPR struct{ mmio.U32 }

func (r *WPR) Bits(mask WPR_Bits) WPR_Bits { return WPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WPR) StoreBits(mask, b WPR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WPR) SetBits(mask WPR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *WPR) ClearBits(mask WPR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *WPR) Load() WPR_Bits              { return WPR_Bits(r.U32.Load()) }
func (r *WPR) Store(b WPR_Bits)            { r.U32.Store(uint32(b)) }

func (r *WPR) AtomicStoreBits(mask, b WPR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WPR) AtomicSetBits(mask WPR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WPR) AtomicClearBits(mask WPR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WPR_Mask struct{ mmio.UM32 }

func (rm WPR_Mask) Load() WPR_Bits   { return WPR_Bits(rm.UM32.Load()) }
func (rm WPR_Mask) Store(b WPR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) KEY() WPR_Mask {
	return WPR_Mask{mmio.UM32{&p.WPR.U32, uint32(KEY)}}
}

type SSR_Bits uint32

func (b SSR_Bits) Field(mask SSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SSR_Bits) J(v int) SSR_Bits {
	return SSR_Bits(bits.Make32(v, uint32(mask)))
}

type SSR struct{ mmio.U32 }

func (r *SSR) Bits(mask SSR_Bits) SSR_Bits { return SSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SSR) StoreBits(mask, b SSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SSR) SetBits(mask SSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *SSR) ClearBits(mask SSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *SSR) Load() SSR_Bits              { return SSR_Bits(r.U32.Load()) }
func (r *SSR) Store(b SSR_Bits)            { r.U32.Store(uint32(b)) }

func (r *SSR) AtomicStoreBits(mask, b SSR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *SSR) AtomicSetBits(mask SSR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SSR) AtomicClearBits(mask SSR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type SSR_Mask struct{ mmio.UM32 }

func (rm SSR_Mask) Load() SSR_Bits   { return SSR_Bits(rm.UM32.Load()) }
func (rm SSR_Mask) Store(b SSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SS() SSR_Mask {
	return SSR_Mask{mmio.UM32{&p.SSR.U32, uint32(SS)}}
}

type SHIFTR_Bits uint32

func (b SHIFTR_Bits) Field(mask SHIFTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHIFTR_Bits) J(v int) SHIFTR_Bits {
	return SHIFTR_Bits(bits.Make32(v, uint32(mask)))
}

type SHIFTR struct{ mmio.U32 }

func (r *SHIFTR) Bits(mask SHIFTR_Bits) SHIFTR_Bits { return SHIFTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SHIFTR) StoreBits(mask, b SHIFTR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SHIFTR) SetBits(mask SHIFTR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *SHIFTR) ClearBits(mask SHIFTR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *SHIFTR) Load() SHIFTR_Bits                 { return SHIFTR_Bits(r.U32.Load()) }
func (r *SHIFTR) Store(b SHIFTR_Bits)               { r.U32.Store(uint32(b)) }

func (r *SHIFTR) AtomicStoreBits(mask, b SHIFTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *SHIFTR) AtomicSetBits(mask SHIFTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *SHIFTR) AtomicClearBits(mask SHIFTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type SHIFTR_Mask struct{ mmio.UM32 }

func (rm SHIFTR_Mask) Load() SHIFTR_Bits   { return SHIFTR_Bits(rm.UM32.Load()) }
func (rm SHIFTR_Mask) Store(b SHIFTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) SUBFS() SHIFTR_Mask {
	return SHIFTR_Mask{mmio.UM32{&p.SHIFTR.U32, uint32(SUBFS)}}
}

func (p *RTC_Periph) ADD1S() SHIFTR_Mask {
	return SHIFTR_Mask{mmio.UM32{&p.SHIFTR.U32, uint32(ADD1S)}}
}

type TSTR_Bits uint32

func (b TSTR_Bits) Field(mask TSTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSTR_Bits) J(v int) TSTR_Bits {
	return TSTR_Bits(bits.Make32(v, uint32(mask)))
}

type TSTR struct{ mmio.U32 }

func (r *TSTR) Bits(mask TSTR_Bits) TSTR_Bits { return TSTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSTR) StoreBits(mask, b TSTR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSTR) SetBits(mask TSTR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TSTR) ClearBits(mask TSTR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TSTR) Load() TSTR_Bits               { return TSTR_Bits(r.U32.Load()) }
func (r *TSTR) Store(b TSTR_Bits)             { r.U32.Store(uint32(b)) }

func (r *TSTR) AtomicStoreBits(mask, b TSTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TSTR) AtomicSetBits(mask TSTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TSTR) AtomicClearBits(mask TSTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TSTR_Mask struct{ mmio.UM32 }

func (rm TSTR_Mask) Load() TSTR_Bits   { return TSTR_Bits(rm.UM32.Load()) }
func (rm TSTR_Mask) Store(b TSTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TPM() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(TPM)}}
}

func (p *RTC_Periph) THT() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(THT)}}
}

func (p *RTC_Periph) THU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(THU)}}
}

func (p *RTC_Periph) TMNT() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(TMNT)}}
}

func (p *RTC_Periph) TMNU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(TMNU)}}
}

func (p *RTC_Periph) TST() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(TST)}}
}

func (p *RTC_Periph) TSU() TSTR_Mask {
	return TSTR_Mask{mmio.UM32{&p.TSTR.U32, uint32(TSU)}}
}

type TSDR_Bits uint32

func (b TSDR_Bits) Field(mask TSDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSDR_Bits) J(v int) TSDR_Bits {
	return TSDR_Bits(bits.Make32(v, uint32(mask)))
}

type TSDR struct{ mmio.U32 }

func (r *TSDR) Bits(mask TSDR_Bits) TSDR_Bits { return TSDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSDR) StoreBits(mask, b TSDR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSDR) SetBits(mask TSDR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *TSDR) ClearBits(mask TSDR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *TSDR) Load() TSDR_Bits               { return TSDR_Bits(r.U32.Load()) }
func (r *TSDR) Store(b TSDR_Bits)             { r.U32.Store(uint32(b)) }

func (r *TSDR) AtomicStoreBits(mask, b TSDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TSDR) AtomicSetBits(mask TSDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TSDR) AtomicClearBits(mask TSDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TSDR_Mask struct{ mmio.UM32 }

func (rm TSDR_Mask) Load() TSDR_Bits   { return TSDR_Bits(rm.UM32.Load()) }
func (rm TSDR_Mask) Store(b TSDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TWDU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{&p.TSDR.U32, uint32(TWDU)}}
}

func (p *RTC_Periph) TMT() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{&p.TSDR.U32, uint32(TMT)}}
}

func (p *RTC_Periph) TMU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{&p.TSDR.U32, uint32(TMU)}}
}

func (p *RTC_Periph) TDT() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{&p.TSDR.U32, uint32(TDT)}}
}

func (p *RTC_Periph) TDU() TSDR_Mask {
	return TSDR_Mask{mmio.UM32{&p.TSDR.U32, uint32(TDU)}}
}

type TSSSR_Bits uint32

func (b TSSSR_Bits) Field(mask TSSSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSSSR_Bits) J(v int) TSSSR_Bits {
	return TSSSR_Bits(bits.Make32(v, uint32(mask)))
}

type TSSSR struct{ mmio.U32 }

func (r *TSSSR) Bits(mask TSSSR_Bits) TSSSR_Bits { return TSSSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSSSR) StoreBits(mask, b TSSSR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSSSR) SetBits(mask TSSSR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *TSSSR) ClearBits(mask TSSSR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *TSSSR) Load() TSSSR_Bits                { return TSSSR_Bits(r.U32.Load()) }
func (r *TSSSR) Store(b TSSSR_Bits)              { r.U32.Store(uint32(b)) }

func (r *TSSSR) AtomicStoreBits(mask, b TSSSR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TSSSR) AtomicSetBits(mask TSSSR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TSSSR) AtomicClearBits(mask TSSSR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TSSSR_Mask struct{ mmio.UM32 }

func (rm TSSSR_Mask) Load() TSSSR_Bits   { return TSSSR_Bits(rm.UM32.Load()) }
func (rm TSSSR_Mask) Store(b TSSSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TSS() TSSSR_Mask {
	return TSSSR_Mask{mmio.UM32{&p.TSSSR.U32, uint32(TSS)}}
}

type CALR_Bits uint32

func (b CALR_Bits) Field(mask CALR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CALR_Bits) J(v int) CALR_Bits {
	return CALR_Bits(bits.Make32(v, uint32(mask)))
}

type CALR struct{ mmio.U32 }

func (r *CALR) Bits(mask CALR_Bits) CALR_Bits { return CALR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CALR) StoreBits(mask, b CALR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CALR) SetBits(mask CALR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CALR) ClearBits(mask CALR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CALR) Load() CALR_Bits               { return CALR_Bits(r.U32.Load()) }
func (r *CALR) Store(b CALR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CALR) AtomicStoreBits(mask, b CALR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CALR) AtomicSetBits(mask CALR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CALR) AtomicClearBits(mask CALR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CALR_Mask struct{ mmio.UM32 }

func (rm CALR_Mask) Load() CALR_Bits   { return CALR_Bits(rm.UM32.Load()) }
func (rm CALR_Mask) Store(b CALR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) CALP() CALR_Mask {
	return CALR_Mask{mmio.UM32{&p.CALR.U32, uint32(CALP)}}
}

func (p *RTC_Periph) CALW8() CALR_Mask {
	return CALR_Mask{mmio.UM32{&p.CALR.U32, uint32(CALW8)}}
}

func (p *RTC_Periph) CALW16() CALR_Mask {
	return CALR_Mask{mmio.UM32{&p.CALR.U32, uint32(CALW16)}}
}

func (p *RTC_Periph) CALM() CALR_Mask {
	return CALR_Mask{mmio.UM32{&p.CALR.U32, uint32(CALM)}}
}

type TAMPCR_Bits uint32

func (b TAMPCR_Bits) Field(mask TAMPCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TAMPCR_Bits) J(v int) TAMPCR_Bits {
	return TAMPCR_Bits(bits.Make32(v, uint32(mask)))
}

type TAMPCR struct{ mmio.U32 }

func (r *TAMPCR) Bits(mask TAMPCR_Bits) TAMPCR_Bits { return TAMPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TAMPCR) StoreBits(mask, b TAMPCR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TAMPCR) SetBits(mask TAMPCR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *TAMPCR) ClearBits(mask TAMPCR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *TAMPCR) Load() TAMPCR_Bits                 { return TAMPCR_Bits(r.U32.Load()) }
func (r *TAMPCR) Store(b TAMPCR_Bits)               { r.U32.Store(uint32(b)) }

func (r *TAMPCR) AtomicStoreBits(mask, b TAMPCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *TAMPCR) AtomicSetBits(mask TAMPCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *TAMPCR) AtomicClearBits(mask TAMPCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type TAMPCR_Mask struct{ mmio.UM32 }

func (rm TAMPCR_Mask) Load() TAMPCR_Bits   { return TAMPCR_Bits(rm.UM32.Load()) }
func (rm TAMPCR_Mask) Store(b TAMPCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) TAMP3MF() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3MF)}}
}

func (p *RTC_Periph) TAMP3NOERASE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3NOERASE)}}
}

func (p *RTC_Periph) TAMP3IE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3IE)}}
}

func (p *RTC_Periph) TAMP2MF() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2MF)}}
}

func (p *RTC_Periph) TAMP2NOERASE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2NOERASE)}}
}

func (p *RTC_Periph) TAMP2IE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2IE)}}
}

func (p *RTC_Periph) TAMP1MF() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1MF)}}
}

func (p *RTC_Periph) TAMP1NOERASE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1NOERASE)}}
}

func (p *RTC_Periph) TAMP1IE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1IE)}}
}

func (p *RTC_Periph) TAMPPUDIS() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPPUDIS)}}
}

func (p *RTC_Periph) TAMPPRCH() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPPRCH)}}
}

func (p *RTC_Periph) TAMPFLT() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPFLT)}}
}

func (p *RTC_Periph) TAMPFREQ() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPFREQ)}}
}

func (p *RTC_Periph) TAMPTS() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPTS)}}
}

func (p *RTC_Periph) TAMP3TRG() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3TRG)}}
}

func (p *RTC_Periph) TAMP3E() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP3E)}}
}

func (p *RTC_Periph) TAMP2TRG() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2TRG)}}
}

func (p *RTC_Periph) TAMP2E() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP2E)}}
}

func (p *RTC_Periph) TAMPIE() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMPIE)}}
}

func (p *RTC_Periph) TAMP1TRG() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1TRG)}}
}

func (p *RTC_Periph) TAMP1E() TAMPCR_Mask {
	return TAMPCR_Mask{mmio.UM32{&p.TAMPCR.U32, uint32(TAMP1E)}}
}

type ALRMSSR_Bits uint32

func (b ALRMSSR_Bits) Field(mask ALRMSSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ALRMSSR_Bits) J(v int) ALRMSSR_Bits {
	return ALRMSSR_Bits(bits.Make32(v, uint32(mask)))
}

type ALRMSSR struct{ mmio.U32 }

func (r *ALRMSSR) Bits(mask ALRMSSR_Bits) ALRMSSR_Bits { return ALRMSSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ALRMSSR) StoreBits(mask, b ALRMSSR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ALRMSSR) SetBits(mask ALRMSSR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *ALRMSSR) ClearBits(mask ALRMSSR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *ALRMSSR) Load() ALRMSSR_Bits                  { return ALRMSSR_Bits(r.U32.Load()) }
func (r *ALRMSSR) Store(b ALRMSSR_Bits)                { r.U32.Store(uint32(b)) }

func (r *ALRMSSR) AtomicStoreBits(mask, b ALRMSSR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *ALRMSSR) AtomicSetBits(mask ALRMSSR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ALRMSSR) AtomicClearBits(mask ALRMSSR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type ALRMSSR_Mask struct{ mmio.UM32 }

func (rm ALRMSSR_Mask) Load() ALRMSSR_Bits   { return ALRMSSR_Bits(rm.UM32.Load()) }
func (rm ALRMSSR_Mask) Store(b ALRMSSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) AMASKSS(n int) ALRMSSR_Mask {
	return ALRMSSR_Mask{mmio.UM32{&p.ALRMSSR[n].U32, uint32(AMASKSS)}}
}

func (p *RTC_Periph) ASS(n int) ALRMSSR_Mask {
	return ALRMSSR_Mask{mmio.UM32{&p.ALRMSSR[n].U32, uint32(ASS)}}
}

type OR_Bits uint32

func (b OR_Bits) Field(mask OR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OR_Bits) J(v int) OR_Bits {
	return OR_Bits(bits.Make32(v, uint32(mask)))
}

type OR struct{ mmio.U32 }

func (r *OR) Bits(mask OR_Bits) OR_Bits { return OR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OR) StoreBits(mask, b OR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OR) SetBits(mask OR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *OR) ClearBits(mask OR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *OR) Load() OR_Bits             { return OR_Bits(r.U32.Load()) }
func (r *OR) Store(b OR_Bits)           { r.U32.Store(uint32(b)) }

func (r *OR) AtomicStoreBits(mask, b OR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OR) AtomicSetBits(mask OR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OR) AtomicClearBits(mask OR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OR_Mask struct{ mmio.UM32 }

func (rm OR_Mask) Load() OR_Bits   { return OR_Bits(rm.UM32.Load()) }
func (rm OR_Mask) Store(b OR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RTC_Periph) OUT_RMP() OR_Mask {
	return OR_Mask{mmio.UM32{&p.OR.U32, uint32(OUT_RMP)}}
}

func (p *RTC_Periph) ALARMOUTTYPE() OR_Mask {
	return OR_Mask{mmio.UM32{&p.OR.U32, uint32(ALARMOUTTYPE)}}
}

type BKPR_Bits uint32

func (b BKPR_Bits) Field(mask BKPR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BKPR_Bits) J(v int) BKPR_Bits {
	return BKPR_Bits(bits.Make32(v, uint32(mask)))
}

type BKPR struct{ mmio.U32 }

func (r *BKPR) Bits(mask BKPR_Bits) BKPR_Bits { return BKPR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BKPR) StoreBits(mask, b BKPR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BKPR) SetBits(mask BKPR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BKPR) ClearBits(mask BKPR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BKPR) Load() BKPR_Bits               { return BKPR_Bits(r.U32.Load()) }
func (r *BKPR) Store(b BKPR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BKPR) AtomicStoreBits(mask, b BKPR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BKPR) AtomicSetBits(mask BKPR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BKPR) AtomicClearBits(mask BKPR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BKPR_Mask struct{ mmio.UM32 }

func (rm BKPR_Mask) Load() BKPR_Bits   { return BKPR_Bits(rm.UM32.Load()) }
func (rm BKPR_Mask) Store(b BKPR_Bits) { rm.UM32.Store(uint32(b)) }
