// +build f303xe

// Peripheral: USART_Periph  Universal Synchronous Asynchronous Receiver Transmitter.
// Instances:
//  USART2  mmap.USART2_BASE
//  USART3  mmap.USART3_BASE
//  UART4   mmap.UART4_BASE
//  UART5   mmap.UART5_BASE
//  USART1  mmap.USART1_BASE
// Registers:
//  0x00 32  CR1  Control register 1.
//  0x04 32  CR2  Control register 2.
//  0x08 32  CR3  Control register 3.
//  0x0C 32  BRR  Baud rate register.
//  0x10 32  GTPR Guard time and prescaler register.
//  0x14 32  RTOR Receiver Time Out register.
//  0x18 32  RQR  Request register.
//  0x1C 32  ISR  Interrupt and status register.
//  0x20 32  ICR  Interrupt flag Clear register.
//  0x24 16  RDR  Receive Data register.
//  0x28 16  TDR  Transmit Data register.
// Import:
//  stm32/o/f303xe/mmap
package usart

// DO NOT EDIT THIS FILE. GENERATED BY stm32xgen.

const (
	UE     CR1_Bits = 0x01 << 0     //+ USART Enable.
	UESM   CR1_Bits = 0x01 << 1     //+ USART Enable in STOP Mode.
	RE     CR1_Bits = 0x01 << 2     //+ Receiver Enable.
	TE     CR1_Bits = 0x01 << 3     //+ Transmitter Enable.
	IDLEIE CR1_Bits = 0x01 << 4     //+ IDLE Interrupt Enable.
	RXNEIE CR1_Bits = 0x01 << 5     //+ RXNE Interrupt Enable.
	TCIE   CR1_Bits = 0x01 << 6     //+ Transmission Complete Interrupt Enable.
	TXEIE  CR1_Bits = 0x01 << 7     //+ TXE Interrupt Enable.
	PEIE   CR1_Bits = 0x01 << 8     //+ PE Interrupt Enable.
	PS     CR1_Bits = 0x01 << 9     //+ Parity Selection.
	PCE    CR1_Bits = 0x01 << 10    //+ Parity Control Enable.
	WAKE   CR1_Bits = 0x01 << 11    //+ Receiver Wakeup method.
	M0     CR1_Bits = 0x01 << 12    //+ Word length bit 0.
	MME    CR1_Bits = 0x01 << 13    //+ Mute Mode Enable.
	CMIE   CR1_Bits = 0x01 << 14    //+ Character match interrupt enable.
	OVER8  CR1_Bits = 0x01 << 15    //+ Oversampling by 8-bit or 16-bit mode.
	DEDT   CR1_Bits = 0x1F << 16    //+ DEDT[4:0] bits (Driver Enable Deassertion Time).
	DEAT   CR1_Bits = 0x1F << 21    //+ DEAT[4:0] bits (Driver Enable Assertion Time).
	RTOIE  CR1_Bits = 0x01 << 26    //+ Receive Time Out interrupt enable.
	EOBIE  CR1_Bits = 0x01 << 27    //+ End of Block interrupt enable.
	M1     CR1_Bits = 0x01 << 28    //+ Word length bit 1.
	M      CR1_Bits = 0x10001 << 12 //  [M1:M0] Word length.
)

const (
	UEn     = 0
	UESMn   = 1
	REn     = 2
	TEn     = 3
	IDLEIEn = 4
	RXNEIEn = 5
	TCIEn   = 6
	TXEIEn  = 7
	PEIEn   = 8
	PSn     = 9
	PCEn    = 10
	WAKEn   = 11
	M0n     = 12
	MMEn    = 13
	CMIEn   = 14
	OVER8n  = 15
	DEDTn   = 16
	DEATn   = 21
	RTOIEn  = 26
	EOBIEn  = 27
	M1n     = 28
)

const (
	ADDM7    CR2_Bits = 0x01 << 4  //+ 7-bit or 4-bit Address Detection.
	LBDL     CR2_Bits = 0x01 << 5  //+ LIN Break Detection Length.
	LBDIE    CR2_Bits = 0x01 << 6  //+ LIN Break Detection Interrupt Enable.
	LBCL     CR2_Bits = 0x01 << 8  //+ Last Bit Clock pulse.
	CPHA     CR2_Bits = 0x01 << 9  //+ Clock Phase.
	CPOL     CR2_Bits = 0x01 << 10 //+ Clock Polarity.
	CLKEN    CR2_Bits = 0x01 << 11 //+ Clock Enable.
	STOP     CR2_Bits = 0x03 << 12 //+ STOP[1:0] bits (STOP bits).
	LINEN    CR2_Bits = 0x01 << 14 //+ LIN mode enable.
	SWAP     CR2_Bits = 0x01 << 15 //+ SWAP TX/RX pins.
	RXINV    CR2_Bits = 0x01 << 16 //+ RX pin active level inversion.
	TXINV    CR2_Bits = 0x01 << 17 //+ TX pin active level inversion.
	DATAINV  CR2_Bits = 0x01 << 18 //+ Binary data inversion.
	MSBFIRST CR2_Bits = 0x01 << 19 //+ Most Significant Bit First.
	ABREN    CR2_Bits = 0x01 << 20 //+ Auto Baud-Rate Enable.
	ABRMODE  CR2_Bits = 0x03 << 21 //+ ABRMOD[1:0] bits (Auto Baud-Rate Mode).
	RTOEN    CR2_Bits = 0x01 << 23 //+ Receiver Time-Out enable.
	ADD      CR2_Bits = 0xFF << 24 //+ Address of the USART node.
)

const (
	ADDM7n    = 4
	LBDLn     = 5
	LBDIEn    = 6
	LBCLn     = 8
	CPHAn     = 9
	CPOLn     = 10
	CLKENn    = 11
	STOPn     = 12
	LINENn    = 14
	SWAPn     = 15
	RXINVn    = 16
	TXINVn    = 17
	DATAINVn  = 18
	MSBFIRSTn = 19
	ABRENn    = 20
	ABRMODEn  = 21
	RTOENn    = 23
	ADDn      = 24
)

const (
	EIE     CR3_Bits = 0x01 << 0  //+ Error Interrupt Enable.
	IREN    CR3_Bits = 0x01 << 1  //+ IrDA mode Enable.
	IRLP    CR3_Bits = 0x01 << 2  //+ IrDA Low-Power.
	HDSEL   CR3_Bits = 0x01 << 3  //+ Half-Duplex Selection.
	NACK    CR3_Bits = 0x01 << 4  //+ SmartCard NACK enable.
	SCEN    CR3_Bits = 0x01 << 5  //+ SmartCard mode enable.
	DMAR    CR3_Bits = 0x01 << 6  //+ DMA Enable Receiver.
	DMAT    CR3_Bits = 0x01 << 7  //+ DMA Enable Transmitter.
	RTSE    CR3_Bits = 0x01 << 8  //+ RTS Enable.
	CTSE    CR3_Bits = 0x01 << 9  //+ CTS Enable.
	CTSIE   CR3_Bits = 0x01 << 10 //+ CTS Interrupt Enable.
	ONEBIT  CR3_Bits = 0x01 << 11 //+ One sample bit method enable.
	OVRDIS  CR3_Bits = 0x01 << 12 //+ Overrun Disable.
	DDRE    CR3_Bits = 0x01 << 13 //+ DMA Disable on Reception Error.
	DEM     CR3_Bits = 0x01 << 14 //+ Driver Enable Mode.
	DEP     CR3_Bits = 0x01 << 15 //+ Driver Enable Polarity Selection.
	SCARCNT CR3_Bits = 0x07 << 17 //+ SCARCNT[2:0] bits (SmartCard Auto-Retry Count).
	WUS     CR3_Bits = 0x03 << 20 //+ WUS[1:0] bits (Wake UP Interrupt Flag Selection).
	WUFIE   CR3_Bits = 0x01 << 22 //+ Wake Up Interrupt Enable.
)

const (
	EIEn     = 0
	IRENn    = 1
	IRLPn    = 2
	HDSELn   = 3
	NACKn    = 4
	SCENn    = 5
	DMARn    = 6
	DMATn    = 7
	RTSEn    = 8
	CTSEn    = 9
	CTSIEn   = 10
	ONEBITn  = 11
	OVRDISn  = 12
	DDREn    = 13
	DEMn     = 14
	DEPn     = 15
	SCARCNTn = 17
	WUSn     = 20
	WUFIEn   = 22
)

const (
	DIV_FRACTION BRR_Bits = 0x0F << 0  //+ Fraction of USARTDIV.
	DIV_MANTISSA BRR_Bits = 0xFFF << 4 //+ Mantissa of USARTDIV.
)

const (
	DIV_FRACTIONn = 0
	DIV_MANTISSAn = 4
)

const (
	PSC GTPR_Bits = 0xFF << 0 //+ PSC[7:0] bits (Prescaler value).
	GT  GTPR_Bits = 0xFF << 8 //+ GT[7:0] bits (Guard time value).
)

const (
	PSCn = 0
	GTn  = 8
)

const (
	RTO  RTOR_Bits = 0xFFFFFF << 0 //+ Receiver Time Out Value.
	BLEN RTOR_Bits = 0xFF << 24    //+ Block Length.
)

const (
	RTOn  = 0
	BLENn = 24
)

const (
	ABRRQ RQR_Bits = 0x01 << 0 //+ Auto-Baud Rate Request.
	SBKRQ RQR_Bits = 0x01 << 1 //+ Send Break Request.
	MMRQ  RQR_Bits = 0x01 << 2 //+ Mute Mode Request.
	RXFRQ RQR_Bits = 0x01 << 3 //+ Receive Data flush Request.
	TXFRQ RQR_Bits = 0x01 << 4 //+ Transmit data flush Request.
)

const (
	ABRRQn = 0
	SBKRQn = 1
	MMRQn  = 2
	RXFRQn = 3
	TXFRQn = 4
)

const (
	PE    ISR_Bits = 0x01 << 0  //+ Parity Error.
	FE    ISR_Bits = 0x01 << 1  //+ Framing Error.
	NE    ISR_Bits = 0x01 << 2  //+ Noise detected Flag.
	ORE   ISR_Bits = 0x01 << 3  //+ OverRun Error.
	IDLE  ISR_Bits = 0x01 << 4  //+ IDLE line detected.
	RXNE  ISR_Bits = 0x01 << 5  //+ Read Data Register Not Empty.
	TC    ISR_Bits = 0x01 << 6  //+ Transmission Complete.
	TXE   ISR_Bits = 0x01 << 7  //+ Transmit Data Register Empty.
	LBDF  ISR_Bits = 0x01 << 8  //+ LIN Break Detection Flag.
	CTSIF ISR_Bits = 0x01 << 9  //+ CTS interrupt flag.
	CTS   ISR_Bits = 0x01 << 10 //+ CTS flag.
	RTOF  ISR_Bits = 0x01 << 11 //+ Receiver Time Out.
	EOBF  ISR_Bits = 0x01 << 12 //+ End Of Block Flag.
	ABRE  ISR_Bits = 0x01 << 14 //+ Auto-Baud Rate Error.
	ABRF  ISR_Bits = 0x01 << 15 //+ Auto-Baud Rate Flag.
	BUSY  ISR_Bits = 0x01 << 16 //+ Busy Flag.
	CMF   ISR_Bits = 0x01 << 17 //+ Character Match Flag.
	SBKF  ISR_Bits = 0x01 << 18 //+ Send Break Flag.
	RWU   ISR_Bits = 0x01 << 19 //+ Receive Wake Up from mute mode Flag.
	WUF   ISR_Bits = 0x01 << 20 //+ Wake Up from stop mode Flag.
	TEACK ISR_Bits = 0x01 << 21 //+ Transmit Enable Acknowledge Flag.
	REACK ISR_Bits = 0x01 << 22 //+ Receive Enable Acknowledge Flag.
)

const (
	PEn    = 0
	FEn    = 1
	NEn    = 2
	OREn   = 3
	IDLEn  = 4
	RXNEn  = 5
	TCn    = 6
	TXEn   = 7
	LBDFn  = 8
	CTSIFn = 9
	CTSn   = 10
	RTOFn  = 11
	EOBFn  = 12
	ABREn  = 14
	ABRFn  = 15
	BUSYn  = 16
	CMFn   = 17
	SBKFn  = 18
	RWUn   = 19
	WUFn   = 20
	TEACKn = 21
	REACKn = 22
)

const (
	PECF   ICR_Bits = 0x01 << 0  //+ Parity Error Clear Flag.
	FECF   ICR_Bits = 0x01 << 1  //+ Framing Error Clear Flag.
	NCF    ICR_Bits = 0x01 << 2  //+ Noise detected Clear Flag.
	ORECF  ICR_Bits = 0x01 << 3  //+ OverRun Error Clear Flag.
	IDLECF ICR_Bits = 0x01 << 4  //+ IDLE line detected Clear Flag.
	TCCF   ICR_Bits = 0x01 << 6  //+ Transmission Complete Clear Flag.
	LBDCF  ICR_Bits = 0x01 << 8  //+ LIN Break Detection Clear Flag.
	CTSCF  ICR_Bits = 0x01 << 9  //+ CTS Interrupt Clear Flag.
	RTOCF  ICR_Bits = 0x01 << 11 //+ Receiver Time Out Clear Flag.
	EOBCF  ICR_Bits = 0x01 << 12 //+ End Of Block Clear Flag.
	CMCF   ICR_Bits = 0x01 << 17 //+ Character Match Clear Flag.
	WUCF   ICR_Bits = 0x01 << 20 //+ Wake Up from stop mode Clear Flag.
)

const (
	PECFn   = 0
	FECFn   = 1
	NCFn    = 2
	ORECFn  = 3
	IDLECFn = 4
	TCCFn   = 6
	LBDCFn  = 8
	CTSCFn  = 9
	RTOCFn  = 11
	EOBCFn  = 12
	CMCFn   = 17
	WUCFn   = 20
)
