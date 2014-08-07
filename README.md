Manchester Small-Scale Experimental Machine Emulator
====================================================

Emulator for the SSEM, aka: "the Baby", written in Go.

## Installation
You need to have Go installed, then just run:
```shell
$ go get github.com/jcla1/gobaby
...
$ go install github.com/jcla1/gobaby
```
Then you're ready to go!

## Running the examples
To run a program, you can either provide it via stdin or as an argument to ```gobaby```:

```bash
$ echo examples/factor.asm | gobaby -p=f -l 27
...
$ gobaby -p=f -l 27 examples/factor.asm
```

If you need any help, look at the usage with: ```gobaby -help```.
