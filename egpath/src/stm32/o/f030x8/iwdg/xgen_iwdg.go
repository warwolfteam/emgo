package iwdg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type IWDG_Periph struct {
	KR   RKR
	PR   RPR
	RLR  RRLR
	SR   RSR
	WINR RWINR
}

func (p *IWDG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var IWDG = (*IWDG_Periph)(unsafe.Pointer(uintptr(mmap.IWDG_BASE)))

type KR uint32

func (b KR) Field(mask KR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KR) J(v int) KR {
	return KR(bits.Make32(v, uint32(mask)))
}

type RKR struct{ mmio.U32 }

func (r *RKR) Bits(mask KR) KR      { return KR(r.U32.Bits(uint32(mask))) }
func (r *RKR) StoreBits(mask, b KR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RKR) SetBits(mask KR)      { r.U32.SetBits(uint32(mask)) }
func (r *RKR) ClearBits(mask KR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RKR) Load() KR             { return KR(r.U32.Load()) }
func (r *RKR) Store(b KR)           { r.U32.Store(uint32(b)) }

func (r *RKR) AtomicStoreBits(mask, b KR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RKR) AtomicSetBits(mask KR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RKR) AtomicClearBits(mask KR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMKR struct{ mmio.UM32 }

func (rm RMKR) Load() KR   { return KR(rm.UM32.Load()) }
func (rm RMKR) Store(b KR) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) KEY() RMKR {
	return RMKR{mmio.UM32{&p.KR.U32, uint32(KEY)}}
}

type PR uint32

func (b PR) Field(mask PR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR) J(v int) PR {
	return PR(bits.Make32(v, uint32(mask)))
}

type RPR struct{ mmio.U32 }

func (r *RPR) Bits(mask PR) PR      { return PR(r.U32.Bits(uint32(mask))) }
func (r *RPR) StoreBits(mask, b PR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPR) SetBits(mask PR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPR) ClearBits(mask PR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPR) Load() PR             { return PR(r.U32.Load()) }
func (r *RPR) Store(b PR)           { r.U32.Store(uint32(b)) }

func (r *RPR) AtomicStoreBits(mask, b PR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPR) AtomicSetBits(mask PR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPR) AtomicClearBits(mask PR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPR struct{ mmio.UM32 }

func (rm RMPR) Load() PR   { return PR(rm.UM32.Load()) }
func (rm RMPR) Store(b PR) { rm.UM32.Store(uint32(b)) }

type RLR uint32

func (b RLR) Field(mask RLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RLR) J(v int) RLR {
	return RLR(bits.Make32(v, uint32(mask)))
}

type RRLR struct{ mmio.U32 }

func (r *RRLR) Bits(mask RLR) RLR     { return RLR(r.U32.Bits(uint32(mask))) }
func (r *RRLR) StoreBits(mask, b RLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRLR) SetBits(mask RLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRLR) ClearBits(mask RLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRLR) Load() RLR             { return RLR(r.U32.Load()) }
func (r *RRLR) Store(b RLR)           { r.U32.Store(uint32(b)) }

func (r *RRLR) AtomicStoreBits(mask, b RLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRLR) AtomicSetBits(mask RLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRLR) AtomicClearBits(mask RLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRLR struct{ mmio.UM32 }

func (rm RMRLR) Load() RLR   { return RLR(rm.UM32.Load()) }
func (rm RMRLR) Store(b RLR) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) RL() RMRLR {
	return RMRLR{mmio.UM32{&p.RLR.U32, uint32(RL)}}
}

type SR uint32

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.Make32(v, uint32(mask)))
}

type RSR struct{ mmio.U32 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U32.Bits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

func (r *RSR) AtomicStoreBits(mask, b SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR) AtomicSetBits(mask SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR) AtomicClearBits(mask SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) PVU() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PVU)}}
}

func (p *IWDG_Periph) RVU() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RVU)}}
}

func (p *IWDG_Periph) WVU() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WVU)}}
}

type WINR uint32

func (b WINR) Field(mask WINR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WINR) J(v int) WINR {
	return WINR(bits.Make32(v, uint32(mask)))
}

type RWINR struct{ mmio.U32 }

func (r *RWINR) Bits(mask WINR) WINR    { return WINR(r.U32.Bits(uint32(mask))) }
func (r *RWINR) StoreBits(mask, b WINR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWINR) SetBits(mask WINR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWINR) ClearBits(mask WINR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWINR) Load() WINR             { return WINR(r.U32.Load()) }
func (r *RWINR) Store(b WINR)           { r.U32.Store(uint32(b)) }

func (r *RWINR) AtomicStoreBits(mask, b WINR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWINR) AtomicSetBits(mask WINR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWINR) AtomicClearBits(mask WINR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWINR struct{ mmio.UM32 }

func (rm RMWINR) Load() WINR   { return WINR(rm.UM32.Load()) }
func (rm RMWINR) Store(b WINR) { rm.UM32.Store(uint32(b)) }

func (p *IWDG_Periph) WIN() RMWINR {
	return RMWINR{mmio.UM32{&p.WINR.U32, uint32(WIN)}}
}