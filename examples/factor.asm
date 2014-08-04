00  JMP  0
; Initialization
01  LDN 24          ; Loads to Acc -(no. to be factored - 1) = initial -b test value
02  STO 26          ; Stores the initial –b test value in line 26
03  LDN 26          ; Loads initial +b value into Acc
04  STO 27          ; Stores the initial +b test value in line 27

; Do subtractions using the current b test value, check sign of difference, jump back if 0 is not passed yet
05  LDN 23          ; Loads in no. to be factored to Acc.
06  SUB 27          ; Subtracts the latest +b test value from the current Acc value
07  CMP             ; Jumps to execute line 9 if Acc is now negative.
08  JRP 20          ; Loops back to execute from line 6 again if Acc value not yet negative

; Form a remainder, Test it and Stop if it is Zero (because we have a result then)
09  SUB 26          ; Subtract current -b test value from Acc (so adds +b back on).
                    ; By adding +b back on, we identify if subtractions have overshot 0
                    ; by less than the amount +b, in which case b isn’t a factor.
                    ; If instead we get back to 0 exactly, then it must be a factor.
10  STO 25          ; Stores the calculated overshoot difference value in line 25.
                    ; If this is 0, we’ve found the factor. If it’s +ve, we haven’t.
11  LDN 25          ; Loads negative of line 25 overshoot difference value.
                    ; If a negative no is loaded now, then we haven’t got a factor.
                    ; If a non negative no is loaded it must be 0 and we have the factor.
12  CMP             ; If Acc is negative: Jumps to execute from 14 with a new test divisor.
                    ; If Acc is not negative: Execute next line 13 to Stop.
13  STP             ; STOP. Acc was NOT negative so Divisor was found. Answer is in Line 27.

; Form a new divisor b to be tested, then jump back and test it as a possible factor using subtractions
14  LDN 26          ; Load the last tested b value as a positive Acc value.
15  SUB 21          ; Decrement the last tested b value by 1.
16  STO 27          ; Store new +b test value in line 27.
17  LDN 27          ; Load new –b test value into Acc.
18  STO 26          ; Store new -b test value in line 26.
19  JMP 22          ; Execute subtractions from line 5 again using new test b value.

; Fixed data
20  NUM -3          ; Value for use in the JRP jump instruction in line 8.
21  NUM  1          ; Value for decrementing value of the test b value in line 15.
22  NUM  4          ; Value for use in the JMP jump instruction in line 19.
23  NUM -262144     ; Negative form of the number to be factored.
24  NUM 262143      ; First b value to check as being a factor of number in line 23.

; Variable data written to during execution (initially all zero)
25  NUM  0          ; Latest overshoot difference (written by line 10).
26  NUM  0          ; Latest -b value under test (written by line 2 or 18).
27  NUM  0          ; Latest +b value under test (written by line 4 or 16).*
28  NUM  0          ; Not used.
29  NUM  0          ; Not used.
30  NUM  0          ; Not used.
31  NUM  0          ; Not used.
