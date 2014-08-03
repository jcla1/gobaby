package baby

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var lineRegex = regexp.MustCompile("(?:[0-9]*)? *(NUM|JMP|JRP|LDN|STO|SUB|CMP|STP)( *)?(-?[0-9]+)?")

var (
	ErrNonInstruction = errors.New("trying to execute non-instruction code")
	ErrUnknownOpcode  = errors.New("unknown opcode")
)

type MemoryImage [32]uint32

type Baby struct {
	CurrentInstruction uint32
	Accumulator        uint32

	// Bits 0-4 make up an optional address, while
	// the actual instruction lives in 13-15
	MemoryImage
}

func (b *Baby) Run() error {
	for {
		b.CurrentInstruction++

		instr := b.MemoryImage[int32(b.CurrentInstruction)%32]
		data := instr & 0x0000001F

		switch instr & 0x0000E000 {
		case 0x00000000: // JMP
			b.CurrentInstruction = b.MemoryImage[data]
		case 0x00002000: // JRP
			b.CurrentInstruction += b.MemoryImage[data]
		case 0x00004000: // LDN
			b.Accumulator = -b.MemoryImage[data]
		case 0x00006000: // STO
			b.MemoryImage[data] = b.Accumulator
		case 0x00008000: // SUB
			b.Accumulator -= b.MemoryImage[data]
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

func (b *Baby) Reset() {
	b.Accumulator = 0
	b.CurrentInstruction = 0
}

func (b *Baby) ASM() string {
	var buf bytes.Buffer
	for i, line := range b.MemoryImage {
		fmt.Fprintf(&buf, "%02d  %s\n", i, memLineToASM(line))
	}
	return buf.String()
}

func memLineToASM(line uint32) string {
	// If any of the high bits are set,
	// we're guessing that it's a NUM
	if (line & 0xFFFF0000) > 0 {
		return fmt.Sprintf("NUM %2d", int32(line))
	}

	data := int32(line & 0x0000001F)

	switch line & 0x0000E000 {
	case 0x00000000:
		return fmt.Sprintf("JMP %2d", data)
	case 0x00002000:
		return fmt.Sprintf("JRP %2d", data)
	case 0x00004000:
		return fmt.Sprintf("LDN %2d", data)
	case 0x00006000:
		return fmt.Sprintf("STO %2d", data)
	case 0x00008000:
		return fmt.Sprintf("SUB %2d", data)
	case 0x0000C000:
		return "CMP"
	case 0x0000E000:
		return "STP"
	}

	return ""
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

func MemoryFromString(prog string) (MemoryImage, error) {
	var err error
	lines := strings.Split(prog, "\n")

	var memory MemoryImage
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

func (mem MemoryImage) String() string {
	var buf bytes.Buffer
	for i, line := range mem {
		fmt.Fprintf(&buf, "%02d   ", i)

		chunks := make([]byte, 4)
		binary.BigEndian.PutUint32(chunks, line)
		for _, chunk := range chunks {
			fmt.Fprintf(&buf, "%08b ", chunk)
		}

		fmt.Fprintf(&buf, "   %10d\n", int32(line))
	}

	return buf.String()
}
