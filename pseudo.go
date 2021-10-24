package rvasm

import (
	"fmt"
	"strings"
)

func isize(inst *Inst) uint32 {
	switch inst.Name {
	case "call", "tail", "la",
		"lb", "lh", "lw", "ld",
		"sb", "sh", "sw", "sd":
		if IsPseudo(inst) {
			return 8
		}
	}
	return 4
}

func IsPseudo(inst *Inst) bool {
	if !pseudo[inst.Name] {
		return false
	}

	switch inst.Name {
	case "lb", "lh", "lw", "ld",
		"sb", "sh", "sw", "sd":
		return inst.Args[1].Ma == nil
	case "jal", "jalr":
		return len(inst.Args) == 1
	case "fence":
		return len(inst.Args) == 0
	}
	return true
}

func FromPseudo(inst *Inst, labels map[string]uint32, base uint32) []*Operation {
	var code string
	switch inst.Name {
	case "la":
		panic("unimplemented")
	case "lb", "lh", "lw", "ld":
		panic("unimplemented")
	case "sb", "sh", "sw", "sd":
		panic("unimplemented")
	case "nop":
		code = "addi x0, x0, 0"
	case "li":
		// TODO: not quite right
		code = fmt.Sprintf("addi %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "mv":
		code = fmt.Sprintf("addi %s, %s, 0", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "not":
		code = fmt.Sprintf("xori %s, %s, -1", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "neg":
		code = fmt.Sprintf("sub %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "seqz":
		code = fmt.Sprintf("sltiu %s, %s, 1", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "snez":
		code = fmt.Sprintf("sltu %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "sltz":
		code = fmt.Sprintf("slt %s, %s, x0", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "sgtz":
		code = fmt.Sprintf("slt %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "beqz":
		code = fmt.Sprintf("beq %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "bnez":
		code = fmt.Sprintf("bne %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "blez":
		code = fmt.Sprintf("bge x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "bgez":
		code = fmt.Sprintf("bge %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "bltz":
		code = fmt.Sprintf("blt %s, x0, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "bgtz":
		code = fmt.Sprintf("blt x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "bgt":
		code = fmt.Sprintf("blt %s, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val, inst.Args[2].Na.Val)
	case "ble":
		code = fmt.Sprintf("bge %s, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val, inst.Args[2].Na.Val)
	case "bgtu":
		code = fmt.Sprintf("bltu %s, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val, inst.Args[2].Na.Val)
	case "bleu":
		code = fmt.Sprintf("bgeu %s, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val, inst.Args[2].Na.Val)
	case "j":
		code = fmt.Sprintf("jal x0, %s", inst.Args[0].Na.Val)
	case "jal":
		code = fmt.Sprintf("jal x1, %s", inst.Args[0].Na.Val)
	case "jr":
		code = fmt.Sprintf("jalr x0, %s, 0", inst.Args[0].Na.Val)
	case "jalr":
		code = fmt.Sprintf("jalr x1, %s, 0", inst.Args[0].Na.Val)
	case "ret":
		code = "jalr x0, x1, 0"
	case "call":
		code = fmt.Sprintf("auipc x6, %d\njalr x1, x6, %d",
			GetBits(inst.Args[0].Na.Imm(labels, 0)+base, 31, 12).Uint32(),
			GetBits(inst.Args[0].Na.Imm(labels, 0)+base, 11, 0).Uint32(),
		)
	case "tail":
		code = fmt.Sprintf("auipc x6, %d\njalr x0, x6, %d",
			GetBits(inst.Args[0].Na.Imm(labels, 0)+base, 31, 12).Uint32(),
			GetBits(inst.Args[0].Na.Imm(labels, 0)+base, 11, 0).Uint32(),
		)
	case "fence":
		code = "fence iorw, iorw"
	case "rdinstret":
		code = fmt.Sprintf("csrrs %s, instret, x0", inst.Args[0].Na.Val)
	case "rdinstreth":
		code = fmt.Sprintf("csrrs %s, instreth, x0", inst.Args[0].Na.Val)
	case "rdcycle":
		code = fmt.Sprintf("csrrs %s, cycle, x0", inst.Args[0].Na.Val)
	case "rdcycleh":
		code = fmt.Sprintf("csrrs %s, cycleh, x0", inst.Args[0].Na.Val)
	case "rdtime":
		code = fmt.Sprintf("csrrs %s, time, x0", inst.Args[0].Na.Val)
	case "rdtimeh":
		code = fmt.Sprintf("csrrs %s, timeh, x0", inst.Args[0].Na.Val)
	case "csrr":
		code = fmt.Sprintf("csrrs %s, %s, x0", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrw":
		code = fmt.Sprintf("csrrw x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrs":
		code = fmt.Sprintf("csrrs x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrc":
		code = fmt.Sprintf("csrrc x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrwi":
		code = fmt.Sprintf("csrrwi x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrsi":
		code = fmt.Sprintf("csrrsi x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	case "csrci":
		code = fmt.Sprintf("csrrci x0, %s, %s", inst.Args[0].Na.Val, inst.Args[1].Na.Val)
	}
	ast, err := Parse(fmt.Sprintf("pseudo:%s", inst.Name), strings.NewReader(code))
	if err != nil {
		panic(fmt.Sprintf("pseudo-instruction conversion error %v", err))
	}
	return ast.Ops
}

var pseudo = map[string]bool{
	"la": true,
	"lb": true,
	"lh": true,
	"lw": true,
	"ld": true,
	"sb": true,
	"sh": true,
	"sw": true,
	"sd": true,

	"nop":  true,
	"li":   true,
	"mv":   true,
	"not":  true,
	"neg":  true,
	"seqz": true,
	"snez": true,
	"sltz": true,
	"sgtz": true,

	"beqz":  true,
	"bnez":  true,
	"blez":  true,
	"bgez":  true,
	"bgltz": true,
	"bgtz":  true,

	"bgt":  true,
	"ble":  true,
	"bgtu": true,
	"bleu": true,

	"j":    true,
	"jal":  true,
	"jr":   true,
	"jalr": true,
	"ret":  true,
	"call": true,
	"tail": true,

	"fence": true,

	"rdinstret":  true,
	"rdinstreth": true,
	"rdcycle":    true,
	"rdcycleh":   true,
	"rdtime":     true,
	"rdtimeh":    true,

	"csrr":  true,
	"csrw":  true,
	"csrs":  true,
	"csrc":  true,
	"csrwi": true,
	"csrsi": true,
	"csrci": true,
}
