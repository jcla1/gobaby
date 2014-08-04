package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/jcla1/gobaby/baby"
)

var printLoc = flag.Int("l", -1, "memory location whose value shall be printed, range -1 - 31")
var printMem = flag.Bool("p", true, "print out the memory as a program, after execution")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s filename\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *printLoc < -1 || *printLoc > 31 {
		fmt.Fprintln(os.Stderr, "can't print outside of memory")
		flag.Usage()
	}

	var fd io.Reader
	var err error
	if flag.NArg() < 1 {
		fd = os.Stdin
	} else {
		fd, err = os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to open file: %s\n\n", flag.Arg(0))
		}
	}

	fc, err := ioutil.ReadAll(fd)
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
		fmt.Printf("Value at location #%02d: %d\n", *printLoc, int32(b.MemoryImage[*printLoc]))
	}

	if *printMem {
		fmt.Print(b.ASM())
	}
}
