package baby

import (
	"fmt"
	"regexp"
	"strconv"
	"bytes"
	"strings"
	"errors"
)

var lineRegex = regexp.MustCompile("[0-9]* *(NUM|JMP|JRP|LDN|STO|SUB|CMP|STP)( *)?(-?[0-9]+)?")

var (
	ErrNonInstruction = errors.New("trying to execute non-instruction code")
	ErrUnknownOpcode = errors.New("unknown opcode")
)

type Baby struct {
	CurrentInstruction uint32
	Accumulator        uint32

	// Bits 0-4 make up an optional address, while
	// the actual instruction lives in 13-15
	Memory [32]uint32
}

func (b *Baby) Run() error {
	for {
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
			return nil
		default:
			return ErrNonInstruction
		}
	}
}

func instrToOpCode(instr string) (uint32, error) {
	matches := lineRegex.FindStringSubmatch(instr)[1:]

	if len(matches) < 1 {
		return 0, ErrUnknownOpcode
	}

	// We ignore the possible error,
	// to keep the code simpler.
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

	return uint32(opCode), nil
}

func MemoryFromString(prog string) ([32]uint32, error) {
	var err error
	lines := strings.Split(prog, "\n")

	var memory [32]uint32
	memoryIndex := 0

	for i := 0; i < len(lines) && memoryIndex < len(memory); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}

		if !strings.HasPrefix(line, ";") {
			memory[memoryIndex], err = instrToOpCode(line)
			if err != nil {
				return memory, err
			}
			memoryIndex++
		}
	}

	return memory, nil
}

func MemoryToString(mem [32]uint32) string {
	buf := bytes.NewBuffer([]byte{})
	for i, line := range mem {
		fmt.Fprintf(buf, "%02d %032b % 10d\n", i, line, int32(line))
	}

	return buf.String()
}
