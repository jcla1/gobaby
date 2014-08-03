package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jcla1/gobaby/baby"
)

var printLoc = flag.Int("l", 31, "memory location whose value shall be printed, range -1 - 31")
var printMem = flag.Bool("p", false, "print out the memory as a program, after execution")

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s filename\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Fprintln(os.Stderr, "not enough arguments given")
		flag.Usage()
	}

	if *printLoc < -1 || *printLoc > 31 {
		fmt.Fprintln(os.Stderr, "can't print outside of memory")
		flag.Usage()
	}

	fc, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	mem, err := baby.MemoryFromString(string(fc))
	if err != nil {
		panic(err)
	}

	b := baby.Baby{0, 0, mem}
	b.Run()

	if *printLoc > -1 {
		fmt.Printf("Value at location #%d: %d\n", *printLoc, int32(b.MemoryImage[*printLoc]))
	}

	if *printMem {
		fmt.Print(b.ASM())
	}
}
