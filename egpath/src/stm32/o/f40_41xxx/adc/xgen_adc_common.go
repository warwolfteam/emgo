package adc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type ADC_Common_Periph struct {
	CSR CSR
	CCR CCR
	CDR CDR
}

func (p *ADC_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var ADC = (*ADC_Common_Periph)(unsafe.Pointer(uintptr(mmap.ADC_BASE)))

type CSR_Bits uint32

func (b CSR_Bits) Field(mask CSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR_Bits) J(v int) CSR_Bits {
	return CSR_Bits(bits.Make32(v, uint32(mask)))
}

type CSR struct{ mmio.U32 }

func (r *CSR) Bits(mask CSR_Bits) CSR_Bits { return CSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CSR) StoreBits(mask, b CSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CSR) SetBits(mask CSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CSR) ClearBits(mask CSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CSR) Load() CSR_Bits              { return CSR_Bits(r.U32.Load()) }
func (r *CSR) Store(b CSR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CSR) AtomicStoreBits(mask, b CSR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CSR) AtomicSetBits(mask CSR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CSR) AtomicClearBits(mask CSR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CSR_Mask struct{ mmio.UM32 }

func (rm CSR_Mask) Load() CSR_Bits   { return CSR_Bits(rm.UM32.Load()) }
func (rm CSR_Mask) Store(b CSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) AWD1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD1)}}
}

func (p *ADC_Common_Periph) EOC1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOC1)}}
}

func (p *ADC_Common_Periph) JEOC1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOC1)}}
}

func (p *ADC_Common_Periph) JSTRT1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JSTRT1)}}
}

func (p *ADC_Common_Periph) STRT1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(STRT1)}}
}

func (p *ADC_Common_Periph) OVR1() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OVR1)}}
}

func (p *ADC_Common_Periph) AWD2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD2)}}
}

func (p *ADC_Common_Periph) EOC2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOC2)}}
}

func (p *ADC_Common_Periph) JEOC2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOC2)}}
}

func (p *ADC_Common_Periph) JSTRT2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JSTRT2)}}
}

func (p *ADC_Common_Periph) STRT2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(STRT2)}}
}

func (p *ADC_Common_Periph) OVR2() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OVR2)}}
}

func (p *ADC_Common_Periph) AWD3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(AWD3)}}
}

func (p *ADC_Common_Periph) EOC3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(EOC3)}}
}

func (p *ADC_Common_Periph) JEOC3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JEOC3)}}
}

func (p *ADC_Common_Periph) JSTRT3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(JSTRT3)}}
}

func (p *ADC_Common_Periph) STRT3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(STRT3)}}
}

func (p *ADC_Common_Periph) OVR3() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OVR3)}}
}

type CCR_Bits uint32

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U32 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U32.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CCR) AtomicStoreBits(mask, b CCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CCR) AtomicSetBits(mask CCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CCR) AtomicClearBits(mask CCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CCR_Mask struct{ mmio.UM32 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM32.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) MULTI() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(MULTI)}}
}

func (p *ADC_Common_Periph) DELAY() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DELAY)}}
}

func (p *ADC_Common_Periph) DDS() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DDS)}}
}

func (p *ADC_Common_Periph) DMA() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(DMA)}}
}

func (p *ADC_Common_Periph) ADCPRE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(ADCPRE)}}
}

func (p *ADC_Common_Periph) VBATE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(VBATE)}}
}

func (p *ADC_Common_Periph) TSVREFE() CCR_Mask {
	return CCR_Mask{mmio.UM32{&p.CCR.U32, uint32(TSVREFE)}}
}

type CDR_Bits uint32

func (b CDR_Bits) Field(mask CDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CDR_Bits) J(v int) CDR_Bits {
	return CDR_Bits(bits.Make32(v, uint32(mask)))
}

type CDR struct{ mmio.U32 }

func (r *CDR) Bits(mask CDR_Bits) CDR_Bits { return CDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CDR) StoreBits(mask, b CDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CDR) SetBits(mask CDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CDR) ClearBits(mask CDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CDR) Load() CDR_Bits              { return CDR_Bits(r.U32.Load()) }
func (r *CDR) Store(b CDR_Bits)            { r.U32.Store(uint32(b)) }

func (r *CDR) AtomicStoreBits(mask, b CDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CDR) AtomicSetBits(mask CDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CDR) AtomicClearBits(mask CDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CDR_Mask struct{ mmio.UM32 }

func (rm CDR_Mask) Load() CDR_Bits   { return CDR_Bits(rm.UM32.Load()) }
func (rm CDR_Mask) Store(b CDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *ADC_Common_Periph) DATA1() CDR_Mask {
	return CDR_Mask{mmio.UM32{&p.CDR.U32, uint32(DATA1)}}
}

func (p *ADC_Common_Periph) DATA2() CDR_Mask {
	return CDR_Mask{mmio.UM32{&p.CDR.U32, uint32(DATA2)}}
}
