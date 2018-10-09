"".fibonacci STEXT size=211 args=0x8 locals=0x58
	0x0000 00000 (test_channel.go:3)	TEXT	"".fibonacci(SB), $88-8
	0x0000 00000 (test_channel.go:3)	MOVL	TLS, CX
	0x0007 00007 (test_channel.go:3)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test_channel.go:3)	CMPL	SP, 8(CX)
	0x0010 00016 (test_channel.go:3)	JLS	201
	0x0016 00022 (test_channel.go:3)	SUBL	$88, SP
	0x0019 00025 (test_channel.go:3)	FUNCDATA	$0, gclocals·5e326a2f5498e528ce8fbe7bd86e7d21(SB)
	0x0019 00025 (test_channel.go:3)	FUNCDATA	$1, gclocals·330e121c4652a955d2e0ff6e9125912d(SB)
	0x0019 00025 (test_channel.go:3)	FUNCDATA	$3, gclocals·e0b5c08eb812f276885e532abca71746(SB)
	0x0019 00025 (test_channel.go:3)	PCDATA	$2, $0
	0x0019 00025 (test_channel.go:3)	PCDATA	$0, $0
	0x0019 00025 (test_channel.go:3)	XORL	AX, AX
	0x001b 00027 (test_channel.go:3)	MOVL	$1, CX
	0x0020 00032 (test_channel.go:5)	JMP	45
	0x0022 00034 (test_channel.go:8)	MOVL	"".y+20(SP), AX
	0x0026 00038 (test_channel.go:8)	MOVL	"".x+24(SP), DX
	0x002a 00042 (test_channel.go:8)	LEAL	(DX)(AX*1), CX
	0x002d 00045 (test_channel.go:8)	MOVL	CX, "".y+20(SP)
	0x0031 00049 (test_channel.go:7)	MOVL	AX, "".x+24(SP)
	0x0035 00053 (test_channel.go:7)	MOVL	AX, ""..autotmp_5+28(SP)
	0x0039 00057 (test_channel.go:6)	PCDATA	$2, $1
	0x0039 00057 (test_channel.go:6)	PCDATA	$0, $1
	0x0039 00057 (test_channel.go:6)	LEAL	""..autotmp_6+40(SP), DI
	0x003d 00061 (test_channel.go:6)	XORL	AX, AX
	0x003f 00063 (test_channel.go:6)	PCDATA	$2, $0
	0x003f 00063 (test_channel.go:6)	DUFFZERO	$116
	0x0044 00068 (test_channel.go:6)	MOVL	$0, ""..autotmp_7+32(SP)
	0x004c 00076 (test_channel.go:6)	MOVL	$0, ""..autotmp_7+36(SP)
	0x0054 00084 (test_channel.go:7)	MOVW	$2, ""..autotmp_6+48(SP)
	0x005b 00091 (test_channel.go:7)	PCDATA	$2, $2
	0x005b 00091 (test_channel.go:7)	MOVL	"".c+92(SP), CX
	0x005f 00095 (test_channel.go:7)	PCDATA	$2, $0
	0x005f 00095 (test_channel.go:7)	MOVL	CX, ""..autotmp_6+40(SP)
	0x0063 00099 (test_channel.go:7)	PCDATA	$2, $3
	0x0063 00099 (test_channel.go:7)	LEAL	""..autotmp_5+28(SP), BX
	0x0067 00103 (test_channel.go:7)	PCDATA	$2, $0
	0x0067 00103 (test_channel.go:7)	MOVL	BX, ""..autotmp_6+44(SP)
	0x006b 00107 (test_channel.go:9)	MOVW	$1, ""..autotmp_6+72(SP)
	0x0072 00114 (test_channel.go:9)	PCDATA	$2, $3
	0x0072 00114 (test_channel.go:9)	MOVL	"".quit+96(SP), BX
	0x0076 00118 (test_channel.go:9)	PCDATA	$2, $0
	0x0076 00118 (test_channel.go:9)	MOVL	BX, ""..autotmp_6+64(SP)
	0x007a 00122 (test_channel.go:6)	PCDATA	$2, $4
	0x007a 00122 (test_channel.go:6)	LEAL	""..autotmp_7+32(SP), BP
	0x007e 00126 (test_channel.go:6)	PCDATA	$2, $0
	0x007e 00126 (test_channel.go:6)	MOVL	BP, 4(SP)
	0x0082 00130 (test_channel.go:6)	PCDATA	$2, $4
	0x0082 00130 (test_channel.go:6)	LEAL	""..autotmp_6+40(SP), BP
	0x0086 00134 (test_channel.go:6)	PCDATA	$2, $0
	0x0086 00134 (test_channel.go:6)	MOVL	BP, (SP)
	0x0089 00137 (test_channel.go:6)	MOVL	$2, 8(SP)
	0x0091 00145 (test_channel.go:6)	CALL	runtime.selectgo(SB)
	0x0096 00150 (test_channel.go:6)	MOVL	12(SP), CX
	0x009a 00154 (test_channel.go:7)	PCDATA	$0, $0
	0x009a 00154 (test_channel.go:7)	TESTL	CX, CX
	0x009c 00156 (test_channel.go:7)	JEQ	34
	0x009e 00158 (test_channel.go:9)	PCDATA	$0, $2
	0x009e 00158 (test_channel.go:9)	CMPL	CX, $1
	0x00a1 00161 (test_channel.go:9)	JNE	199
	0x00a3 00163 (test_channel.go:10)	CALL	runtime.printlock(SB)
	0x00a8 00168 (test_channel.go:10)	PCDATA	$2, $5
	0x00a8 00168 (test_channel.go:10)	LEAL	go.string."quit\n"(SB), AX
	0x00ae 00174 (test_channel.go:10)	PCDATA	$2, $0
	0x00ae 00174 (test_channel.go:10)	MOVL	AX, (SP)
	0x00b1 00177 (test_channel.go:10)	MOVL	$5, 4(SP)
	0x00b9 00185 (test_channel.go:10)	CALL	runtime.printstring(SB)
	0x00be 00190 (test_channel.go:10)	CALL	runtime.printunlock(SB)
	0x00c3 00195 (test_channel.go:11)	ADDL	$88, SP
	0x00c6 00198 (test_channel.go:11)	RET
	0x00c7 00199 (<unknown line number>)	PCDATA	$2, $-2
	0x00c7 00199 (<unknown line number>)	PCDATA	$0, $-2
	0x00c7 00199 (<unknown line number>)	UNDEF
	0x00c9 00201 (<unknown line number>)	NOP
	0x00c9 00201 (test_channel.go:3)	PCDATA	$0, $-1
	0x00c9 00201 (test_channel.go:3)	PCDATA	$2, $-1
	0x00c9 00201 (test_channel.go:3)	CALL	runtime.morestack_noctxt(SB)
	0x00ce 00206 (test_channel.go:3)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 0f 86 b3 00 00 00 83 ec 58 31 c0 b9 01 00 00 00  ........X1......
	0x0020 eb 0b 8b 44 24 14 8b 54 24 18 8d 0c 02 89 4c 24  ...D$..T$.....L$
	0x0030 14 89 44 24 18 89 44 24 1c 8d 7c 24 28 31 c0 e8  ..D$..D$..|$(1..
	0x0040 00 00 00 00 c7 44 24 20 00 00 00 00 c7 44 24 24  .....D$ .....D$$
	0x0050 00 00 00 00 66 c7 44 24 30 02 00 8b 4c 24 5c 89  ....f.D$0...L$\.
	0x0060 4c 24 28 8d 5c 24 1c 89 5c 24 2c 66 c7 44 24 48  L$(.\$..\$,f.D$H
	0x0070 01 00 8b 5c 24 60 89 5c 24 40 8d 6c 24 20 89 6c  ...\$`.\$@.l$ .l
	0x0080 24 04 8d 6c 24 28 89 2c 24 c7 44 24 08 02 00 00  $..l$(.,$.D$....
	0x0090 00 e8 00 00 00 00 8b 4c 24 0c 85 c9 74 84 83 f9  .......L$...t...
	0x00a0 01 75 24 e8 00 00 00 00 8d 05 00 00 00 00 89 04  .u$.............
	0x00b0 24 c7 44 24 04 05 00 00 00 e8 00 00 00 00 e8 00  $.D$............
	0x00c0 00 00 00 83 c4 58 c3 0f 0b e8 00 00 00 00 e9 2d  .....X.........-
	0x00d0 ff ff ff                                         ...
	rel 9+4 t=16 TLS+0
	rel 64+4 t=8 runtime.duffzero+116
	rel 146+4 t=8 runtime.selectgo+0
	rel 164+4 t=8 runtime.printlock+0
	rel 170+4 t=1 go.string."quit\n"+0
	rel 186+4 t=8 runtime.printstring+0
	rel 191+4 t=8 runtime.printunlock+0
	rel 202+4 t=8 runtime.morestack_noctxt+0
"".main1 STEXT size=149 args=0x0 locals=0x18
	0x0000 00000 (test_channel.go:15)	TEXT	"".main1(SB), $24-0
	0x0000 00000 (test_channel.go:15)	MOVL	TLS, CX
	0x0007 00007 (test_channel.go:15)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test_channel.go:15)	CMPL	SP, 8(CX)
	0x0010 00016 (test_channel.go:15)	JLS	139
	0x0012 00018 (test_channel.go:15)	SUBL	$24, SP
	0x0015 00021 (test_channel.go:15)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x0015 00021 (test_channel.go:15)	FUNCDATA	$1, gclocals·a57540db91d4df0a95a343aea8a16f48(SB)
	0x0015 00021 (test_channel.go:15)	FUNCDATA	$3, gclocals·f9e9e71b3d85c7d9ecf49fc49dd56c0d(SB)
	0x0015 00021 (test_channel.go:16)	PCDATA	$2, $1
	0x0015 00021 (test_channel.go:16)	PCDATA	$0, $0
	0x0015 00021 (test_channel.go:16)	LEAL	type.chan int(SB), AX
	0x001b 00027 (test_channel.go:16)	PCDATA	$2, $0
	0x001b 00027 (test_channel.go:16)	MOVL	AX, (SP)
	0x001e 00030 (test_channel.go:16)	MOVL	$0, 4(SP)
	0x0026 00038 (test_channel.go:16)	CALL	runtime.makechan(SB)
	0x002b 00043 (test_channel.go:16)	PCDATA	$2, $1
	0x002b 00043 (test_channel.go:16)	MOVL	8(SP), AX
	0x002f 00047 (test_channel.go:16)	PCDATA	$2, $0
	0x002f 00047 (test_channel.go:16)	PCDATA	$0, $1
	0x002f 00047 (test_channel.go:16)	MOVL	AX, "".c+20(SP)
	0x0033 00051 (test_channel.go:17)	PCDATA	$2, $2
	0x0033 00051 (test_channel.go:17)	LEAL	type.chan int(SB), CX
	0x0039 00057 (test_channel.go:17)	PCDATA	$2, $0
	0x0039 00057 (test_channel.go:17)	MOVL	CX, (SP)
	0x003c 00060 (test_channel.go:17)	MOVL	$0, 4(SP)
	0x0044 00068 (test_channel.go:17)	CALL	runtime.makechan(SB)
	0x0049 00073 (test_channel.go:17)	PCDATA	$2, $1
	0x0049 00073 (test_channel.go:17)	MOVL	8(SP), AX
	0x004d 00077 (test_channel.go:17)	PCDATA	$0, $2
	0x004d 00077 (test_channel.go:17)	MOVL	AX, "".quit+16(SP)
	0x0051 00081 (test_channel.go:23)	PCDATA	$2, $3
	0x0051 00081 (test_channel.go:23)	MOVL	"".c+20(SP), CX
	0x0055 00085 (test_channel.go:23)	PCDATA	$2, $1
	0x0055 00085 (test_channel.go:23)	MOVL	CX, 8(SP)
	0x0059 00089 (test_channel.go:23)	PCDATA	$2, $0
	0x0059 00089 (test_channel.go:23)	MOVL	AX, 12(SP)
	0x005d 00093 (test_channel.go:18)	MOVL	$8, (SP)
	0x0064 00100 (test_channel.go:18)	PCDATA	$2, $4
	0x0064 00100 (test_channel.go:18)	LEAL	"".main1.func1·f(SB), DX
	0x006a 00106 (test_channel.go:18)	PCDATA	$2, $0
	0x006a 00106 (test_channel.go:18)	MOVL	DX, 4(SP)
	0x006e 00110 (test_channel.go:18)	CALL	runtime.newproc(SB)
	0x0073 00115 (test_channel.go:24)	PCDATA	$2, $1
	0x0073 00115 (test_channel.go:24)	PCDATA	$0, $3
	0x0073 00115 (test_channel.go:24)	MOVL	"".c+20(SP), AX
	0x0077 00119 (test_channel.go:24)	PCDATA	$2, $0
	0x0077 00119 (test_channel.go:24)	MOVL	AX, (SP)
	0x007a 00122 (test_channel.go:24)	PCDATA	$2, $1
	0x007a 00122 (test_channel.go:24)	PCDATA	$0, $0
	0x007a 00122 (test_channel.go:24)	MOVL	"".quit+16(SP), AX
	0x007e 00126 (test_channel.go:24)	PCDATA	$2, $0
	0x007e 00126 (test_channel.go:24)	MOVL	AX, 4(SP)
	0x0082 00130 (test_channel.go:24)	CALL	"".fibonacci(SB)
	0x0087 00135 (test_channel.go:25)	ADDL	$24, SP
	0x008a 00138 (test_channel.go:25)	RET
	0x008b 00139 (test_channel.go:25)	NOP
	0x008b 00139 (test_channel.go:15)	PCDATA	$0, $-1
	0x008b 00139 (test_channel.go:15)	PCDATA	$2, $-1
	0x008b 00139 (test_channel.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x0090 00144 (test_channel.go:15)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 79 83 ec 18 8d 05 00 00 00 00 89 04 24 c7 44  vy...........$.D
	0x0020 24 04 00 00 00 00 e8 00 00 00 00 8b 44 24 08 89  $...........D$..
	0x0030 44 24 14 8d 0d 00 00 00 00 89 0c 24 c7 44 24 04  D$.........$.D$.
	0x0040 00 00 00 00 e8 00 00 00 00 8b 44 24 08 89 44 24  ..........D$..D$
	0x0050 10 8b 4c 24 14 89 4c 24 08 89 44 24 0c c7 04 24  ..L$..L$..D$...$
	0x0060 08 00 00 00 8d 15 00 00 00 00 89 54 24 04 e8 00  ...........T$...
	0x0070 00 00 00 8b 44 24 14 89 04 24 8b 44 24 10 89 44  ....D$...$.D$..D
	0x0080 24 04 e8 00 00 00 00 83 c4 18 c3 e8 00 00 00 00  $...............
	0x0090 e9 6b ff ff ff                                   .k...
	rel 9+4 t=16 TLS+0
	rel 23+4 t=1 type.chan int+0
	rel 39+4 t=8 runtime.makechan+0
	rel 53+4 t=1 type.chan int+0
	rel 69+4 t=8 runtime.makechan+0
	rel 102+4 t=1 "".main1.func1·f+0
	rel 111+4 t=8 runtime.newproc+0
	rel 131+4 t=8 "".fibonacci+0
	rel 140+4 t=8 runtime.morestack_noctxt+0
"".main STEXT size=192 args=0x0 locals=0x20
	0x0000 00000 (test_channel.go:27)	TEXT	"".main(SB), $32-0
	0x0000 00000 (test_channel.go:27)	MOVL	TLS, CX
	0x0007 00007 (test_channel.go:27)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test_channel.go:27)	CMPL	SP, 8(CX)
	0x0010 00016 (test_channel.go:27)	JLS	182
	0x0016 00022 (test_channel.go:27)	SUBL	$32, SP
	0x0019 00025 (test_channel.go:27)	FUNCDATA	$0, gclocals·f6bd6b3389b872033d462029172c8612(SB)
	0x0019 00025 (test_channel.go:27)	FUNCDATA	$1, gclocals·20fcec35d1696a32d1314d1646792dfe(SB)
	0x0019 00025 (test_channel.go:27)	FUNCDATA	$3, gclocals·f9e9e71b3d85c7d9ecf49fc49dd56c0d(SB)
	0x0019 00025 (test_channel.go:28)	PCDATA	$2, $1
	0x0019 00025 (test_channel.go:28)	PCDATA	$0, $0
	0x0019 00025 (test_channel.go:28)	LEAL	type.string(SB), AX
	0x001f 00031 (test_channel.go:28)	PCDATA	$2, $0
	0x001f 00031 (test_channel.go:28)	MOVL	AX, (SP)
	0x0022 00034 (test_channel.go:28)	CALL	runtime.newobject(SB)
	0x0027 00039 (test_channel.go:28)	PCDATA	$2, $1
	0x0027 00039 (test_channel.go:28)	MOVL	4(SP), AX
	0x002b 00043 (test_channel.go:28)	PCDATA	$2, $0
	0x002b 00043 (test_channel.go:28)	PCDATA	$0, $1
	0x002b 00043 (test_channel.go:28)	MOVL	AX, "".&a+28(SP)
	0x002f 00047 (test_channel.go:29)	PCDATA	$2, $2
	0x002f 00047 (test_channel.go:29)	LEAL	type.chan int(SB), CX
	0x0035 00053 (test_channel.go:29)	PCDATA	$2, $0
	0x0035 00053 (test_channel.go:29)	MOVL	CX, (SP)
	0x0038 00056 (test_channel.go:29)	MOVL	$1, 4(SP)
	0x0040 00064 (test_channel.go:29)	CALL	runtime.makechan(SB)
	0x0045 00069 (test_channel.go:29)	PCDATA	$2, $1
	0x0045 00069 (test_channel.go:29)	MOVL	8(SP), AX
	0x0049 00073 (test_channel.go:29)	PCDATA	$0, $2
	0x0049 00073 (test_channel.go:29)	MOVL	AX, "".c+20(SP)
	0x004d 00077 (test_channel.go:35)	PCDATA	$2, $3
	0x004d 00077 (test_channel.go:35)	MOVL	"".&a+28(SP), CX
	0x0051 00081 (test_channel.go:35)	PCDATA	$2, $1
	0x0051 00081 (test_channel.go:35)	MOVL	CX, 8(SP)
	0x0055 00085 (test_channel.go:35)	PCDATA	$2, $0
	0x0055 00085 (test_channel.go:35)	MOVL	AX, 12(SP)
	0x0059 00089 (test_channel.go:31)	MOVL	$8, (SP)
	0x0060 00096 (test_channel.go:31)	PCDATA	$2, $4
	0x0060 00096 (test_channel.go:31)	LEAL	"".main.func1·f(SB), DX
	0x0066 00102 (test_channel.go:31)	PCDATA	$2, $0
	0x0066 00102 (test_channel.go:31)	MOVL	DX, 4(SP)
	0x006a 00106 (test_channel.go:31)	CALL	runtime.newproc(SB)
	0x006f 00111 (test_channel.go:36)	PCDATA	$2, $1
	0x006f 00111 (test_channel.go:36)	PCDATA	$0, $1
	0x006f 00111 (test_channel.go:36)	MOVL	"".c+20(SP), AX
	0x0073 00115 (test_channel.go:36)	PCDATA	$2, $0
	0x0073 00115 (test_channel.go:36)	MOVL	AX, (SP)
	0x0076 00118 (test_channel.go:36)	MOVL	$0, 4(SP)
	0x007e 00126 (test_channel.go:36)	CALL	runtime.chanrecv1(SB)
	0x0083 00131 (test_channel.go:38)	PCDATA	$2, $1
	0x0083 00131 (test_channel.go:38)	PCDATA	$0, $0
	0x0083 00131 (test_channel.go:38)	MOVL	"".&a+28(SP), AX
	0x0087 00135 (test_channel.go:38)	PCDATA	$2, $3
	0x0087 00135 (test_channel.go:38)	MOVL	(AX), CX
	0x0089 00137 (test_channel.go:38)	PCDATA	$2, $1
	0x0089 00137 (test_channel.go:38)	PCDATA	$0, $3
	0x0089 00137 (test_channel.go:38)	MOVL	CX, ""..autotmp_5+24(SP)
	0x008d 00141 (test_channel.go:38)	PCDATA	$2, $0
	0x008d 00141 (test_channel.go:38)	MOVL	4(AX), AX
	0x0090 00144 (test_channel.go:38)	MOVL	AX, ""..autotmp_6+16(SP)
	0x0094 00148 (test_channel.go:38)	CALL	runtime.printlock(SB)
	0x0099 00153 (test_channel.go:38)	PCDATA	$2, $1
	0x0099 00153 (test_channel.go:38)	PCDATA	$0, $0
	0x0099 00153 (test_channel.go:38)	MOVL	""..autotmp_5+24(SP), AX
	0x009d 00157 (test_channel.go:38)	PCDATA	$2, $0
	0x009d 00157 (test_channel.go:38)	MOVL	AX, (SP)
	0x00a0 00160 (test_channel.go:38)	MOVL	""..autotmp_6+16(SP), AX
	0x00a4 00164 (test_channel.go:38)	MOVL	AX, 4(SP)
	0x00a8 00168 (test_channel.go:38)	CALL	runtime.printstring(SB)
	0x00ad 00173 (test_channel.go:38)	CALL	runtime.printunlock(SB)
	0x00b2 00178 (test_channel.go:39)	ADDL	$32, SP
	0x00b5 00181 (test_channel.go:39)	RET
	0x00b6 00182 (test_channel.go:39)	NOP
	0x00b6 00182 (test_channel.go:27)	PCDATA	$0, $-1
	0x00b6 00182 (test_channel.go:27)	PCDATA	$2, $-1
	0x00b6 00182 (test_channel.go:27)	CALL	runtime.morestack_noctxt(SB)
	0x00bb 00187 (test_channel.go:27)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 0f 86 a0 00 00 00 83 ec 20 8d 05 00 00 00 00 89  ........ .......
	0x0020 04 24 e8 00 00 00 00 8b 44 24 04 89 44 24 1c 8d  .$......D$..D$..
	0x0030 0d 00 00 00 00 89 0c 24 c7 44 24 04 01 00 00 00  .......$.D$.....
	0x0040 e8 00 00 00 00 8b 44 24 08 89 44 24 14 8b 4c 24  ......D$..D$..L$
	0x0050 1c 89 4c 24 08 89 44 24 0c c7 04 24 08 00 00 00  ..L$..D$...$....
	0x0060 8d 15 00 00 00 00 89 54 24 04 e8 00 00 00 00 8b  .......T$.......
	0x0070 44 24 14 89 04 24 c7 44 24 04 00 00 00 00 e8 00  D$...$.D$.......
	0x0080 00 00 00 8b 44 24 1c 8b 08 89 4c 24 18 8b 40 04  ....D$....L$..@.
	0x0090 89 44 24 10 e8 00 00 00 00 8b 44 24 18 89 04 24  .D$.......D$...$
	0x00a0 8b 44 24 10 89 44 24 04 e8 00 00 00 00 e8 00 00  .D$..D$.........
	0x00b0 00 00 83 c4 20 c3 e8 00 00 00 00 e9 40 ff ff ff  .... .......@...
	rel 9+4 t=16 TLS+0
	rel 27+4 t=1 type.string+0
	rel 35+4 t=8 runtime.newobject+0
	rel 49+4 t=1 type.chan int+0
	rel 65+4 t=8 runtime.makechan+0
	rel 98+4 t=1 "".main.func1·f+0
	rel 107+4 t=8 runtime.newproc+0
	rel 127+4 t=8 runtime.chanrecv1+0
	rel 149+4 t=8 runtime.printlock+0
	rel 169+4 t=8 runtime.printstring+0
	rel 174+4 t=8 runtime.printunlock+0
	rel 183+4 t=8 runtime.morestack_noctxt+0
"".main1.func1 STEXT size=137 args=0x8 locals=0x10
	0x0000 00000 (test_channel.go:18)	TEXT	"".main1.func1(SB), $16-8
	0x0000 00000 (test_channel.go:18)	MOVL	TLS, CX
	0x0007 00007 (test_channel.go:18)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test_channel.go:18)	CMPL	SP, 8(CX)
	0x0010 00016 (test_channel.go:18)	JLS	127
	0x0012 00018 (test_channel.go:18)	SUBL	$16, SP
	0x0015 00021 (test_channel.go:18)	FUNCDATA	$0, gclocals·dc9b0298814590ca3ffc3a889546fc8b(SB)
	0x0015 00021 (test_channel.go:18)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0015 00021 (test_channel.go:18)	FUNCDATA	$3, gclocals·1cf923758aae2e428391d1783fe59973(SB)
	0x0015 00021 (test_channel.go:18)	PCDATA	$2, $0
	0x0015 00021 (test_channel.go:18)	PCDATA	$0, $0
	0x0015 00021 (test_channel.go:18)	XORL	AX, AX
	0x0017 00023 (test_channel.go:19)	JMP	96
	0x0019 00025 (test_channel.go:19)	MOVL	AX, "".i+8(SP)
	0x001d 00029 (test_channel.go:20)	MOVL	$0, ""..autotmp_3+12(SP)
	0x0025 00037 (test_channel.go:20)	PCDATA	$2, $1
	0x0025 00037 (test_channel.go:20)	MOVL	"".c+20(SP), AX
	0x0029 00041 (test_channel.go:20)	PCDATA	$2, $0
	0x0029 00041 (test_channel.go:20)	MOVL	AX, (SP)
	0x002c 00044 (test_channel.go:20)	PCDATA	$2, $2
	0x002c 00044 (test_channel.go:20)	LEAL	""..autotmp_3+12(SP), CX
	0x0030 00048 (test_channel.go:20)	PCDATA	$2, $0
	0x0030 00048 (test_channel.go:20)	MOVL	CX, 4(SP)
	0x0034 00052 (test_channel.go:20)	CALL	runtime.chanrecv1(SB)
	0x0039 00057 (test_channel.go:20)	CALL	runtime.printlock(SB)
	0x003e 00062 (test_channel.go:20)	MOVL	""..autotmp_3+12(SP), AX
	0x0042 00066 (test_channel.go:20)	MOVL	AX, (SP)
	0x0045 00069 (test_channel.go:20)	SARL	$31, AX
	0x0048 00072 (test_channel.go:20)	MOVL	AX, 4(SP)
	0x004c 00076 (test_channel.go:20)	CALL	runtime.printint(SB)
	0x0051 00081 (test_channel.go:20)	CALL	runtime.printnl(SB)
	0x0056 00086 (test_channel.go:20)	CALL	runtime.printunlock(SB)
	0x005b 00091 (test_channel.go:19)	MOVL	"".i+8(SP), AX
	0x005f 00095 (test_channel.go:19)	INCL	AX
	0x0060 00096 (test_channel.go:19)	CMPL	AX, $10
	0x0063 00099 (test_channel.go:19)	JLT	25
	0x0065 00101 (test_channel.go:22)	PCDATA	$2, $1
	0x0065 00101 (test_channel.go:22)	PCDATA	$0, $1
	0x0065 00101 (test_channel.go:22)	MOVL	"".quit+24(SP), AX
	0x0069 00105 (test_channel.go:22)	PCDATA	$2, $0
	0x0069 00105 (test_channel.go:22)	MOVL	AX, (SP)
	0x006c 00108 (test_channel.go:22)	PCDATA	$2, $1
	0x006c 00108 (test_channel.go:22)	LEAL	"".statictmp_0(SB), AX
	0x0072 00114 (test_channel.go:22)	PCDATA	$2, $0
	0x0072 00114 (test_channel.go:22)	MOVL	AX, 4(SP)
	0x0076 00118 (test_channel.go:22)	CALL	runtime.chansend1(SB)
	0x007b 00123 (test_channel.go:23)	ADDL	$16, SP
	0x007e 00126 (test_channel.go:23)	RET
	0x007f 00127 (test_channel.go:23)	NOP
	0x007f 00127 (test_channel.go:18)	PCDATA	$0, $-1
	0x007f 00127 (test_channel.go:18)	PCDATA	$2, $-1
	0x007f 00127 (test_channel.go:18)	CALL	runtime.morestack_noctxt(SB)
	0x0084 00132 (test_channel.go:18)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 6d 83 ec 10 31 c0 eb 47 89 44 24 08 c7 44 24  vm...1..G.D$..D$
	0x0020 0c 00 00 00 00 8b 44 24 14 89 04 24 8d 4c 24 0c  ......D$...$.L$.
	0x0030 89 4c 24 04 e8 00 00 00 00 e8 00 00 00 00 8b 44  .L$............D
	0x0040 24 0c 89 04 24 c1 f8 1f 89 44 24 04 e8 00 00 00  $...$....D$.....
	0x0050 00 e8 00 00 00 00 e8 00 00 00 00 8b 44 24 08 40  ............D$.@
	0x0060 83 f8 0a 7c b4 8b 44 24 18 89 04 24 8d 05 00 00  ...|..D$...$....
	0x0070 00 00 89 44 24 04 e8 00 00 00 00 83 c4 10 c3 e8  ...D$...........
	0x0080 00 00 00 00 e9 77 ff ff ff                       .....w...
	rel 9+4 t=16 TLS+0
	rel 53+4 t=8 runtime.chanrecv1+0
	rel 58+4 t=8 runtime.printlock+0
	rel 77+4 t=8 runtime.printint+0
	rel 82+4 t=8 runtime.printnl+0
	rel 87+4 t=8 runtime.printunlock+0
	rel 110+4 t=1 "".statictmp_0+0
	rel 119+4 t=8 runtime.chansend1+0
	rel 128+4 t=8 runtime.morestack_noctxt+0
"".main.func1 STEXT size=96 args=0x8 locals=0x8
	0x0000 00000 (test_channel.go:31)	TEXT	"".main.func1(SB), $8-8
	0x0000 00000 (test_channel.go:31)	MOVL	TLS, CX
	0x0007 00007 (test_channel.go:31)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (test_channel.go:31)	CMPL	SP, 8(CX)
	0x0010 00016 (test_channel.go:31)	JLS	89
	0x0012 00018 (test_channel.go:31)	SUBL	$8, SP
	0x0015 00021 (test_channel.go:31)	FUNCDATA	$0, gclocals·fdbf1f5761f6d17e8ae3f0aaecb6a3c5(SB)
	0x0015 00021 (test_channel.go:31)	FUNCDATA	$1, gclocals·7d2d5fca80364273fb07d5820a76fef4(SB)
	0x0015 00021 (test_channel.go:31)	FUNCDATA	$3, gclocals·cd666f9a7f09fcd2aca7dadbf3522159(SB)
	0x0015 00021 (test_channel.go:32)	PCDATA	$2, $1
	0x0015 00021 (test_channel.go:32)	PCDATA	$0, $1
	0x0015 00021 (test_channel.go:32)	MOVL	"".&a+12(SP), DI
	0x0019 00025 (test_channel.go:32)	MOVL	$11, 4(DI)
	0x0020 00032 (test_channel.go:32)	PCDATA	$2, $-2
	0x0020 00032 (test_channel.go:32)	PCDATA	$0, $-2
	0x0020 00032 (test_channel.go:32)	MOVL	runtime.writeBarrier(SB), CX
	0x0026 00038 (test_channel.go:32)	TESTL	CX, CX
	0x0028 00040 (test_channel.go:32)	JNE	76
	0x002a 00042 (test_channel.go:32)	LEAL	go.string."hello world"(SB), AX
	0x0030 00048 (test_channel.go:32)	MOVL	AX, (DI)
	0x0032 00050 (test_channel.go:33)	PCDATA	$2, $2
	0x0032 00050 (test_channel.go:33)	PCDATA	$0, $2
	0x0032 00050 (test_channel.go:33)	MOVL	"".c+16(SP), AX
	0x0036 00054 (test_channel.go:33)	PCDATA	$2, $0
	0x0036 00054 (test_channel.go:33)	MOVL	AX, (SP)
	0x0039 00057 (test_channel.go:33)	PCDATA	$2, $2
	0x0039 00057 (test_channel.go:33)	LEAL	"".statictmp_1(SB), AX
	0x003f 00063 (test_channel.go:33)	PCDATA	$2, $0
	0x003f 00063 (test_channel.go:33)	MOVL	AX, 4(SP)
	0x0043 00067 (test_channel.go:33)	CALL	runtime.chansend1(SB)
	0x0048 00072 (test_channel.go:35)	ADDL	$8, SP
	0x004b 00075 (test_channel.go:35)	RET
	0x004c 00076 (test_channel.go:32)	PCDATA	$2, $-2
	0x004c 00076 (test_channel.go:32)	PCDATA	$0, $-2
	0x004c 00076 (test_channel.go:32)	LEAL	go.string."hello world"(SB), AX
	0x0052 00082 (test_channel.go:32)	CALL	runtime.gcWriteBarrier(SB)
	0x0057 00087 (test_channel.go:32)	JMP	50
	0x0059 00089 (test_channel.go:32)	NOP
	0x0059 00089 (test_channel.go:31)	PCDATA	$0, $-1
	0x0059 00089 (test_channel.go:31)	PCDATA	$2, $-1
	0x0059 00089 (test_channel.go:31)	CALL	runtime.morestack_noctxt(SB)
	0x005e 00094 (test_channel.go:31)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 47 83 ec 08 8b 7c 24 0c c7 47 04 0b 00 00 00  vG....|$..G.....
	0x0020 8b 0d 00 00 00 00 85 c9 75 22 8d 05 00 00 00 00  ........u"......
	0x0030 89 07 8b 44 24 10 89 04 24 8d 05 00 00 00 00 89  ...D$...$.......
	0x0040 44 24 04 e8 00 00 00 00 83 c4 08 c3 8d 05 00 00  D$..............
	0x0050 00 00 e8 00 00 00 00 eb d9 e8 00 00 00 00 eb a0  ................
	rel 9+4 t=16 TLS+0
	rel 34+4 t=1 runtime.writeBarrier+0
	rel 44+4 t=1 go.string."hello world"+0
	rel 59+4 t=1 "".statictmp_1+0
	rel 68+4 t=8 runtime.chansend1+0
	rel 78+4 t=1 go.string."hello world"+0
	rel 83+4 t=8 runtime.gcWriteBarrier+0
	rel 90+4 t=8 runtime.morestack_noctxt+0
"".init STEXT size=55 args=0x0 locals=0x0
	0x0000 00000 (<autogenerated>:1)	TEXT	"".init(SB), $0-0
	0x0000 00000 (<autogenerated>:1)	MOVL	TLS, CX
	0x0007 00007 (<autogenerated>:1)	MOVL	(CX)(TLS*2), CX
	0x000d 00013 (<autogenerated>:1)	CMPL	SP, 8(CX)
	0x0010 00016 (<autogenerated>:1)	JLS	48
	0x0012 00018 (<autogenerated>:1)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (<autogenerated>:1)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (<autogenerated>:1)	FUNCDATA	$3, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0012 00018 (<autogenerated>:1)	PCDATA	$2, $0
	0x0012 00018 (<autogenerated>:1)	PCDATA	$0, $0
	0x0012 00018 (<autogenerated>:1)	MOVBLZX	"".initdone·(SB), AX
	0x0019 00025 (<autogenerated>:1)	CMPB	AX, $1
	0x001c 00028 (<autogenerated>:1)	JLS	31
	0x001e 00030 (<autogenerated>:1)	PCDATA	$2, $-2
	0x001e 00030 (<autogenerated>:1)	PCDATA	$0, $-2
	0x001e 00030 (<autogenerated>:1)	RET
	0x001f 00031 (<autogenerated>:1)	JNE	40
	0x0021 00033 (<autogenerated>:1)	PCDATA	$2, $0
	0x0021 00033 (<autogenerated>:1)	PCDATA	$0, $0
	0x0021 00033 (<autogenerated>:1)	CALL	runtime.throwinit(SB)
	0x0026 00038 (<autogenerated>:1)	UNDEF
	0x0028 00040 (<autogenerated>:1)	MOVB	$2, "".initdone·(SB)
	0x002f 00047 (<autogenerated>:1)	RET
	0x0030 00048 (<autogenerated>:1)	NOP
	0x0030 00048 (<autogenerated>:1)	PCDATA	$0, $-1
	0x0030 00048 (<autogenerated>:1)	PCDATA	$2, $-1
	0x0030 00048 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x0035 00053 (<autogenerated>:1)	JMP	0
	0x0000 65 8b 0d 00 00 00 00 8b 89 00 00 00 00 3b 61 08  e............;a.
	0x0010 76 1e 0f b6 05 00 00 00 00 80 f8 01 76 01 c3 75  v...........v..u
	0x0020 07 e8 00 00 00 00 0f 0b c6 05 00 00 00 00 02 c3  ................
	0x0030 e8 00 00 00 00 eb c9                             .......
	rel 9+4 t=16 TLS+0
	rel 21+4 t=1 "".initdone·+0
	rel 34+4 t=8 runtime.throwinit+0
	rel 42+4 t=1 "".initdone·+0
	rel 49+4 t=8 runtime.morestack_noctxt+0
go.string."quit" SRODATA dupok size=4
	0x0000 71 75 69 74                                      quit
go.string."quit\n" SRODATA dupok size=5
	0x0000 71 75 69 74 0a                                   quit.
go.loc."".fibonacci SDWARFLOC size=139
	0x0000 ff ff ff ff 00 00 00 00 1b 00 00 00 d3 00 00 00  ................
	0x0010 01 00 9c 00 00 00 00 00 00 00 00 ff ff ff ff 00  ................
	0x0020 00 00 00 1b 00 00 00 d3 00 00 00 02 00 91 04 00  ................
	0x0030 00 00 00 00 00 00 00 ff ff ff ff 00 00 00 00 26  ...............&
	0x0040 00 00 00 31 00 00 00 03 00 91 bc 7f 31 00 00 00  ...1........1...
	0x0050 35 00 00 00 03 00 91 b8 7f 35 00 00 00 d3 00 00  5........5......
	0x0060 00 03 00 91 bc 7f 00 00 00 00 00 00 00 00 ff ff  ................
	0x0070 ff ff 00 00 00 00 26 00 00 00 d3 00 00 00 03 00  ......&.........
	0x0080 91 b8 7f 00 00 00 00 00 00 00 00                 ...........
	rel 4+4 t=1 "".fibonacci+0
	rel 31+4 t=1 "".fibonacci+0
	rel 59+4 t=1 "".fibonacci+0
	rel 114+4 t=1 "".fibonacci+0
go.info."".fibonacci SDWARFINFO size=83
	0x0000 02 22 22 2e 66 69 62 6f 6e 61 63 63 69 00 00 00  ."".fibonacci...
	0x0010 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f 63 00  ..............c.
	0x0020 00 03 00 00 00 00 00 00 00 00 0f 71 75 69 74 00  ...........quit.
	0x0030 00 03 00 00 00 00 00 00 00 00 0a 78 00 04 00 00  ...........x....
	0x0040 00 00 00 00 00 00 0a 79 00 04 00 00 00 00 00 00  .......y........
	0x0050 00 00 00                                         ...
	rel 14+4 t=1 "".fibonacci+0
	rel 18+4 t=1 "".fibonacci+211
	rel 24+4 t=29 gofile../Users/majing/workspace/Go/Development/src/myGoPractices/test_channel.go+0
	rel 34+4 t=28 go.info.chan int+0
	rel 38+4 t=28 go.loc."".fibonacci+0
	rel 50+4 t=28 go.info.chan int+0
	rel 54+4 t=28 go.loc."".fibonacci+27
	rel 62+4 t=28 go.info.int+0
	rel 66+4 t=28 go.loc."".fibonacci+55
	rel 74+4 t=28 go.info.int+0
	rel 78+4 t=28 go.loc."".fibonacci+110
go.range."".fibonacci SDWARFRANGE size=0
go.isstmt."".fibonacci SDWARFMISC size=0
	0x0000 04 16 04 03 05 07 02 06 01 13 02 04 01 17 02 07  ................
	0x0010 01 10 02 07 01 08 02 04 01 1c 02 02 01 02 02 03  ................
	0x0020 01 02 02 05 01 1b 02 04 01 02 02 0a 00           .............
go.loc."".main1 SDWARFLOC size=78
	0x0000 ff ff ff ff 00 00 00 00 2f 00 00 00 49 00 00 00  ......../...I...
	0x0010 01 00 50 49 00 00 00 95 00 00 00 02 00 91 78 00  ..PI..........x.
	0x0020 00 00 00 00 00 00 00 ff ff ff ff 00 00 00 00 4d  ...............M
	0x0030 00 00 00 73 00 00 00 01 00 50 73 00 00 00 95 00  ...s.....Ps.....
	0x0040 00 00 02 00 91 74 00 00 00 00 00 00 00 00        .....t........
	rel 4+4 t=1 "".main1+0
	rel 43+4 t=1 "".main1+0
go.info."".main1 SDWARFINFO size=53
	0x0000 02 22 22 2e 6d 61 69 6e 31 00 00 00 00 00 00 00  ."".main1.......
	0x0010 00 00 01 9c 00 00 00 00 01 0a 63 00 10 00 00 00  ..........c.....
	0x0020 00 00 00 00 00 0a 71 75 69 74 00 11 00 00 00 00  ......quit......
	0x0030 00 00 00 00 00                                   .....
	rel 10+4 t=1 "".main1+0
	rel 14+4 t=1 "".main1+149
	rel 20+4 t=29 gofile../Users/majing/workspace/Go/Development/src/myGoPractices/test_channel.go+0
	rel 29+4 t=28 go.info.chan int+0
	rel 33+4 t=28 go.loc."".main1+0
	rel 44+4 t=28 go.info.chan int+0
	rel 48+4 t=28 go.loc."".main1+39
go.range."".main1 SDWARFRANGE size=0
go.isstmt."".main1 SDWARFMISC size=0
	0x0000 04 12 04 03 03 06 01 18 02 06 01 18 02 04 01 08  ................
	0x0010 02 07 01 0f 02 04 01 10 02 0e 00                 ...........
go.loc."".main SDWARFLOC size=78
	0x0000 ff ff ff ff 00 00 00 00 2b 00 00 00 45 00 00 00  ........+...E...
	0x0010 01 00 50 45 00 00 00 c0 00 00 00 02 00 91 78 00  ..PE..........x.
	0x0020 00 00 00 00 00 00 00 ff ff ff ff 00 00 00 00 49  ...............I
	0x0030 00 00 00 6f 00 00 00 01 00 50 6f 00 00 00 c0 00  ...o.....Po.....
	0x0040 00 00 02 00 91 70 00 00 00 00 00 00 00 00        .....p........
	rel 4+4 t=1 "".main+0
	rel 43+4 t=1 "".main+0
go.info."".main SDWARFINFO size=50
	0x0000 02 22 22 2e 6d 61 69 6e 00 00 00 00 00 00 00 00  ."".main........
	0x0010 00 01 9c 00 00 00 00 01 0a 26 61 00 1c 00 00 00  .........&a.....
	0x0020 00 00 00 00 00 0a 63 00 1d 00 00 00 00 00 00 00  ......c.........
	0x0030 00 00                                            ..
	rel 9+4 t=1 "".main+0
	rel 13+4 t=1 "".main+192
	rel 19+4 t=29 gofile../Users/majing/workspace/Go/Development/src/myGoPractices/test_channel.go+0
	rel 29+4 t=28 go.info.*string+0
	rel 33+4 t=28 go.loc."".main+0
	rel 41+4 t=28 go.info.chan int+0
	rel 45+4 t=28 go.loc."".main+39
go.range."".main SDWARFRANGE size=0
go.isstmt."".main SDWARFMISC size=0
	0x0000 04 16 04 03 03 06 01 10 02 06 01 18 02 04 01 08  ................
	0x0010 02 07 01 0f 02 04 01 10 02 04 01 2b 02 0e 00     ...........+...
go.loc."".main1.func1 SDWARFLOC size=94
	0x0000 ff ff ff ff 00 00 00 00 17 00 00 00 89 00 00 00  ................
	0x0010 01 00 9c 00 00 00 00 00 00 00 00 ff ff ff ff 00  ................
	0x0020 00 00 00 17 00 00 00 89 00 00 00 02 00 91 04 00  ................
	0x0030 00 00 00 00 00 00 00 ff ff ff ff 00 00 00 00 1d  ................
	0x0040 00 00 00 63 00 00 00 02 00 91 74 63 00 00 00 69  ...c......tc...i
	0x0050 00 00 00 01 00 50 00 00 00 00 00 00 00 00        .....P........
	rel 4+4 t=1 "".main1.func1+0
	rel 31+4 t=1 "".main1.func1+0
	rel 59+4 t=1 "".main1.func1+0
go.info."".main1.func1 SDWARFINFO size=83
	0x0000 02 22 22 2e 6d 61 69 6e 31 2e 66 75 6e 63 31 00  ."".main1.func1.
	0x0010 00 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0f  ................
	0x0020 63 00 00 14 00 00 00 00 00 00 00 00 0f 71 75 69  c............qui
	0x0030 74 00 00 16 00 00 00 00 00 00 00 00 14 00 00 00  t...............
	0x0040 00 00 00 00 00 0a 69 00 13 00 00 00 00 00 00 00  ......i.........
	0x0050 00 00 00                                         ...
	rel 16+4 t=1 "".main1.func1+0
	rel 20+4 t=1 "".main1.func1+137
	rel 26+4 t=29 gofile../Users/majing/workspace/Go/Development/src/myGoPractices/test_channel.go+0
	rel 36+4 t=28 go.info.chan int+0
	rel 40+4 t=28 go.loc."".main1.func1+0
	rel 52+4 t=28 go.info.chan int+0
	rel 56+4 t=28 go.loc."".main1.func1+27
	rel 61+4 t=1 "".main1.func1+23
	rel 65+4 t=1 "".main1.func1+101
	rel 73+4 t=28 go.info.int+0
	rel 77+4 t=28 go.loc."".main1.func1+55
go.range."".main1.func1 SDWARFRANGE size=0
go.isstmt."".main1.func1 SDWARFMISC size=0
	0x0000 04 12 04 03 05 02 02 02 01 04 02 08 01 36 02 04  .............6..
	0x0010 01 01 02 03 01 02 02 04 01 12 02 0e 00           .............
go.string."hello world" SRODATA dupok size=11
	0x0000 68 65 6c 6c 6f 20 77 6f 72 6c 64                 hello world
go.loc."".main.func1 SDWARFLOC size=0
go.info."".main.func1 SDWARFINFO size=52
	0x0000 02 22 22 2e 6d 61 69 6e 2e 66 75 6e 63 31 00 00  ."".main.func1..
	0x0010 00 00 00 00 00 00 00 01 9c 00 00 00 00 01 0e 26  ...............&
	0x0020 61 00 00 1f 00 00 00 00 00 0e 63 00 00 21 00 00  a.........c..!..
	0x0030 00 00 00 00                                      ....
	rel 15+4 t=1 "".main.func1+0
	rel 19+4 t=1 "".main.func1+96
	rel 25+4 t=29 gofile../Users/majing/workspace/Go/Development/src/myGoPractices/test_channel.go+0
	rel 36+4 t=28 go.info.*string+0
	rel 46+4 t=28 go.info.chan int+0
go.range."".main.func1 SDWARFRANGE size=0
go.isstmt."".main.func1 SDWARFMISC size=0
	0x0000 04 12 04 03 03 04 01 19 02 04 01 12 02 04 01 0d  ................
	0x0010 02 07 00                                         ...
go.loc."".init SDWARFLOC size=0
go.info."".init SDWARFINFO size=25
	0x0000 02 22 22 2e 69 6e 69 74 00 00 00 00 00 00 00 00  ."".init........
	0x0010 00 01 9c 00 00 00 00 01 00                       .........
	rel 9+4 t=1 "".init+0
	rel 13+4 t=1 "".init+55
	rel 19+4 t=29 gofile..<autogenerated>+0
go.range."".init SDWARFRANGE size=0
go.isstmt."".init SDWARFMISC size=0
	0x0000 04 19 01 05 02 01 01 07 02 09 01 01 02 07 00     ...............
"".statictmp_0 SRODATA size=4
"".statictmp_1 SRODATA size=4
"".initdone· SNOPTRBSS size=1
"".main.func1·f SRODATA dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".main.func1+0
"".main1.func1·f SRODATA dupok size=4
	0x0000 00 00 00 00                                      ....
	rel 0+4 t=1 "".main1.func1+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*[]uint16- SRODATA dupok size=12
	0x0000 00 00 09 2a 5b 5d 75 69 6e 74 31 36              ...*[]uint16
type.*[]uint16 SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 5f 18 a0 4c 00 04 04 36  ........_..L...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint16-+0
	rel 32+4 t=1 type.[]uint16+0
type.[]uint16 SRODATA dupok size=36
	0x0000 0c 00 00 00 04 00 00 00 e7 8e e3 20 02 04 04 17  ........... ....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]uint16-+0
	rel 28+4 t=6 type.*[]uint16+0
	rel 32+4 t=1 type.uint16+0
type..namedata.*[4]uint16- SRODATA dupok size=13
	0x0000 00 00 0a 2a 5b 34 5d 75 69 6e 74 31 36           ...*[4]uint16
type.*[4]uint16 SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 42 a7 e4 3c 00 04 04 36  ........B..<...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[4]uint16-+0
	rel 32+4 t=1 type.[4]uint16+0
runtime.gcbits. SRODATA dupok size=0
type.[4]uint16 SRODATA dupok size=44
	0x0000 08 00 00 00 00 00 00 00 3f 77 35 8e 02 02 02 91  ........?w5.....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 04 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+40
	rel 20+4 t=1 runtime.gcbits.+0
	rel 24+4 t=5 type..namedata.*[4]uint16-+0
	rel 28+4 t=6 type.*[4]uint16+0
	rel 32+4 t=1 type.uint16+0
	rel 36+4 t=1 type.[]uint16+0
type..namedata.*chan int- SRODATA dupok size=12
	0x0000 00 00 09 2a 63 68 61 6e 20 69 6e 74              ...*chan int
type.*chan int SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 ed 7b ed 3b 00 04 04 36  .........{.;...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*chan int-+0
	rel 32+4 t=1 type.chan int+0
type.chan int SRODATA dupok size=40
	0x0000 04 00 00 00 04 00 00 00 91 55 cb 71 02 04 04 32  .........U.q...2
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 03 00 00 00                          ........
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*chan int-+0
	rel 28+4 t=6 type.*chan int+0
	rel 32+4 t=1 type.int+0
type..namedata.*struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }- SRODATA dupok size=97
	0x0000 00 00 5e 2a 73 74 72 75 63 74 20 7b 20 63 20 75  ..^*struct { c u
	0x0010 6e 73 61 66 65 2e 50 6f 69 6e 74 65 72 3b 20 65  nsafe.Pointer; e
	0x0020 6c 65 6d 20 75 6e 73 61 66 65 2e 50 6f 69 6e 74  lem unsafe.Point
	0x0030 65 72 3b 20 6b 69 6e 64 20 75 69 6e 74 31 36 3b  er; kind uint16;
	0x0040 20 70 63 20 75 69 6e 74 70 74 72 3b 20 72 65 6c   pc uintptr; rel
	0x0050 65 61 73 65 74 69 6d 65 20 75 69 6e 74 36 34 20  easetime uint64
	0x0060 7d                                               }
type.*struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 3f 06 8e 78 00 04 04 36  ........?..x...6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 32+4 t=1 type.noalg.struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
runtime.gcbits.03 SRODATA dupok size=1
	0x0000 03                                               .
type..namedata.c- SRODATA dupok size=4
	0x0000 00 00 01 63                                      ...c
type..namedata.elem- SRODATA dupok size=7
	0x0000 00 00 04 65 6c 65 6d                             ...elem
type..namedata.kind- SRODATA dupok size=7
	0x0000 00 00 04 6b 69 6e 64                             ...kind
type..namedata.pc- SRODATA dupok size=5
	0x0000 00 00 02 70 63                                   ...pc
type..namedata.releasetime- SRODATA dupok size=14
	0x0000 00 00 0b 72 65 6c 65 61 73 65 74 69 6d 65        ...releasetime
type.noalg.struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=108
	0x0000 18 00 00 00 08 00 00 00 f0 9d 7c 2a 02 04 04 19  ..........|*....
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 05 00 00 00 05 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 08 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 10 00 00 00 00 00 00 00 00 00 00 00 18 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00 20 00 00 00              ........ ...
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.03+0
	rel 24+4 t=5 type..namedata.*struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 28+4 t=6 type.*struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
	rel 32+4 t=1 type..importpath."".+0
	rel 36+4 t=1 type.noalg.struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+48
	rel 48+4 t=1 type..namedata.c-+0
	rel 52+4 t=1 type.unsafe.Pointer+0
	rel 60+4 t=1 type..namedata.elem-+0
	rel 64+4 t=1 type.unsafe.Pointer+0
	rel 72+4 t=1 type..namedata.kind-+0
	rel 76+4 t=1 type.uint16+0
	rel 84+4 t=1 type..namedata.pc-+0
	rel 88+4 t=1 type.uintptr+0
	rel 96+4 t=1 type..namedata.releasetime-+0
	rel 100+4 t=1 type.uint64+0
type..namedata.*[]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }- SRODATA dupok size=99
	0x0000 00 00 60 2a 5b 5d 73 74 72 75 63 74 20 7b 20 63  ..`*[]struct { c
	0x0010 20 75 6e 73 61 66 65 2e 50 6f 69 6e 74 65 72 3b   unsafe.Pointer;
	0x0020 20 65 6c 65 6d 20 75 6e 73 61 66 65 2e 50 6f 69   elem unsafe.Poi
	0x0030 6e 74 65 72 3b 20 6b 69 6e 64 20 75 69 6e 74 31  nter; kind uint1
	0x0040 36 3b 20 70 63 20 75 69 6e 74 70 74 72 3b 20 72  6; pc uintptr; r
	0x0050 65 6c 65 61 73 65 74 69 6d 65 20 75 69 6e 74 36  eleasetime uint6
	0x0060 34 20 7d                                         4 }
type.*[]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 a5 f5 87 b3 00 04 04 36  ...............6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 32+4 t=1 type.[]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
type.[]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=36
	0x0000 0c 00 00 00 04 00 00 00 66 98 f1 9f 02 04 04 17  ........f.......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 28+4 t=6 type.*[]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
	rel 32+4 t=1 type.noalg.struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
type..namedata.*[2]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }- SRODATA dupok size=100
	0x0000 00 00 61 2a 5b 32 5d 73 74 72 75 63 74 20 7b 20  ..a*[2]struct {
	0x0010 63 20 75 6e 73 61 66 65 2e 50 6f 69 6e 74 65 72  c unsafe.Pointer
	0x0020 3b 20 65 6c 65 6d 20 75 6e 73 61 66 65 2e 50 6f  ; elem unsafe.Po
	0x0030 69 6e 74 65 72 3b 20 6b 69 6e 64 20 75 69 6e 74  inter; kind uint
	0x0040 31 36 3b 20 70 63 20 75 69 6e 74 70 74 72 3b 20  16; pc uintptr;
	0x0050 72 65 6c 65 61 73 65 74 69 6d 65 20 75 69 6e 74  releasetime uint
	0x0060 36 34 20 7d                                      64 }
type.*[2]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=36
	0x0000 04 00 00 00 04 00 00 00 4a 03 db 90 00 04 04 36  ........J......6
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00                                      ....
	rel 16+4 t=1 runtime.algarray+32
	rel 20+4 t=1 runtime.gcbits.01+0
	rel 24+4 t=5 type..namedata.*[2]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 32+4 t=1 type.noalg.[2]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
runtime.gcbits.c3 SRODATA dupok size=1
	0x0000 c3                                               .
type.noalg.[2]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 } SRODATA dupok size=44
	0x0000 30 00 00 00 20 00 00 00 32 aa 86 16 02 04 04 11  0... ...2.......
	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0020 00 00 00 00 00 00 00 00 02 00 00 00              ............
	rel 16+4 t=1 runtime.algarray+0
	rel 20+4 t=1 runtime.gcbits.c3+0
	rel 24+4 t=5 type..namedata.*[2]struct { c unsafe.Pointer; elem unsafe.Pointer; kind uint16; pc uintptr; releasetime uint64 }-+0
	rel 28+4 t=6 type.*[2]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
	rel 32+4 t=1 type.noalg.struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
	rel 36+4 t=1 type.[]struct { "".c unsafe.Pointer; "".elem unsafe.Pointer; "".kind uint16; "".pc uintptr; "".releasetime uint64 }+0
gclocals·5e326a2f5498e528ce8fbe7bd86e7d21 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 03 03 00                 ...........
gclocals·330e121c4652a955d2e0ff6e9125912d SRODATA dupok size=14
	0x0000 03 00 00 00 0c 00 00 00 00 00 c3 00 00 00        ..............
gclocals·e0b5c08eb812f276885e532abca71746 SRODATA dupok size=14
	0x0000 06 00 00 00 07 00 00 00 00 40 02 08 10 01        .........@....
gclocals·f6bd6b3389b872033d462029172c8612 SRODATA dupok size=8
	0x0000 04 00 00 00 00 00 00 00                          ........
gclocals·a57540db91d4df0a95a343aea8a16f48 SRODATA dupok size=12
	0x0000 04 00 00 00 02 00 00 00 00 02 03 01              ............
gclocals·f9e9e71b3d85c7d9ecf49fc49dd56c0d SRODATA dupok size=13
	0x0000 05 00 00 00 03 00 00 00 00 01 02 03 04           .............
gclocals·20fcec35d1696a32d1314d1646792dfe SRODATA dupok size=12
	0x0000 04 00 00 00 03 00 00 00 00 04 05 02              ............
gclocals·dc9b0298814590ca3ffc3a889546fc8b SRODATA dupok size=10
	0x0000 02 00 00 00 02 00 00 00 03 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
gclocals·1cf923758aae2e428391d1783fe59973 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 00 01 02                 ...........
gclocals·fdbf1f5761f6d17e8ae3f0aaecb6a3c5 SRODATA dupok size=11
	0x0000 03 00 00 00 02 00 00 00 03 02 00                 ...........
gclocals·7d2d5fca80364273fb07d5820a76fef4 SRODATA dupok size=8
	0x0000 03 00 00 00 00 00 00 00                          ........
gclocals·cd666f9a7f09fcd2aca7dadbf3522159 SRODATA dupok size=11
	0x0000 03 00 00 00 07 00 00 00 00 40 01                 .........@.
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
