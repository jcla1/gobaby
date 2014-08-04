; Run with: gobaby -l 9 -p=f examples/simple_calc.asm
; Will perform the calculation: 5 - 3 = 2
; So output will be along the lines of:
;     Value at memory location #09: 2

; The lines are formatted in the following manner:
;     [word index] [3 letter mnemonic] [optional address or value parameter] [comment]

00  JMP  0
01  LDN  7          ; We load in our initial number, negatively
02  SUB  8          ; next, we just subtract
03  STO  9          ; and put out result back into memory.
04  STP             ; Finally, we need to make the machine STOP!

05  JMP  0          ; Take these "JMP 0" instructions as empty lines.
06  JMP  0

07  NUM -5          ; Here is the data
08  NUM  3          ; for our calculation!

09  JMP  0          ; This line is only included for clarity,
                    ; it will hold the result of the calculation.

; P.S. we're missing lines 10 - 31, which is just fine!