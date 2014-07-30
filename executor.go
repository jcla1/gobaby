package baby

import (
   "fmt"
)

type Baby struct {
    CurrentInstruction int8
    Accumulator int32

    // Bits 0-4 make up an optional address, while
    // the actual instruction lives in 13-15
    Memory [32]int32
}

func (b *Baby) Step() {
    b.CurrentInstruction++

    switch instr := b.Memory[b.CurrentInstruction]; {
    case (instr & 0x0000000F) == 0: // JMP
        fmt.Printf("%02d JMP %d\n", b.CurrentInstruction, instr & 0x0000000F)
    }
}