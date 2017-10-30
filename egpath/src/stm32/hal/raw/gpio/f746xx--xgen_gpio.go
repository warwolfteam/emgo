// +build f746xx

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type GPIO_Periph struct {
	MODER   MODER
	OTYPER  OTYPER
	OSPEEDR OSPEEDR
	PUPDR   PUPDR
	IDR     IDR
	ODR     ODR
	BSRR    BSRR
	LCKR    LCKR
	AFR     [2]AFR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

//emgo:const
var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

//emgo:const
var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

//emgo:const
var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

//emgo:const
var GPIOE = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOE_BASE)))

//emgo:const
var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

//emgo:const
var GPIOG = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOG_BASE)))

//emgo:const
var GPIOH = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOH_BASE)))

//emgo:const
var GPIOI = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOI_BASE)))

//emgo:const
var GPIOJ = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOJ_BASE)))

//emgo:const
var GPIOK = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOK_BASE)))

type MODER_Bits uint32

func (b MODER_Bits) Field(mask MODER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MODER_Bits) J(v int) MODER_Bits {
	return MODER_Bits(bits.Make32(v, uint32(mask)))
}

type MODER struct{ mmio.U32 }

func (r *MODER) Bits(mask MODER_Bits) MODER_Bits { return MODER_Bits(r.U32.Bits(uint32(mask))) }
func (r *MODER) StoreBits(mask, b MODER_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MODER) SetBits(mask MODER_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *MODER) ClearBits(mask MODER_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *MODER) Load() MODER_Bits                { return MODER_Bits(r.U32.Load()) }
func (r *MODER) Store(b MODER_Bits)              { r.U32.Store(uint32(b)) }

func (r *MODER) AtomicStoreBits(mask, b MODER_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *MODER) AtomicSetBits(mask MODER_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *MODER) AtomicClearBits(mask MODER_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type MODER_Mask struct{ mmio.UM32 }

func (rm MODER_Mask) Load() MODER_Bits   { return MODER_Bits(rm.UM32.Load()) }
func (rm MODER_Mask) Store(b MODER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODER0() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER0)}}
}

func (p *GPIO_Periph) MODER1() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER1)}}
}

func (p *GPIO_Periph) MODER2() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER2)}}
}

func (p *GPIO_Periph) MODER3() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER3)}}
}

func (p *GPIO_Periph) MODER4() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER4)}}
}

func (p *GPIO_Periph) MODER5() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER5)}}
}

func (p *GPIO_Periph) MODER6() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER6)}}
}

func (p *GPIO_Periph) MODER7() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER7)}}
}

func (p *GPIO_Periph) MODER8() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER8)}}
}

func (p *GPIO_Periph) MODER9() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER9)}}
}

func (p *GPIO_Periph) MODER10() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER10)}}
}

func (p *GPIO_Periph) MODER11() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER11)}}
}

func (p *GPIO_Periph) MODER12() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER12)}}
}

func (p *GPIO_Periph) MODER13() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER13)}}
}

func (p *GPIO_Periph) MODER14() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER14)}}
}

func (p *GPIO_Periph) MODER15() MODER_Mask {
	return MODER_Mask{mmio.UM32{&p.MODER.U32, uint32(MODER15)}}
}

type OTYPER_Bits uint32

func (b OTYPER_Bits) Field(mask OTYPER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTYPER_Bits) J(v int) OTYPER_Bits {
	return OTYPER_Bits(bits.Make32(v, uint32(mask)))
}

type OTYPER struct{ mmio.U32 }

func (r *OTYPER) Bits(mask OTYPER_Bits) OTYPER_Bits { return OTYPER_Bits(r.U32.Bits(uint32(mask))) }
func (r *OTYPER) StoreBits(mask, b OTYPER_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OTYPER) SetBits(mask OTYPER_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *OTYPER) ClearBits(mask OTYPER_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *OTYPER) Load() OTYPER_Bits                 { return OTYPER_Bits(r.U32.Load()) }
func (r *OTYPER) Store(b OTYPER_Bits)               { r.U32.Store(uint32(b)) }

func (r *OTYPER) AtomicStoreBits(mask, b OTYPER_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *OTYPER) AtomicSetBits(mask OTYPER_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OTYPER) AtomicClearBits(mask OTYPER_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type OTYPER_Mask struct{ mmio.UM32 }

func (rm OTYPER_Mask) Load() OTYPER_Bits   { return OTYPER_Bits(rm.UM32.Load()) }
func (rm OTYPER_Mask) Store(b OTYPER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT_0() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_0)}}
}

func (p *GPIO_Periph) OT_1() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_1)}}
}

func (p *GPIO_Periph) OT_2() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_2)}}
}

func (p *GPIO_Periph) OT_3() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_3)}}
}

func (p *GPIO_Periph) OT_4() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_4)}}
}

func (p *GPIO_Periph) OT_5() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_5)}}
}

func (p *GPIO_Periph) OT_6() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_6)}}
}

func (p *GPIO_Periph) OT_7() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_7)}}
}

func (p *GPIO_Periph) OT_8() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_8)}}
}

func (p *GPIO_Periph) OT_9() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_9)}}
}

func (p *GPIO_Periph) OT_10() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_10)}}
}

func (p *GPIO_Periph) OT_11() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_11)}}
}

func (p *GPIO_Periph) OT_12() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_12)}}
}

func (p *GPIO_Periph) OT_13() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_13)}}
}

func (p *GPIO_Periph) OT_14() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_14)}}
}

func (p *GPIO_Periph) OT_15() OTYPER_Mask {
	return OTYPER_Mask{mmio.UM32{&p.OTYPER.U32, uint32(OT_15)}}
}

type OSPEEDR_Bits uint32

func (b OSPEEDR_Bits) Field(mask OSPEEDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OSPEEDR_Bits) J(v int) OSPEEDR_Bits {
	return OSPEEDR_Bits(bits.Make32(v, uint32(mask)))
}

type OSPEEDR struct{ mmio.U32 }

func (r *OSPEEDR) Bits(mask OSPEEDR_Bits) OSPEEDR_Bits { return OSPEEDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *OSPEEDR) StoreBits(mask, b OSPEEDR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *OSPEEDR) SetBits(mask OSPEEDR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *OSPEEDR) ClearBits(mask OSPEEDR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *OSPEEDR) Load() OSPEEDR_Bits                  { return OSPEEDR_Bits(r.U32.Load()) }
func (r *OSPEEDR) Store(b OSPEEDR_Bits)                { r.U32.Store(uint32(b)) }

func (r *OSPEEDR) AtomicStoreBits(mask, b OSPEEDR_Bits) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *OSPEEDR) AtomicSetBits(mask OSPEEDR_Bits)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *OSPEEDR) AtomicClearBits(mask OSPEEDR_Bits) { r.U32.AtomicClearBits(uint32(mask)) }

type OSPEEDR_Mask struct{ mmio.UM32 }

func (rm OSPEEDR_Mask) Load() OSPEEDR_Bits   { return OSPEEDR_Bits(rm.UM32.Load()) }
func (rm OSPEEDR_Mask) Store(b OSPEEDR_Bits) { rm.UM32.Store(uint32(b)) }

type PUPDR_Bits uint32

func (b PUPDR_Bits) Field(mask PUPDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PUPDR_Bits) J(v int) PUPDR_Bits {
	return PUPDR_Bits(bits.Make32(v, uint32(mask)))
}

type PUPDR struct{ mmio.U32 }

func (r *PUPDR) Bits(mask PUPDR_Bits) PUPDR_Bits { return PUPDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *PUPDR) StoreBits(mask, b PUPDR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *PUPDR) SetBits(mask PUPDR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *PUPDR) ClearBits(mask PUPDR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *PUPDR) Load() PUPDR_Bits                { return PUPDR_Bits(r.U32.Load()) }
func (r *PUPDR) Store(b PUPDR_Bits)              { r.U32.Store(uint32(b)) }

func (r *PUPDR) AtomicStoreBits(mask, b PUPDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *PUPDR) AtomicSetBits(mask PUPDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *PUPDR) AtomicClearBits(mask PUPDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type PUPDR_Mask struct{ mmio.UM32 }

func (rm PUPDR_Mask) Load() PUPDR_Bits   { return PUPDR_Bits(rm.UM32.Load()) }
func (rm PUPDR_Mask) Store(b PUPDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPDR0() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR0)}}
}

func (p *GPIO_Periph) PUPDR1() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR1)}}
}

func (p *GPIO_Periph) PUPDR2() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR2)}}
}

func (p *GPIO_Periph) PUPDR3() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR3)}}
}

func (p *GPIO_Periph) PUPDR4() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR4)}}
}

func (p *GPIO_Periph) PUPDR5() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR5)}}
}

func (p *GPIO_Periph) PUPDR6() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR6)}}
}

func (p *GPIO_Periph) PUPDR7() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR7)}}
}

func (p *GPIO_Periph) PUPDR8() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR8)}}
}

func (p *GPIO_Periph) PUPDR9() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR9)}}
}

func (p *GPIO_Periph) PUPDR10() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR10)}}
}

func (p *GPIO_Periph) PUPDR11() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR11)}}
}

func (p *GPIO_Periph) PUPDR12() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR12)}}
}

func (p *GPIO_Periph) PUPDR13() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR13)}}
}

func (p *GPIO_Periph) PUPDR14() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR14)}}
}

func (p *GPIO_Periph) PUPDR15() PUPDR_Mask {
	return PUPDR_Mask{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR15)}}
}

type IDR_Bits uint32

func (b IDR_Bits) Field(mask IDR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR_Bits) J(v int) IDR_Bits {
	return IDR_Bits(bits.Make32(v, uint32(mask)))
}

type IDR struct{ mmio.U32 }

func (r *IDR) Bits(mask IDR_Bits) IDR_Bits { return IDR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IDR) StoreBits(mask, b IDR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IDR) SetBits(mask IDR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IDR) ClearBits(mask IDR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IDR) Load() IDR_Bits              { return IDR_Bits(r.U32.Load()) }
func (r *IDR) Store(b IDR_Bits)            { r.U32.Store(uint32(b)) }

func (r *IDR) AtomicStoreBits(mask, b IDR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *IDR) AtomicSetBits(mask IDR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *IDR) AtomicClearBits(mask IDR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type IDR_Mask struct{ mmio.UM32 }

func (rm IDR_Mask) Load() IDR_Bits   { return IDR_Bits(rm.UM32.Load()) }
func (rm IDR_Mask) Store(b IDR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) IDR_0() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_0)}}
}

func (p *GPIO_Periph) IDR_1() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_1)}}
}

func (p *GPIO_Periph) IDR_2() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_2)}}
}

func (p *GPIO_Periph) IDR_3() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_3)}}
}

func (p *GPIO_Periph) IDR_4() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_4)}}
}

func (p *GPIO_Periph) IDR_5() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_5)}}
}

func (p *GPIO_Periph) IDR_6() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_6)}}
}

func (p *GPIO_Periph) IDR_7() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_7)}}
}

func (p *GPIO_Periph) IDR_8() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_8)}}
}

func (p *GPIO_Periph) IDR_9() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_9)}}
}

func (p *GPIO_Periph) IDR_10() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_10)}}
}

func (p *GPIO_Periph) IDR_11() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_11)}}
}

func (p *GPIO_Periph) IDR_12() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_12)}}
}

func (p *GPIO_Periph) IDR_13() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_13)}}
}

func (p *GPIO_Periph) IDR_14() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_14)}}
}

func (p *GPIO_Periph) IDR_15() IDR_Mask {
	return IDR_Mask{mmio.UM32{&p.IDR.U32, uint32(IDR_15)}}
}

type ODR_Bits uint32

func (b ODR_Bits) Field(mask ODR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR_Bits) J(v int) ODR_Bits {
	return ODR_Bits(bits.Make32(v, uint32(mask)))
}

type ODR struct{ mmio.U32 }

func (r *ODR) Bits(mask ODR_Bits) ODR_Bits { return ODR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ODR) StoreBits(mask, b ODR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ODR) SetBits(mask ODR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ODR) ClearBits(mask ODR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ODR) Load() ODR_Bits              { return ODR_Bits(r.U32.Load()) }
func (r *ODR) Store(b ODR_Bits)            { r.U32.Store(uint32(b)) }

func (r *ODR) AtomicStoreBits(mask, b ODR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ODR) AtomicSetBits(mask ODR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ODR) AtomicClearBits(mask ODR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type ODR_Mask struct{ mmio.UM32 }

func (rm ODR_Mask) Load() ODR_Bits   { return ODR_Bits(rm.UM32.Load()) }
func (rm ODR_Mask) Store(b ODR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) ODR_0() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_0)}}
}

func (p *GPIO_Periph) ODR_1() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_1)}}
}

func (p *GPIO_Periph) ODR_2() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_2)}}
}

func (p *GPIO_Periph) ODR_3() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_3)}}
}

func (p *GPIO_Periph) ODR_4() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_4)}}
}

func (p *GPIO_Periph) ODR_5() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_5)}}
}

func (p *GPIO_Periph) ODR_6() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_6)}}
}

func (p *GPIO_Periph) ODR_7() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_7)}}
}

func (p *GPIO_Periph) ODR_8() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_8)}}
}

func (p *GPIO_Periph) ODR_9() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_9)}}
}

func (p *GPIO_Periph) ODR_10() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_10)}}
}

func (p *GPIO_Periph) ODR_11() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_11)}}
}

func (p *GPIO_Periph) ODR_12() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_12)}}
}

func (p *GPIO_Periph) ODR_13() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_13)}}
}

func (p *GPIO_Periph) ODR_14() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_14)}}
}

func (p *GPIO_Periph) ODR_15() ODR_Mask {
	return ODR_Mask{mmio.UM32{&p.ODR.U32, uint32(ODR_15)}}
}

type BSRR_Bits uint32

func (b BSRR_Bits) Field(mask BSRR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRR_Bits) J(v int) BSRR_Bits {
	return BSRR_Bits(bits.Make32(v, uint32(mask)))
}

type BSRR struct{ mmio.U32 }

func (r *BSRR) Bits(mask BSRR_Bits) BSRR_Bits { return BSRR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BSRR) StoreBits(mask, b BSRR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BSRR) SetBits(mask BSRR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *BSRR) ClearBits(mask BSRR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *BSRR) Load() BSRR_Bits               { return BSRR_Bits(r.U32.Load()) }
func (r *BSRR) Store(b BSRR_Bits)             { r.U32.Store(uint32(b)) }

func (r *BSRR) AtomicStoreBits(mask, b BSRR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *BSRR) AtomicSetBits(mask BSRR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *BSRR) AtomicClearBits(mask BSRR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type BSRR_Mask struct{ mmio.UM32 }

func (rm BSRR_Mask) Load() BSRR_Bits   { return BSRR_Bits(rm.UM32.Load()) }
func (rm BSRR_Mask) Store(b BSRR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BS_0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_0)}}
}

func (p *GPIO_Periph) BS_1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_1)}}
}

func (p *GPIO_Periph) BS_2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_2)}}
}

func (p *GPIO_Periph) BS_3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_3)}}
}

func (p *GPIO_Periph) BS_4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_4)}}
}

func (p *GPIO_Periph) BS_5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_5)}}
}

func (p *GPIO_Periph) BS_6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_6)}}
}

func (p *GPIO_Periph) BS_7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_7)}}
}

func (p *GPIO_Periph) BS_8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_8)}}
}

func (p *GPIO_Periph) BS_9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_9)}}
}

func (p *GPIO_Periph) BS_10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_10)}}
}

func (p *GPIO_Periph) BS_11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_11)}}
}

func (p *GPIO_Periph) BS_12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_12)}}
}

func (p *GPIO_Periph) BS_13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_13)}}
}

func (p *GPIO_Periph) BS_14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_14)}}
}

func (p *GPIO_Periph) BS_15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BS_15)}}
}

func (p *GPIO_Periph) BR_0() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_0)}}
}

func (p *GPIO_Periph) BR_1() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_1)}}
}

func (p *GPIO_Periph) BR_2() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_2)}}
}

func (p *GPIO_Periph) BR_3() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_3)}}
}

func (p *GPIO_Periph) BR_4() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_4)}}
}

func (p *GPIO_Periph) BR_5() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_5)}}
}

func (p *GPIO_Periph) BR_6() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_6)}}
}

func (p *GPIO_Periph) BR_7() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_7)}}
}

func (p *GPIO_Periph) BR_8() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_8)}}
}

func (p *GPIO_Periph) BR_9() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_9)}}
}

func (p *GPIO_Periph) BR_10() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_10)}}
}

func (p *GPIO_Periph) BR_11() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_11)}}
}

func (p *GPIO_Periph) BR_12() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_12)}}
}

func (p *GPIO_Periph) BR_13() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_13)}}
}

func (p *GPIO_Periph) BR_14() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_14)}}
}

func (p *GPIO_Periph) BR_15() BSRR_Mask {
	return BSRR_Mask{mmio.UM32{&p.BSRR.U32, uint32(BR_15)}}
}

type LCKR_Bits uint32

func (b LCKR_Bits) Field(mask LCKR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR_Bits) J(v int) LCKR_Bits {
	return LCKR_Bits(bits.Make32(v, uint32(mask)))
}

type LCKR struct{ mmio.U32 }

func (r *LCKR) Bits(mask LCKR_Bits) LCKR_Bits { return LCKR_Bits(r.U32.Bits(uint32(mask))) }
func (r *LCKR) StoreBits(mask, b LCKR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *LCKR) SetBits(mask LCKR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *LCKR) ClearBits(mask LCKR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *LCKR) Load() LCKR_Bits               { return LCKR_Bits(r.U32.Load()) }
func (r *LCKR) Store(b LCKR_Bits)             { r.U32.Store(uint32(b)) }

func (r *LCKR) AtomicStoreBits(mask, b LCKR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *LCKR) AtomicSetBits(mask LCKR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *LCKR) AtomicClearBits(mask LCKR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type LCKR_Mask struct{ mmio.UM32 }

func (rm LCKR_Mask) Load() LCKR_Bits   { return LCKR_Bits(rm.UM32.Load()) }
func (rm LCKR_Mask) Store(b LCKR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) LCK0() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *GPIO_Periph) LCK1() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *GPIO_Periph) LCK2() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *GPIO_Periph) LCK3() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *GPIO_Periph) LCK4() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *GPIO_Periph) LCK5() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *GPIO_Periph) LCK6() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *GPIO_Periph) LCK7() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *GPIO_Periph) LCK8() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *GPIO_Periph) LCK9() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *GPIO_Periph) LCK10() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *GPIO_Periph) LCK11() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *GPIO_Periph) LCK12() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *GPIO_Periph) LCK13() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *GPIO_Periph) LCK14() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *GPIO_Periph) LCK15() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *GPIO_Periph) LCKK() LCKR_Mask {
	return LCKR_Mask{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}

type AFR_Bits uint32

func (b AFR_Bits) Field(mask AFR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFR_Bits) J(v int) AFR_Bits {
	return AFR_Bits(bits.Make32(v, uint32(mask)))
}

type AFR struct{ mmio.U32 }

func (r *AFR) Bits(mask AFR_Bits) AFR_Bits { return AFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *AFR) StoreBits(mask, b AFR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *AFR) SetBits(mask AFR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *AFR) ClearBits(mask AFR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *AFR) Load() AFR_Bits              { return AFR_Bits(r.U32.Load()) }
func (r *AFR) Store(b AFR_Bits)            { r.U32.Store(uint32(b)) }

func (r *AFR) AtomicStoreBits(mask, b AFR_Bits) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *AFR) AtomicSetBits(mask AFR_Bits)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *AFR) AtomicClearBits(mask AFR_Bits)    { r.U32.AtomicClearBits(uint32(mask)) }

type AFR_Mask struct{ mmio.UM32 }

func (rm AFR_Mask) Load() AFR_Bits   { return AFR_Bits(rm.UM32.Load()) }
func (rm AFR_Mask) Store(b AFR_Bits) { rm.UM32.Store(uint32(b)) }
