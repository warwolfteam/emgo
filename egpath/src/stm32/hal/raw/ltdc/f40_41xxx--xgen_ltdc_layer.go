// +build f40_41xxx

package ltdc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type LTDC_Layer_Periph struct {
	CR     CR
	WHPCR  WHPCR
	WVPCR  WVPCR
	CKCR   CKCR
	PFCR   PFCR
	CACR   CACR
	DCCR   DCCR
	BFCR   BFCR
	_      [2]uint32
	CFBAR  CFBAR
	CFBLR  CFBLR
	CFBLNR CFBLNR
	_      [3]uint32
	CLUTWR CLUTWR
}

func (p *LTDC_Layer_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var LTDC_Layer1 = (*LTDC_Layer_Periph)(unsafe.Pointer(uintptr(mmap.LTDC_Layer1_BASE)))

//emgo:const
var LTDC_Layer2 = (*LTDC_Layer_Periph)(unsafe.Pointer(uintptr(mmap.LTDC_Layer2_BASE)))

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

type WHPCR_Bits uint32

func (b WHPCR_Bits) Field(mask WHPCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WHPCR_Bits) J(v int) WHPCR_Bits {
	return WHPCR_Bits(bits.Make32(v, uint32(mask)))
}

type WHPCR struct{ mmio.U32 }

func (r *WHPCR) Bits(mask WHPCR_Bits) WHPCR_Bits { return WHPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WHPCR) StoreBits(mask, b WHPCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WHPCR) SetBits(mask WHPCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WHPCR) ClearBits(mask WHPCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WHPCR) Load() WHPCR_Bits                { return WHPCR_Bits(r.U32.Load()) }
func (r *WHPCR) Store(b WHPCR_Bits)              { r.U32.Store(uint32(b)) }

func (r *WHPCR) AtomicStoreBits(mask, b WHPCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WHPCR) AtomicSetBits(mask WHPCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WHPCR) AtomicClearBits(mask WHPCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WHPCR_Mask struct{ mmio.UM32 }

func (rm WHPCR_Mask) Load() WHPCR_Bits   { return WHPCR_Bits(rm.UM32.Load()) }
func (rm WHPCR_Mask) Store(b WHPCR_Bits) { rm.UM32.Store(uint32(b)) }

type WVPCR_Bits uint32

func (b WVPCR_Bits) Field(mask WVPCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WVPCR_Bits) J(v int) WVPCR_Bits {
	return WVPCR_Bits(bits.Make32(v, uint32(mask)))
}

type WVPCR struct{ mmio.U32 }

func (r *WVPCR) Bits(mask WVPCR_Bits) WVPCR_Bits { return WVPCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *WVPCR) StoreBits(mask, b WVPCR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *WVPCR) SetBits(mask WVPCR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *WVPCR) ClearBits(mask WVPCR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *WVPCR) Load() WVPCR_Bits                { return WVPCR_Bits(r.U32.Load()) }
func (r *WVPCR) Store(b WVPCR_Bits)              { r.U32.Store(uint32(b)) }

func (r *WVPCR) AtomicStoreBits(mask, b WVPCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *WVPCR) AtomicSetBits(mask WVPCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *WVPCR) AtomicClearBits(mask WVPCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type WVPCR_Mask struct{ mmio.UM32 }

func (rm WVPCR_Mask) Load() WVPCR_Bits   { return WVPCR_Bits(rm.UM32.Load()) }
func (rm WVPCR_Mask) Store(b WVPCR_Bits) { rm.UM32.Store(uint32(b)) }

type CKCR_Bits uint32

func (b CKCR_Bits) Field(mask CKCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CKCR_Bits) J(v int) CKCR_Bits {
	return CKCR_Bits(bits.Make32(v, uint32(mask)))
}

type CKCR struct{ mmio.U32 }

func (r *CKCR) Bits(mask CKCR_Bits) CKCR_Bits { return CKCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CKCR) StoreBits(mask, b CKCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CKCR) SetBits(mask CKCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CKCR) ClearBits(mask CKCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CKCR) Load() CKCR_Bits               { return CKCR_Bits(r.U32.Load()) }
func (r *CKCR) Store(b CKCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CKCR) AtomicStoreBits(mask, b CKCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CKCR) AtomicSetBits(mask CKCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CKCR) AtomicClearBits(mask CKCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CKCR_Mask struct{ mmio.UM32 }

func (rm CKCR_Mask) Load() CKCR_Bits   { return CKCR_Bits(rm.UM32.Load()) }
func (rm CKCR_Mask) Store(b CKCR_Bits) { rm.UM32.Store(uint32(b)) }

type PFCR_Bits uint32

func (b PFCR_Bits) Field(mask PFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PFCR_Bits) J(v int) PFCR_Bits {
	return PFCR_Bits(bits.Make32(v, uint32(mask)))
}

type PFCR struct{ mmio.U32 }

func (r *PFCR) Bits(mask PFCR_Bits) PFCR_Bits { return PFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PFCR) StoreBits(mask, b PFCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PFCR) SetBits(mask PFCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *PFCR) ClearBits(mask PFCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *PFCR) Load() PFCR_Bits               { return PFCR_Bits(r.U32.Load()) }
func (r *PFCR) Store(b PFCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *PFCR) AtomicStoreBits(mask, b PFCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PFCR) AtomicSetBits(mask PFCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PFCR) AtomicClearBits(mask PFCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PFCR_Mask struct{ mmio.UM32 }

func (rm PFCR_Mask) Load() PFCR_Bits   { return PFCR_Bits(rm.UM32.Load()) }
func (rm PFCR_Mask) Store(b PFCR_Bits) { rm.UM32.Store(uint32(b)) }

type CACR_Bits uint32

func (b CACR_Bits) Field(mask CACR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CACR_Bits) J(v int) CACR_Bits {
	return CACR_Bits(bits.Make32(v, uint32(mask)))
}

type CACR struct{ mmio.U32 }

func (r *CACR) Bits(mask CACR_Bits) CACR_Bits { return CACR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CACR) StoreBits(mask, b CACR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CACR) SetBits(mask CACR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CACR) ClearBits(mask CACR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CACR) Load() CACR_Bits               { return CACR_Bits(r.U32.Load()) }
func (r *CACR) Store(b CACR_Bits)             { r.U32.Store(uint32(b)) }

func (r *CACR) AtomicStoreBits(mask, b CACR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CACR) AtomicSetBits(mask CACR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CACR) AtomicClearBits(mask CACR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CACR_Mask struct{ mmio.UM32 }

func (rm CACR_Mask) Load() CACR_Bits   { return CACR_Bits(rm.UM32.Load()) }
func (rm CACR_Mask) Store(b CACR_Bits) { rm.UM32.Store(uint32(b)) }

type DCCR_Bits uint32

func (b DCCR_Bits) Field(mask DCCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DCCR_Bits) J(v int) DCCR_Bits {
	return DCCR_Bits(bits.Make32(v, uint32(mask)))
}

type DCCR struct{ mmio.U32 }

func (r *DCCR) Bits(mask DCCR_Bits) DCCR_Bits { return DCCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DCCR) StoreBits(mask, b DCCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DCCR) SetBits(mask DCCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *DCCR) ClearBits(mask DCCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *DCCR) Load() DCCR_Bits               { return DCCR_Bits(r.U32.Load()) }
func (r *DCCR) Store(b DCCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *DCCR) AtomicStoreBits(mask, b DCCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *DCCR) AtomicSetBits(mask DCCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DCCR) AtomicClearBits(mask DCCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type DCCR_Mask struct{ mmio.UM32 }

func (rm DCCR_Mask) Load() DCCR_Bits   { return DCCR_Bits(rm.UM32.Load()) }
func (rm DCCR_Mask) Store(b DCCR_Bits) { rm.UM32.Store(uint32(b)) }

type BFCR_Bits uint32

func (b BFCR_Bits) Field(mask BFCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BFCR_Bits) J(v int) BFCR_Bits {
	return BFCR_Bits(bits.Make32(v, uint32(mask)))
}

type BFCR struct{ mmio.U32 }

func (r *BFCR) Bits(mask BFCR_Bits) BFCR_Bits { return BFCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BFCR) StoreBits(mask, b BFCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BFCR) SetBits(mask BFCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BFCR) ClearBits(mask BFCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BFCR) Load() BFCR_Bits               { return BFCR_Bits(r.U32.Load()) }
func (r *BFCR) Store(b BFCR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BFCR) AtomicStoreBits(mask, b BFCR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BFCR) AtomicSetBits(mask BFCR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BFCR) AtomicClearBits(mask BFCR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BFCR_Mask struct{ mmio.UM32 }

func (rm BFCR_Mask) Load() BFCR_Bits   { return BFCR_Bits(rm.UM32.Load()) }
func (rm BFCR_Mask) Store(b BFCR_Bits) { rm.UM32.Store(uint32(b)) }

type CFBAR_Bits uint32

func (b CFBAR_Bits) Field(mask CFBAR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFBAR_Bits) J(v int) CFBAR_Bits {
	return CFBAR_Bits(bits.Make32(v, uint32(mask)))
}

type CFBAR struct{ mmio.U32 }

func (r *CFBAR) Bits(mask CFBAR_Bits) CFBAR_Bits { return CFBAR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFBAR) StoreBits(mask, b CFBAR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFBAR) SetBits(mask CFBAR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CFBAR) ClearBits(mask CFBAR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CFBAR) Load() CFBAR_Bits                { return CFBAR_Bits(r.U32.Load()) }
func (r *CFBAR) Store(b CFBAR_Bits)              { r.U32.Store(uint32(b)) }

func (r *CFBAR) AtomicStoreBits(mask, b CFBAR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CFBAR) AtomicSetBits(mask CFBAR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFBAR) AtomicClearBits(mask CFBAR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CFBAR_Mask struct{ mmio.UM32 }

func (rm CFBAR_Mask) Load() CFBAR_Bits   { return CFBAR_Bits(rm.UM32.Load()) }
func (rm CFBAR_Mask) Store(b CFBAR_Bits) { rm.UM32.Store(uint32(b)) }

type CFBLR_Bits uint32

func (b CFBLR_Bits) Field(mask CFBLR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFBLR_Bits) J(v int) CFBLR_Bits {
	return CFBLR_Bits(bits.Make32(v, uint32(mask)))
}

type CFBLR struct{ mmio.U32 }

func (r *CFBLR) Bits(mask CFBLR_Bits) CFBLR_Bits { return CFBLR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFBLR) StoreBits(mask, b CFBLR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFBLR) SetBits(mask CFBLR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CFBLR) ClearBits(mask CFBLR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CFBLR) Load() CFBLR_Bits                { return CFBLR_Bits(r.U32.Load()) }
func (r *CFBLR) Store(b CFBLR_Bits)              { r.U32.Store(uint32(b)) }

func (r *CFBLR) AtomicStoreBits(mask, b CFBLR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CFBLR) AtomicSetBits(mask CFBLR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFBLR) AtomicClearBits(mask CFBLR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CFBLR_Mask struct{ mmio.UM32 }

func (rm CFBLR_Mask) Load() CFBLR_Bits   { return CFBLR_Bits(rm.UM32.Load()) }
func (rm CFBLR_Mask) Store(b CFBLR_Bits) { rm.UM32.Store(uint32(b)) }

type CFBLNR_Bits uint32

func (b CFBLNR_Bits) Field(mask CFBLNR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFBLNR_Bits) J(v int) CFBLNR_Bits {
	return CFBLNR_Bits(bits.Make32(v, uint32(mask)))
}

type CFBLNR struct{ mmio.U32 }

func (r *CFBLNR) Bits(mask CFBLNR_Bits) CFBLNR_Bits { return CFBLNR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CFBLNR) StoreBits(mask, b CFBLNR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CFBLNR) SetBits(mask CFBLNR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CFBLNR) ClearBits(mask CFBLNR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CFBLNR) Load() CFBLNR_Bits                 { return CFBLNR_Bits(r.U32.Load()) }
func (r *CFBLNR) Store(b CFBLNR_Bits)               { r.U32.Store(uint32(b)) }

func (r *CFBLNR) AtomicStoreBits(mask, b CFBLNR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CFBLNR) AtomicSetBits(mask CFBLNR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CFBLNR) AtomicClearBits(mask CFBLNR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CFBLNR_Mask struct{ mmio.UM32 }

func (rm CFBLNR_Mask) Load() CFBLNR_Bits   { return CFBLNR_Bits(rm.UM32.Load()) }
func (rm CFBLNR_Mask) Store(b CFBLNR_Bits) { rm.UM32.Store(uint32(b)) }

type CLUTWR_Bits uint32

func (b CLUTWR_Bits) Field(mask CLUTWR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLUTWR_Bits) J(v int) CLUTWR_Bits {
	return CLUTWR_Bits(bits.Make32(v, uint32(mask)))
}

type CLUTWR struct{ mmio.U32 }

func (r *CLUTWR) Bits(mask CLUTWR_Bits) CLUTWR_Bits { return CLUTWR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CLUTWR) StoreBits(mask, b CLUTWR_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CLUTWR) SetBits(mask CLUTWR_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *CLUTWR) ClearBits(mask CLUTWR_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *CLUTWR) Load() CLUTWR_Bits                 { return CLUTWR_Bits(r.U32.Load()) }
func (r *CLUTWR) Store(b CLUTWR_Bits)               { r.U32.Store(uint32(b)) }

func (r *CLUTWR) AtomicStoreBits(mask, b CLUTWR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *CLUTWR) AtomicSetBits(mask CLUTWR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *CLUTWR) AtomicClearBits(mask CLUTWR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type CLUTWR_Mask struct{ mmio.UM32 }

func (rm CLUTWR_Mask) Load() CLUTWR_Bits   { return CLUTWR_Bits(rm.UM32.Load()) }
func (rm CLUTWR_Mask) Store(b CLUTWR_Bits) { rm.UM32.Store(uint32(b)) }
