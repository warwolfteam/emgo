// Peripheral: ADC_Common_Periph  Analog to Digital Converter.
// Instances:
//  ADC12_COMMON  mmap.ADC1_2_COMMON_BASE
//  ADC34_COMMON  mmap.ADC3_4_COMMON_BASE
// Registers:
//  0x00 32  CSR ADC Common status register.
//  0x08 32  CCR ADC common control register.
//  0x0C 32  CDR ADC common regular data register for dual.
// Import:
//  stm32/o/f303xe/mmap
package adc

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	ADRDY_MST       CSR_Bits = 0x01 << 0  //+ Master ADC ready.
	ADRDY_EOSMP_MST CSR_Bits = 0x01 << 1  //+ End of sampling phase flag of the master ADC.
	ADRDY_EOC_MST   CSR_Bits = 0x01 << 2  //+ End of regular conversion of the master ADC.
	ADRDY_EOS_MST   CSR_Bits = 0x01 << 3  //+ End of regular sequence flag of the master ADC.
	ADRDY_OVR_MST   CSR_Bits = 0x01 << 4  //+ Overrun flag of the master ADC.
	ADRDY_JEOC_MST  CSR_Bits = 0x01 << 5  //+ End of injected conversion of the master ADC.
	ADRDY_JEOS_MST  CSR_Bits = 0x01 << 6  //+ End of injected sequence flag of the master ADC.
	AWD1_MST        CSR_Bits = 0x01 << 7  //+ Analog watchdog 1 flag of the master ADC.
	AWD2_MST        CSR_Bits = 0x01 << 8  //+ Analog watchdog 2 flag of the master ADC.
	AWD3_MST        CSR_Bits = 0x01 << 9  //+ Analog watchdog 3 flag of the master ADC.
	JQOVF_MST       CSR_Bits = 0x01 << 10 //+ Injected context queue overflow flag of the master ADC.
	ADRDY_SLV       CSR_Bits = 0x01 << 16 //+ Slave ADC ready.
	ADRDY_EOSMP_SLV CSR_Bits = 0x01 << 17 //+ End of sampling phase flag of the slave ADC.
	ADRDY_EOC_SLV   CSR_Bits = 0x01 << 18 //+ End of regular conversion of the slave ADC.
	ADRDY_EOS_SLV   CSR_Bits = 0x01 << 19 //+ End of regular sequence flag of the slave ADC.
	ADRDY_OVR_SLV   CSR_Bits = 0x01 << 20 //+ Overrun flag of the slave ADC.
	ADRDY_JEOC_SLV  CSR_Bits = 0x01 << 21 //+ End of injected conversion of the slave ADC.
	ADRDY_JEOS_SLV  CSR_Bits = 0x01 << 22 //+ End of injected sequence flag of the slave ADC.
	AWD1_SLV        CSR_Bits = 0x01 << 23 //+ Analog watchdog 1 flag of the slave ADC.
	AWD2_SLV        CSR_Bits = 0x01 << 24 //+ Analog watchdog 2 flag of the slave ADC.
	AWD3_SLV        CSR_Bits = 0x01 << 25 //+ Analog watchdog 3 flag of the slave ADC.
	JQOVF_SLV       CSR_Bits = 0x01 << 26 //+ Injected context queue overflow flag of the slave ADC.
	EOSMP_MST       CSR_Bits = 0x01 << 1  //  ADC multimode master group regular end of sampling flag.
	EOC_MST         CSR_Bits = 0x01 << 2  //  ADC multimode master group regular end of unitary conversion flag.
	EOS_MST         CSR_Bits = 0x01 << 3  //  ADC multimode master group regular end of sequence conversions flag.
	OVR_MST         CSR_Bits = 0x01 << 4  //  ADC multimode master group regular overrun flag.
	JEOC_MST        CSR_Bits = 0x01 << 5  //  ADC multimode master group injected end of unitary conversion flag.
	JEOS_MST        CSR_Bits = 0x01 << 6  //  ADC multimode master group injected end of sequence conversions flag.
	EOSMP_SLV       CSR_Bits = 0x01 << 17 //  ADC multimode slave group regular end of sampling flag.
	EOC_SLV         CSR_Bits = 0x01 << 18 //  ADC multimode slave group regular end of unitary conversion flag.
	EOS_SLV         CSR_Bits = 0x01 << 19 //  ADC multimode slave group regular end of sequence conversions flag.
	OVR_SLV         CSR_Bits = 0x01 << 20 //  ADC multimode slave group regular overrun flag.
	JEOC_SLV        CSR_Bits = 0x01 << 21 //  ADC multimode slave group injected end of unitary conversion flag.
	JEOS_SLV        CSR_Bits = 0x01 << 22 //  ADC multimode slave group injected end of sequence conversions flag.
)

const (
	ADRDY_MSTn       = 0
	ADRDY_EOSMP_MSTn = 1
	ADRDY_EOC_MSTn   = 2
	ADRDY_EOS_MSTn   = 3
	ADRDY_OVR_MSTn   = 4
	ADRDY_JEOC_MSTn  = 5
	ADRDY_JEOS_MSTn  = 6
	AWD1_MSTn        = 7
	AWD2_MSTn        = 8
	AWD3_MSTn        = 9
	JQOVF_MSTn       = 10
	ADRDY_SLVn       = 16
	ADRDY_EOSMP_SLVn = 17
	ADRDY_EOC_SLVn   = 18
	ADRDY_EOS_SLVn   = 19
	ADRDY_OVR_SLVn   = 20
	ADRDY_JEOC_SLVn  = 21
	ADRDY_JEOS_SLVn  = 22
	AWD1_SLVn        = 23
	AWD2_SLVn        = 24
	AWD3_SLVn        = 25
	JQOVF_SLVn       = 26
)

const (
	MULTI   CCR_Bits = 0x1F << 0  //+ Multi ADC mode selection.
	DELAY   CCR_Bits = 0x0F << 8  //+ Delay between 2 sampling phases.
	MDMACFG CCR_Bits = 0x01 << 13 //+ DMA configuration for multi-ADC mode.
	MDMA    CCR_Bits = 0x03 << 14 //+ DMA mode for multi-ADC mode.
	CKMODE  CCR_Bits = 0x03 << 16 //+ ADC clock mode.
	VREFEN  CCR_Bits = 0x01 << 22 //+ VREFINT enable.
	TSEN    CCR_Bits = 0x01 << 23 //+ Temperature sensor enable.
	VBATEN  CCR_Bits = 0x01 << 24 //+ VBAT enable.
	DUAL    CCR_Bits = 0x1F << 0  //  ADC multimode mode selection.
)

const (
	MULTIn   = 0
	DELAYn   = 8
	MDMACFGn = 13
	MDMAn    = 14
	CKMODEn  = 16
	VREFENn  = 22
	TSENn    = 23
	VBATENn  = 24
)

const (
	RDATA_MST CDR_Bits = 0xFFFF << 0  //+ Regular Data of the master ADC.
	RDATA_SLV CDR_Bits = 0xFFFF << 16 //+ Regular Data of the master ADC.
)

const (
	RDATA_MSTn = 0
	RDATA_SLVn = 16
)
