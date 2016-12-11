// +build f746xx

// Peripheral: USB_OTG_Host_Periph  USB_OTG_Host_Mode_Register_Structures.
// Instances:
// Registers:
//  0x00 32  HCFG     Host Configuration Register          400h.
//  0x04 32  HFIR     Host Frame Interval Register         404h.
//  0x08 32  HFNUM    Host Frame Nbr/Frame Remaining       408h.
//  0x10 32  HPTXSTS  Host Periodic Tx FIFO/ Queue Status  410h.
//  0x14 32  HAINT    Host All Channels Interrupt Register 414h.
//  0x18 32  HAINTMSK Host All Channels Interrupt Mask     418h.
// Import:
//  stm32/o/f746xx/mmap
package usb

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	FSLSPCS   HCFG_Bits = 0x03 << 0 //+ FS/LS PHY clock select.
	FSLSPCS_0 HCFG_Bits = 0x01 << 0 //  Bit 0.
	FSLSPCS_1 HCFG_Bits = 0x02 << 0 //  Bit 1.
	FSLSS     HCFG_Bits = 0x01 << 2 //+ FS- and LS-only support.
)

const (
	FSLSPCSn = 0
	FSLSSn   = 2
)

const (
	FRIVL HFIR_Bits = 0xFFFF << 0 //+ Frame interval.
)

const (
	FRIVLn = 0
)

const (
	FRNUM HFNUM_Bits = 0xFFFF << 0  //+ Frame number.
	FTREM HFNUM_Bits = 0xFFFF << 16 //+ Frame time remaining.
)

const (
	FRNUMn = 0
	FTREMn = 16
)

const (
	PTXFSAVL  HPTXSTS_Bits = 0xFFFF << 0 //+ Periodic transmit data FIFO space available.
	PTXQSAV   HPTXSTS_Bits = 0xFF << 16  //+ Periodic transmit request queue space available.
	PTXQSAV_0 HPTXSTS_Bits = 0x01 << 16  //  Bit 0.
	PTXQSAV_1 HPTXSTS_Bits = 0x02 << 16  //  Bit 1.
	PTXQSAV_2 HPTXSTS_Bits = 0x04 << 16  //  Bit 2.
	PTXQSAV_3 HPTXSTS_Bits = 0x08 << 16  //  Bit 3.
	PTXQSAV_4 HPTXSTS_Bits = 0x10 << 16  //  Bit 4.
	PTXQSAV_5 HPTXSTS_Bits = 0x20 << 16  //  Bit 5.
	PTXQSAV_6 HPTXSTS_Bits = 0x40 << 16  //  Bit 6.
	PTXQSAV_7 HPTXSTS_Bits = 0x80 << 16  //  Bit 7.
	PTXQTOP   HPTXSTS_Bits = 0xFF << 24  //+ Top of the periodic transmit request queue.
	PTXQTOP_0 HPTXSTS_Bits = 0x01 << 24  //  Bit 0.
	PTXQTOP_1 HPTXSTS_Bits = 0x02 << 24  //  Bit 1.
	PTXQTOP_2 HPTXSTS_Bits = 0x04 << 24  //  Bit 2.
	PTXQTOP_3 HPTXSTS_Bits = 0x08 << 24  //  Bit 3.
	PTXQTOP_4 HPTXSTS_Bits = 0x10 << 24  //  Bit 4.
	PTXQTOP_5 HPTXSTS_Bits = 0x20 << 24  //  Bit 5.
	PTXQTOP_6 HPTXSTS_Bits = 0x40 << 24  //  Bit 6.
	PTXQTOP_7 HPTXSTS_Bits = 0x80 << 24  //  Bit 7.
)

const (
	PTXFSAVLn = 0
	PTXQSAVn  = 16
	PTXQTOPn  = 24
)

const (
	HAINTM HAINTMSK_Bits = 0xFFFF << 0 //+ Channel interrupt mask.
)

const (
	HAINTMn = 0
)