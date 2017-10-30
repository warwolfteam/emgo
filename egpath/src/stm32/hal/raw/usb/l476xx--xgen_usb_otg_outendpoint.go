// +build l476xx

package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_OUTEndpoint_Periph struct {
	DOEPCTL  DOEPCTL
	_        uint32
	DOEPINT  DOEPINT
	_        uint32
	DOEPTSIZ DOEPTSIZ
	DOEPDMA  DOEPDMA
}

func (p *USB_OTG_OUTEndpoint_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DOEPCTL_Bits uint32

func (b DOEPCTL_Bits) Field(mask DOEPCTL_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPCTL_Bits) J(v int) DOEPCTL_Bits {
	return DOEPCTL_Bits(bits.Make32(v, uint32(mask)))
}

type DOEPCTL struct{ mmio.U32 }

func (r *DOEPCTL) Bits(mask DOEPCTL_Bits) DOEPCTL_Bits { return DOEPCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOEPCTL) StoreBits(mask, b DOEPCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPCTL) SetBits(mask DOEPCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DOEPCTL) ClearBits(mask DOEPCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPCTL) Load() DOEPCTL_Bits                  { return DOEPCTL_Bits(r.U32.Load()) }
func (r *DOEPCTL) Store(b DOEPCTL_Bits)                { r.U32.Store(uint32(b)) }

func (r *DOEPCTL) AtomicStoreBits(mask, b DOEPCTL_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DOEPCTL) AtomicSetBits(mask DOEPCTL_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOEPCTL) AtomicClearBits(mask DOEPCTL_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOEPCTL_Mask struct{ mmio.UM32 }

func (rm DOEPCTL_Mask) Load() DOEPCTL_Bits   { return DOEPCTL_Bits(rm.UM32.Load()) }
func (rm DOEPCTL_Mask) Store(b DOEPCTL_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) MPSIZ() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) USBAEP() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(USBAEP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) NAKSTS() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(NAKSTS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SD0PID_SEVNFRM() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(SD0PID_SEVNFRM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SODDFRM() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(SODDFRM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPTYP() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SNPM() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(SNPM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STALL() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(STALL)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) CNAK() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(CNAK)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SNAK() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(SNAK)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPDIS() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(EPDIS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPENA() DOEPCTL_Mask {
	return DOEPCTL_Mask{mmio.UM32{&p.DOEPCTL.U32, uint32(EPENA)}}
}

type DOEPINT_Bits uint32

func (b DOEPINT_Bits) Field(mask DOEPINT_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPINT_Bits) J(v int) DOEPINT_Bits {
	return DOEPINT_Bits(bits.Make32(v, uint32(mask)))
}

type DOEPINT struct{ mmio.U32 }

func (r *DOEPINT) Bits(mask DOEPINT_Bits) DOEPINT_Bits { return DOEPINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOEPINT) StoreBits(mask, b DOEPINT_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPINT) SetBits(mask DOEPINT_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DOEPINT) ClearBits(mask DOEPINT_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPINT) Load() DOEPINT_Bits                  { return DOEPINT_Bits(r.U32.Load()) }
func (r *DOEPINT) Store(b DOEPINT_Bits)                { r.U32.Store(uint32(b)) }

func (r *DOEPINT) AtomicStoreBits(mask, b DOEPINT_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DOEPINT) AtomicSetBits(mask DOEPINT_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOEPINT) AtomicClearBits(mask DOEPINT_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOEPINT_Mask struct{ mmio.UM32 }

func (rm DOEPINT_Mask) Load() DOEPINT_Bits   { return DOEPINT_Bits(rm.UM32.Load()) }
func (rm DOEPINT_Mask) Store(b DOEPINT_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) XFRC() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPDISD() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(EPDISD)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STUP() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(STUP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) OTEPDIS() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(OTEPDIS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) B2BSTUP() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(B2BSTUP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) NYET() DOEPINT_Mask {
	return DOEPINT_Mask{mmio.UM32{&p.DOEPINT.U32, uint32(NYET)}}
}

type DOEPTSIZ_Bits uint32

func (b DOEPTSIZ_Bits) Field(mask DOEPTSIZ_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPTSIZ_Bits) J(v int) DOEPTSIZ_Bits {
	return DOEPTSIZ_Bits(bits.Make32(v, uint32(mask)))
}

type DOEPTSIZ struct{ mmio.U32 }

func (r *DOEPTSIZ) Bits(mask DOEPTSIZ_Bits) DOEPTSIZ_Bits {
	return DOEPTSIZ_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DOEPTSIZ) StoreBits(mask, b DOEPTSIZ_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPTSIZ) SetBits(mask DOEPTSIZ_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DOEPTSIZ) ClearBits(mask DOEPTSIZ_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPTSIZ) Load() DOEPTSIZ_Bits             { return DOEPTSIZ_Bits(r.U32.Load()) }
func (r *DOEPTSIZ) Store(b DOEPTSIZ_Bits)           { r.U32.Store(uint32(b)) }

func (r *DOEPTSIZ) AtomicStoreBits(mask, b DOEPTSIZ_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DOEPTSIZ) AtomicSetBits(mask DOEPTSIZ_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOEPTSIZ) AtomicClearBits(mask DOEPTSIZ_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOEPTSIZ_Mask struct{ mmio.UM32 }

func (rm DOEPTSIZ_Mask) Load() DOEPTSIZ_Bits   { return DOEPTSIZ_Bits(rm.UM32.Load()) }
func (rm DOEPTSIZ_Mask) Store(b DOEPTSIZ_Bits) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) XFRSIZ() DOEPTSIZ_Mask {
	return DOEPTSIZ_Mask{mmio.UM32{&p.DOEPTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) PKTCNT() DOEPTSIZ_Mask {
	return DOEPTSIZ_Mask{mmio.UM32{&p.DOEPTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STUPCNT() DOEPTSIZ_Mask {
	return DOEPTSIZ_Mask{mmio.UM32{&p.DOEPTSIZ.U32, uint32(STUPCNT)}}
}

type DOEPDMA_Bits uint32

func (b DOEPDMA_Bits) Field(mask DOEPDMA_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPDMA_Bits) J(v int) DOEPDMA_Bits {
	return DOEPDMA_Bits(bits.Make32(v, uint32(mask)))
}

type DOEPDMA struct{ mmio.U32 }

func (r *DOEPDMA) Bits(mask DOEPDMA_Bits) DOEPDMA_Bits { return DOEPDMA_Bits(r.U32.Bits(uint32(mask))) }
func (r *DOEPDMA) StoreBits(mask, b DOEPDMA_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DOEPDMA) SetBits(mask DOEPDMA_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DOEPDMA) ClearBits(mask DOEPDMA_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DOEPDMA) Load() DOEPDMA_Bits                  { return DOEPDMA_Bits(r.U32.Load()) }
func (r *DOEPDMA) Store(b DOEPDMA_Bits)                { r.U32.Store(uint32(b)) }

func (r *DOEPDMA) AtomicStoreBits(mask, b DOEPDMA_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *DOEPDMA) AtomicSetBits(mask DOEPDMA_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *DOEPDMA) AtomicClearBits(mask DOEPDMA_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type DOEPDMA_Mask struct{ mmio.UM32 }

func (rm DOEPDMA_Mask) Load() DOEPDMA_Bits   { return DOEPDMA_Bits(rm.UM32.Load()) }
func (rm DOEPDMA_Mask) Store(b DOEPDMA_Bits) { rm.UM32.Store(uint32(b)) }
