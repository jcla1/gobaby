; Run with: gobaby -l 21 examples/primegen.asm to get the next prime number.
; This is a chainable program, which means if you just keep reexecuting the
; memory, you'll get successive numbers.
; Here, for example, we get the 3rd prime number:
;     gobaby examples/primegen.asm | gobaby | gobaby -l 21

00  JMP 24
01  LDN 21
02  STO 21
03  LDN 21
04  SUB 15
05  STO 21
06  LDN 15
07  STO 22
08  LDN 22
09  STO 22
10  LDN 22
11  SUB 15
12  STO 22
13  SUB 21
14  CMP
15  NUM -1
16  LDN 21
17  STO 23
18  LDN 23
19  SUB 22
20  JMP  0
21  JMP  1
22  JMP  0
23  JMP  0
24  JMP  7
25  CMP
26  JRP  0
27  STO 23
28  LDN 23
29  SUB 22
30  CMP
31  JMP 20
