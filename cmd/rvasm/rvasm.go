package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/deadsy/rvda"
	"github.com/zyedidia/rvasm"
)

var base = flag.Int("base", 0, "base address")
var disas = flag.Bool("disas", false, "disassemble input")
var raw = flag.Bool("raw", false, "use raw machine code representation instead of hex")

func main() {
	flag.Parse()

	var f *os.File
	var fname string
	if len(flag.Args()) < 1 {
		f = os.Stdin
		fname = "stdin"
	} else {
		var err error
		fname := flag.Args()[0]
		f, err = os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *disas {
		isa, err := rvda.New(32, rvda.RV32gc)
		if err != nil {
			log.Fatal(err)
		}
		if *raw {
			insbytes := make([]byte, 4)
			i := 0
			for n, err := f.Read(insbytes); err == nil && n == 4; n, err = f.Read(insbytes) {
				ins := uint(insbytes[3])<<24 | uint(insbytes[2])<<16 | uint(insbytes[1])<<8 | uint(insbytes[0])
				fmt.Println(isa.Disassemble(uint(i*4+*base), ins))
				i++
			}
		} else {
			hex, err := io.ReadAll(f)
			if err != nil {
				log.Fatal(err)
			}
			insns := strings.Split(strings.TrimSpace(string(hex)), "\n")
			for i, ins := range insns {
				insbytes, err := strconv.ParseUint(ins, 16, 32)
				if err != nil {
					log.Fatalf("error:%d:%v", i+1, err)
				}
				fmt.Println(isa.Disassemble(uint(i*4+*base), uint(insbytes)))
			}
		}
	} else {
		prog, err := rvasm.Assemble(fname, f, uint32(*base))
		if err != nil {
			log.Fatal(err)
		}
		if *raw {
			mcode := prog.EncodeToBin()
			fmt.Print(string(mcode))
		} else {
			hex := prog.EncodeToHex()
			fmt.Println(strings.Join(hex, "\n"))
		}
	}
	f.Close()
}
