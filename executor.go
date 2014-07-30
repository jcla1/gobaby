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

    fmt.Printf("%02d ", b.CurrentInstruction)

    switch instr := b.Memory[b.CurrentInstruction]; instr & 0x0000F000 {
    case 0: // JMP
        fmt.Printf("JMP %d\n", instr & 0x0000000F)
    case 16384: // LDN
        fmt.Printf("LDN %d\n", instr & 0x0000000F)
    case 24576: // STO
        fmt.Printf("STO %d\n", instr & 0x0000000F)
    case 32768: // SUB
        fmt.Printf("SUB %d\n", instr & 0x0000000F)
    default:
        fmt.Printf("unknown instr: %d\n", instr & 0x0000F000)
    }
}