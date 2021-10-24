package rvasm_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zyedidia/rvasm"
)

func dump(hex, expect []string) {
	for i := 0; i < max(len(hex), len(expect)); i++ {
		if i < len(hex) {
			fmt.Print(hex[i])
			fmt.Print(" ")
		}
		if i < len(expect) {
			fmt.Print(expect[i])
			fmt.Print(" ")
		}
		if i < len(expect) && i < len(hex) && hex[i] != expect[i] {
			fmt.Print("X")
		}
		fmt.Print("\n")
	}
}

func check(fname string, asm io.Reader, expect []string, t *testing.T) {
	prog, err := rvasm.Assemble(fname, asm, 0)
	if err != nil {
		t.Fatal(err)
	}
	hex := prog.EncodeToHex()
	if len(hex) != len(expect) {
		dump(hex, expect)
		os.Exit(1)
	}
	for i, s := range hex {
		if s != expect[i] {
			dump(hex, expect)
			os.Exit(1)
		}
	}
}

func TestAssemble(t *testing.T) {
	err := filepath.Walk("testdata",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".s") {
				mem := path[:len(path)-1] + "mem"
				memstr, err := ioutil.ReadFile(mem)
				if err != nil {
					return nil
				}
				expect := strings.Split(strings.TrimSpace(string(memstr)), "\n")
				f, err := os.Open(path)
				if err != nil {
					return nil
				}
				fmt.Printf("checking %s...", path)
				check(path, f, expect, t)
				fmt.Println(" PASS")
				f.Close()
			}
			return nil
		})
	if err != nil {
		t.Fatal(err)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
