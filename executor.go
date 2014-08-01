package baby

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile("[0-9]* *(NUM|JMP|JRP|LDN|STO|SUB|CMP|STP)( *)?(-?[0-9]+)?")

type Baby struct {
	CurrentInstruction uint32
	Accumulator        uint32

	// Bits 0-4 make up an optional address, while
	// the actual instruction lives in 13-15
	Memory [32]uint32
}

func (b *Baby) Run() {
	for {
		// fmt.Printf("CI : %d\nACC: %d\n-----\n", int32(b.CurrentInstruction), int32(b.Accumulator))
		// fmt.Scanln()
		b.CurrentInstruction++

		instr := b.Memory[int32(b.CurrentInstruction)%32]
		data := instr & 0x0000001F

		switch instr & 0x0000E000 {
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
			if int32(b.Accumulator) < 0 {
				b.CurrentInstruction++
			}
		case 0x0000E000: // STP
			return
		default:
			fmt.Println("malicious position:", b.CurrentInstruction)
			panic("trying to execute non-instruction code!")
		}
	}
}

func instrToOpCode(instr string) uint32 {
	matches := lineRegex.FindStringSubmatch(instr)[1:]

	if len(matches) < 1 {
		panic("unknown opcode!")
	}

	opCode, _ := strconv.Atoi(matches[2])

	switch matches[0] {
	case "JMP", "NUM":
		break
	case "JRP":
		opCode |= 0x00002000
	case "LDN":
		opCode |= 0x00004000
	case "STO":
		opCode |= 0x00006000
	case "SUB":
		opCode |= 0x00008000
	case "CMP":
		opCode |= 0x0000C000
	case "STP":
		opCode |= 0x0000E000
	}
	// fmt.Printf("%032b\n", uint32(opCode))
	return uint32(opCode)
}

func MemoryFromString(prog string) [32]uint32 {
	lines := strings.Split(prog, "\n")

	var memory [32]uint32
	memoryIndex := 0

	for i := 0; i < len(lines) && memoryIndex < len(memory); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, ";") {
			memory[memoryIndex] = instrToOpCode(line)
			memoryIndex++
		}
	}

	return memory
}
