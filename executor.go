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
        b.CurrentInstruction = data
    case 0x00002000: // CMP
        fmt.Printf("JRP %d\n", data)
    case 0x00004000: // LDN
        fmt.Printf("LDN %d\n", data)
    case 0x00006000: // STO
        fmt.Printf("STO %d\n", data)
    case 0x00008000: // SUB
        fmt.Printf("SUB %d\n", data)
    case 0x0000C000: // CMP
        fmt.Printf("CMP\n")
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