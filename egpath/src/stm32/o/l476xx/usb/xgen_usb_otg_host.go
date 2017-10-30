package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_Host_Periph struct {
	HCFG     HCFG
	HFIR     HFIR
	HFNUM    HFNUM
	_        uint32
	HPTXSTS  HPTXSTS
	HAINT    HAINT
	HAINTMSK HAINTMSK
}

func (p *USB_OTG_Host_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type HCFG_Bits uint32

func (b HCFG_Bits) Field(mask HCFG_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCFG_Bits) J(v int) HCFG_Bits {
	return HCFG_Bits(bits.Make32(v, uint32(mask)))
}

type HCFG struct{ mmio.U32 }

func (r *HCFG) Bits(mask HCFG_Bits) HCFG_Bits { return HCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *HCFG) StoreBits(mask, b HCFG_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HCFG) SetBits(mask HCFG_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *HCFG) ClearBits(mask HCFG_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *HCFG) Load() HCFG_Bits               { return HCFG_Bits(r.U32.Load()) }
func (r *HCFG) Store(b HCFG_Bits)             { r.U32.Store(uint32(b)) }

func (r *HCFG) AtomicStoreBits(mask, b HCFG_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *HCFG) AtomicSetBits(mask HCFG_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HCFG) AtomicClearBits(mask HCFG_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type HCFG_Mask struct{ mmio.UM32 }

func (rm HCFG_Mask) Load() HCFG_Bits   { return HCFG_Bits(rm.UM32.Load()) }
func (rm HCFG_Mask) Store(b HCFG_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FSLSPCS() HCFG_Mask {
	return HCFG_Mask{mmio.UM32{&p.HCFG.U32, uint32(FSLSPCS)}}
}

func (p *USB_OTG_Host_Periph) FSLSS() HCFG_Mask {
	return HCFG_Mask{mmio.UM32{&p.HCFG.U32, uint32(FSLSS)}}
}

type HFIR_Bits uint32

func (b HFIR_Bits) Field(mask HFIR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HFIR_Bits) J(v int) HFIR_Bits {
	return HFIR_Bits(bits.Make32(v, uint32(mask)))
}

type HFIR struct{ mmio.U32 }

func (r *HFIR) Bits(mask HFIR_Bits) HFIR_Bits { return HFIR_Bits(r.U32.Bits(uint32(mask))) }
func (r *HFIR) StoreBits(mask, b HFIR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HFIR) SetBits(mask HFIR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *HFIR) ClearBits(mask HFIR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *HFIR) Load() HFIR_Bits               { return HFIR_Bits(r.U32.Load()) }
func (r *HFIR) Store(b HFIR_Bits)             { r.U32.Store(uint32(b)) }

func (r *HFIR) AtomicStoreBits(mask, b HFIR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *HFIR) AtomicSetBits(mask HFIR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HFIR) AtomicClearBits(mask HFIR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type HFIR_Mask struct{ mmio.UM32 }

func (rm HFIR_Mask) Load() HFIR_Bits   { return HFIR_Bits(rm.UM32.Load()) }
func (rm HFIR_Mask) Store(b HFIR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FRIVL() HFIR_Mask {
	return HFIR_Mask{mmio.UM32{&p.HFIR.U32, uint32(FRIVL)}}
}

type HFNUM_Bits uint32

func (b HFNUM_Bits) Field(mask HFNUM_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HFNUM_Bits) J(v int) HFNUM_Bits {
	return HFNUM_Bits(bits.Make32(v, uint32(mask)))
}

type HFNUM struct{ mmio.U32 }

func (r *HFNUM) Bits(mask HFNUM_Bits) HFNUM_Bits { return HFNUM_Bits(r.U32.Bits(uint32(mask))) }
func (r *HFNUM) StoreBits(mask, b HFNUM_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HFNUM) SetBits(mask HFNUM_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *HFNUM) ClearBits(mask HFNUM_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *HFNUM) Load() HFNUM_Bits                { return HFNUM_Bits(r.U32.Load()) }
func (r *HFNUM) Store(b HFNUM_Bits)              { r.U32.Store(uint32(b)) }

func (r *HFNUM) AtomicStoreBits(mask, b HFNUM_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *HFNUM) AtomicSetBits(mask HFNUM_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HFNUM) AtomicClearBits(mask HFNUM_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type HFNUM_Mask struct{ mmio.UM32 }

func (rm HFNUM_Mask) Load() HFNUM_Bits   { return HFNUM_Bits(rm.UM32.Load()) }
func (rm HFNUM_Mask) Store(b HFNUM_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) FRNUM() HFNUM_Mask {
	return HFNUM_Mask{mmio.UM32{&p.HFNUM.U32, uint32(FRNUM)}}
}

func (p *USB_OTG_Host_Periph) FTREM() HFNUM_Mask {
	return HFNUM_Mask{mmio.UM32{&p.HFNUM.U32, uint32(FTREM)}}
}

type HPTXSTS_Bits uint32

func (b HPTXSTS_Bits) Field(mask HPTXSTS_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HPTXSTS_Bits) J(v int) HPTXSTS_Bits {
	return HPTXSTS_Bits(bits.Make32(v, uint32(mask)))
}

type HPTXSTS struct{ mmio.U32 }

func (r *HPTXSTS) Bits(mask HPTXSTS_Bits) HPTXSTS_Bits { return HPTXSTS_Bits(r.U32.Bits(uint32(mask))) }
func (r *HPTXSTS) StoreBits(mask, b HPTXSTS_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HPTXSTS) SetBits(mask HPTXSTS_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *HPTXSTS) ClearBits(mask HPTXSTS_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *HPTXSTS) Load() HPTXSTS_Bits                  { return HPTXSTS_Bits(r.U32.Load()) }
func (r *HPTXSTS) Store(b HPTXSTS_Bits)                { r.U32.Store(uint32(b)) }

func (r *HPTXSTS) AtomicStoreBits(mask, b HPTXSTS_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *HPTXSTS) AtomicSetBits(mask HPTXSTS_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HPTXSTS) AtomicClearBits(mask HPTXSTS_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type HPTXSTS_Mask struct{ mmio.UM32 }

func (rm HPTXSTS_Mask) Load() HPTXSTS_Bits   { return HPTXSTS_Bits(rm.UM32.Load()) }
func (rm HPTXSTS_Mask) Store(b HPTXSTS_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) PTXFSAVL() HPTXSTS_Mask {
	return HPTXSTS_Mask{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXFSAVL)}}
}

func (p *USB_OTG_Host_Periph) PTXQSAV() HPTXSTS_Mask {
	return HPTXSTS_Mask{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXQSAV)}}
}

func (p *USB_OTG_Host_Periph) PTXQTOP() HPTXSTS_Mask {
	return HPTXSTS_Mask{mmio.UM32{&p.HPTXSTS.U32, uint32(PTXQTOP)}}
}

type HAINT_Bits uint32

func (b HAINT_Bits) Field(mask HAINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HAINT_Bits) J(v int) HAINT_Bits {
	return HAINT_Bits(bits.Make32(v, uint32(mask)))
}

type HAINT struct{ mmio.U32 }

func (r *HAINT) Bits(mask HAINT_Bits) HAINT_Bits { return HAINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *HAINT) StoreBits(mask, b HAINT_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HAINT) SetBits(mask HAINT_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *HAINT) ClearBits(mask HAINT_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *HAINT) Load() HAINT_Bits                { return HAINT_Bits(r.U32.Load()) }
func (r *HAINT) Store(b HAINT_Bits)              { r.U32.Store(uint32(b)) }

func (r *HAINT) AtomicStoreBits(mask, b HAINT_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *HAINT) AtomicSetBits(mask HAINT_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HAINT) AtomicClearBits(mask HAINT_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type HAINT_Mask struct{ mmio.UM32 }

func (rm HAINT_Mask) Load() HAINT_Bits   { return HAINT_Bits(rm.UM32.Load()) }
func (rm HAINT_Mask) Store(b HAINT_Bits) { rm.UM32.Store(uint32(b)) }

type HAINTMSK_Bits uint32

func (b HAINTMSK_Bits) Field(mask HAINTMSK_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HAINTMSK_Bits) J(v int) HAINTMSK_Bits {
	return HAINTMSK_Bits(bits.Make32(v, uint32(mask)))
}

type HAINTMSK struct{ mmio.U32 }

func (r *HAINTMSK) Bits(mask HAINTMSK_Bits) HAINTMSK_Bits {
	return HAINTMSK_Bits(r.U32.Bits(uint32(mask)))
}
func (r *HAINTMSK) StoreBits(mask, b HAINTMSK_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HAINTMSK) SetBits(mask HAINTMSK_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HAINTMSK) ClearBits(mask HAINTMSK_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HAINTMSK) Load() HAINTMSK_Bits             { return HAINTMSK_Bits(r.U32.Load()) }
func (r *HAINTMSK) Store(b HAINTMSK_Bits)           { r.U32.Store(uint32(b)) }

func (r *HAINTMSK) AtomicStoreBits(mask, b HAINTMSK_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *HAINTMSK) AtomicSetBits(mask HAINTMSK_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *HAINTMSK) AtomicClearBits(mask HAINTMSK_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type HAINTMSK_Mask struct{ mmio.UM32 }

func (rm HAINTMSK_Mask) Load() HAINTMSK_Bits   { return HAINTMSK_Bits(rm.UM32.Load()) }
func (rm HAINTMSK_Mask) Store(b HAINTMSK_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_Host_Periph) HAINTM() HAINTMSK_Mask {
	return HAINTMSK_Mask{mmio.UM32{&p.HAINTMSK.U32, uint32(HAINTM)}}
}
