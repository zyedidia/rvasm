package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zyedidia/rvasm"
)

var base = flag.Int("base", 0, "base address")

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("no input file given")
		flag.Usage()
		return
	}

	fname := flag.Args()[0]
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	prog, err := rvasm.Assemble(fname, f, uint32(*base))
	if err != nil {
		log.Fatal(err)
	}
	hex := prog.EncodeToHex()
	fmt.Println(strings.Join(hex, "\n"))
}
