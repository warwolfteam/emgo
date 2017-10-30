// +build l476xx

package opamp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type OPAMP_Periph struct {
	CSR   CSR
	OTR   OTR
	LPOTR LPOTR
}

func (p *OPAMP_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var OPAMP = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP_BASE)))

//emgo:const
var OPAMP1 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP1_BASE)))

//emgo:const
var OPAMP2 = (*OPAMP_Periph)(unsafe.Pointer(uintptr(mmap.OPAMP2_BASE)))

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

func (p *OPAMP_Periph) OPAMPxEN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPAMPxEN)}}
}

func (p *OPAMP_Periph) OPALPM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPALPM)}}
}

func (p *OPAMP_Periph) OPAMODE() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(OPAMODE)}}
}

func (p *OPAMP_Periph) PGGAIN() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(PGGAIN)}}
}

func (p *OPAMP_Periph) VMSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VMSEL)}}
}

func (p *OPAMP_Periph) VPSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(VPSEL)}}
}

func (p *OPAMP_Periph) CALON() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CALON)}}
}

func (p *OPAMP_Periph) CALSEL() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CALSEL)}}
}

func (p *OPAMP_Periph) USERTRIM() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(USERTRIM)}}
}

func (p *OPAMP_Periph) CALOUT() CSR_Mask {
	return CSR_Mask{mmio.UM32{&p.CSR.U32, uint32(CALOUT)}}
}

type OTR_Bits uint32

func (b OTR_Bits) Field(mask OTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTR_Bits) J(v int) OTR_Bits {
	return OTR_Bits(bits.Make32(v, uint32(mask)))
}

type OTR struct{ mmio.U32 }

func (r *OTR) Bits(mask OTR_Bits) OTR_Bits { return OTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OTR) StoreBits(mask, b OTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OTR) SetBits(mask OTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *OTR) ClearBits(mask OTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *OTR) Load() OTR_Bits              { return OTR_Bits(r.U32.Load()) }
func (r *OTR) Store(b OTR_Bits)            { r.U32.Store(uint32(b)) }

func (r *OTR) AtomicStoreBits(mask, b OTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OTR) AtomicSetBits(mask OTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OTR) AtomicClearBits(mask OTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OTR_Mask struct{ mmio.UM32 }

func (rm OTR_Mask) Load() OTR_Bits   { return OTR_Bits(rm.UM32.Load()) }
func (rm OTR_Mask) Store(b OTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) TRIMOFFSETN() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(TRIMOFFSETN)}}
}

func (p *OPAMP_Periph) TRIMOFFSETP() OTR_Mask {
	return OTR_Mask{mmio.UM32{&p.OTR.U32, uint32(TRIMOFFSETP)}}
}

type LPOTR_Bits uint32

func (b LPOTR_Bits) Field(mask LPOTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LPOTR_Bits) J(v int) LPOTR_Bits {
	return LPOTR_Bits(bits.Make32(v, uint32(mask)))
}

type LPOTR struct{ mmio.U32 }

func (r *LPOTR) Bits(mask LPOTR_Bits) LPOTR_Bits { return LPOTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LPOTR) StoreBits(mask, b LPOTR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LPOTR) SetBits(mask LPOTR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *LPOTR) ClearBits(mask LPOTR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *LPOTR) Load() LPOTR_Bits                { return LPOTR_Bits(r.U32.Load()) }
func (r *LPOTR) Store(b LPOTR_Bits)              { r.U32.Store(uint32(b)) }

func (r *LPOTR) AtomicStoreBits(mask, b LPOTR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *LPOTR) AtomicSetBits(mask LPOTR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *LPOTR) AtomicClearBits(mask LPOTR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type LPOTR_Mask struct{ mmio.UM32 }

func (rm LPOTR_Mask) Load() LPOTR_Bits   { return LPOTR_Bits(rm.UM32.Load()) }
func (rm LPOTR_Mask) Store(b LPOTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *OPAMP_Periph) TRIMLPOFFSETN() LPOTR_Mask {
	return LPOTR_Mask{mmio.UM32{&p.LPOTR.U32, uint32(TRIMLPOFFSETN)}}
}

func (p *OPAMP_Periph) TRIMLPOFFSETP() LPOTR_Mask {
	return LPOTR_Mask{mmio.UM32{&p.LPOTR.U32, uint32(TRIMLPOFFSETP)}}
}
