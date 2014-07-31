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

    instr := b.Memory[b.CurrentInstruction]
    data := instr & 0x0000001F

    switch instr & 0x0000F000 {
    case 0x00000000: // JMP
        b.CurrentInstruction = b.Memory[data]
    case 0x00002000: // JRP
        b.CurrentInstruction += b.Memory[data]
    case 0x00004000: // LDN
        b.Accumulator = -b.Memory[data]
    case 0x00006000: // STO
        b.Memory[data] = b.Accumulator
    case 0x00008000: // SUB
        b.Accumulator -= b.Memory[data]
    case 0x0000C000: // CMP
        if (b.Accumulator < 0) {
            b.CurrentInstruction++
        }
    case 0x0000E000: // STP
        return
    default:
        panic("trying to execute non-instruction code!")
    }
}

func (b *Baby) Run() {
    for (b.Memory[b.CurrentInstruction] & 0x0000F000) != 0x0000E000 {
        b.Step()
    }
}