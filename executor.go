package baby

type Baby struct {
    CurrentInstruction int8
    Accumulator int32

    // Bits 0-4 make up an optional address, while
    // the actual instruction lives in 13-15
    Memory [32]int32
}
