package baby

import (
    "regexp"
    "strconv"
)

var lineRegex = regexp.MustCompile("[0-9]* (JMP|JRP|LDN|STO|SUB|CMP|STP) ?([0-9]+)?")

type Baby struct {
    CurrentInstruction int32
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

func InstrToOpCode(instr string) int32 {
    matches := lineRegex.FindStringSubmatch(instr)[1:]

    if len(matches) < 1 {
        panic("unknown opcode!")
    }

    opCode, _ := strconv.Atoi(matches[2])

    switch matches[1] {
    case "JMP": break
    case "JRP": opCode |= 0x00002000
    case "LDN": opCode |= 0x00004000
    case "STO": opCode |= 0x00006000
    case "SUB": opCode |= 0x00008000
    case "CMP": opCode |= 0x0000C000
    case "STP": opCode |= 0x0000E000
    }

    return int32(opCode)
}
