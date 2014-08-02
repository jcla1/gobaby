package main

import (
    "github.com/jcla1/gobaby/baby"
    "io/ioutil"
    "fmt"
)

func main() {
    fc, err := ioutil.ReadFile("examples/primegen.asm")
    if err != nil {
        panic(err)
    }

    mem, err := baby.MemoryFromString(string(fc))
    if err != nil {
        panic(err)
    }

    b := baby.Baby{0, 0, mem}

    for i := 0; i < 20; i++ {
        b.Run()
        fmt.Println("prime is:", int32(b.Memory[21]))
        b.Accumulator = 0
        b.CurrentInstruction = 0
    }

    fmt.Print(baby.MemoryToString(b.Memory))
}