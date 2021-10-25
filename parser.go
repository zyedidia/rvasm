package rvasm

import (
	"fmt"
	"io"
	"strconv"

	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var rv32iLexer = lexer.MustSimple([]lexer.Rule{
	{`Ident`, `[_@.a-zA-Z][a-zA-Z_\d]*`, nil},
	{`Number`, `[-+]?(0x)?[0-9a-fA-F]+\b`, nil},
	{`Punct`, `[,:()]`, nil},
	{`Newline`, `\r?\n`, nil},
	{"comment", `//.*|/\*.*?\*/`, nil},
	{"whitespace", `\s+`, nil},
})

var rv32iParser = participle.MustBuild(&RV32i{},
	participle.Lexer(rv32iLexer),
)

type RV32i struct {
	Pos lexer.Position

	Ops []*Operation `@@*`
}

type Operation struct {
	Label string `  @Ident ":" Newline*`
	Inst  *Inst  `| @@`
}

type Inst struct {
	Name string `@Ident`
	Args []*Arg `( @@ ( "," @@ )* )? (Newline+|EOF)`
}

type Arg struct {
	Ma *MemArg    `  @@`
	Na *NormalArg `| @@`
}

type MemArg struct {
	Imm int    `@Number`
	Reg string `"(" @Ident ")"`
}

type NormalArg struct {
	Val string `@(Ident | Number)`
}

func (n *NormalArg) IsReg() bool {
	_, ok := Reg[n.Val]
	return ok
}

func (n *NormalArg) Reg() uint32 {
	// TODO: fix this
	if !n.IsReg() {
		panic(fmt.Sprintf("%s is not a reg", n.Val))
	}
	return Reg[n.Val]
}

func (n *NormalArg) Imm(locs map[string]uint32, rel uint32) uint32 {
	if v, err := strconv.ParseInt(n.Val, 0, 32); err == nil {
		return uint32(v)
	} else if v, err := strconv.ParseUint(n.Val, 0, 32); err == nil {
		return uint32(v)
	}
	// TODO: fix this
	if l, ok := locs[n.Val]; ok {
		return l - rel
	}
	panic(fmt.Sprintf("invalid label: %s", n.Val))
}

func (o *Operation) String() string {
	if o.Label != "" {
		return fmt.Sprintf("%s:", o.Label)
	}
	return o.Inst.String()
}

func (i *Inst) String() string {
	s := i.Name + " "
	for j, a := range i.Args {
		s += a.String()
		if j != len(i.Args)-1 {
			s += ", "
		}
	}
	return s
}

func (a *Arg) String() string {
	if a.Ma != nil {
		return a.Ma.String()
	}
	return a.Na.String()
}

func (m *MemArg) String() string {
	return fmt.Sprintf("%d(%s)", m.Imm, m.Reg)
}

func (n *NormalArg) String() string {
	return n.Val
}

func Parse(name string, r io.Reader) (*RV32i, error) {
	rv32i := &RV32i{}
	err := rv32iParser.Parse(name, r, rv32i)
	return rv32i, err
}
