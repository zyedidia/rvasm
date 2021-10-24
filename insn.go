package rvasm

type InsnLayout int

const (
	RLayout InsnLayout = iota
	JLayout
	ULayout
	BLayout
	SLayout
	ILayout
	FLayout // fence instructions
)

type InsnType int

const (
	RType InsnType = iota
	JType
	IType
	SType
	LType
	EType
	FType
	CType
	BType
	UType
	PType
)

type Insn struct {
	Name   string
	Layout InsnLayout
	Type   InsnType
}

var Insns = map[string]Insn{
	"lui":   Insn{"lui", ULayout, UType},
	"auipc": Insn{"auipc", ULayout, UType},
	"jal":   Insn{"jal", JLayout, JType},
	"jalr":  Insn{"jalr", ILayout, IType},
	"beq":   Insn{"beq", BLayout, BType},
	"bne":   Insn{"bne", BLayout, BType},
	"blt":   Insn{"blt", BLayout, BType},
	"bge":   Insn{"bge", BLayout, BType},
	"bltu":  Insn{"bltu", BLayout, BType},
	"bgeu":  Insn{"bgeu", BLayout, BType},

	"lb":  Insn{"lb", ILayout, LType},
	"lh":  Insn{"lh", ILayout, LType},
	"lw":  Insn{"lw", ILayout, LType},
	"lbu": Insn{"lbu", ILayout, LType},
	"lhu": Insn{"lhu", ILayout, LType},

	"sb": Insn{"sb", SLayout, SType},
	"sh": Insn{"sh", SLayout, SType},
	"sw": Insn{"sw", SLayout, SType},

	"addi":  Insn{"addi", ILayout, IType},
	"slti":  Insn{"slti", ILayout, IType},
	"sltiu": Insn{"sltiu", ILayout, IType},
	"xori":  Insn{"xori", ILayout, IType},
	"ori":   Insn{"ori", ILayout, IType},
	"andi":  Insn{"andi", ILayout, IType},

	"slli": Insn{"slli", RLayout, IType},
	"srli": Insn{"srli", RLayout, IType},
	"srai": Insn{"srai", RLayout, IType},

	"add":  Insn{"add", RLayout, RType},
	"sub":  Insn{"sub", RLayout, RType},
	"sll":  Insn{"sll", RLayout, RType},
	"slt":  Insn{"slt", RLayout, RType},
	"sltu": Insn{"sltu", RLayout, RType},
	"xor":  Insn{"xor", RLayout, RType},
	"srl":  Insn{"srl", RLayout, RType},
	"sra":  Insn{"sra", RLayout, RType},
	"or":   Insn{"or", RLayout, RType},
	"and":  Insn{"and", RLayout, RType},

	"fence": Insn{"fence", FLayout, FType},

	"ecall":  Insn{"ecall", ILayout, EType},
	"ebreak": Insn{"ebreak", ILayout, EType},

	"csrrw":  Insn{"csrrw", ILayout, CType},
	"csrrs":  Insn{"csrrs", ILayout, CType},
	"csrrc":  Insn{"csrrc", ILayout, CType},
	"csrrwi": Insn{"csrrwi", ILayout, CType},
	"csrrsi": Insn{"csrrsi", ILayout, CType},
	"csrrci": Insn{"csrrci", ILayout, CType},

	"uret": Insn{"uret", RLayout, PType},
	"sret": Insn{"sret", RLayout, PType},
	"mret": Insn{"mret", RLayout, PType},
	"wfi":  Insn{"wfi", RLayout, PType},
}

var shifti = map[string]bool{
	"slli": true,
	"srli": true,
	"srai": true,
}

var prs2 = map[string]uint32{
	"uret": 0b00010,
	"sret": 0b00010,
	"mret": 0b00010,
	"wfi":  0b00101,
}

func (i Insn) Encode(rd, rs1, rs2, imm uint32) uint32 {
	op := i.Op()
	switch i.Layout {
	case RLayout:
		return CatBits(
			Bits{funct7[i.Name], 7},
			Bits{rs2, 5},
			Bits{rs1, 5},
			Bits{funct3[i.Name], 3},
			Bits{rd, 5},
			Bits{op, 7},
		).Uint32()
	case JLayout:
		return CatBits(
			GetBits(imm, 20, 20),
			GetBits(imm, 10, 1),
			GetBits(imm, 11, 11),
			GetBits(imm, 19, 12),
			Bits{rd, 5},
			Bits{op, 7},
		).Uint32()
	case ULayout:
		return CatBits(
			GetBits(imm, 19, 0),
			Bits{rd, 5},
			Bits{op, 7},
		).Uint32()
	case BLayout:
		return CatBits(
			GetBits(imm, 12, 12),
			GetBits(imm, 10, 5),
			Bits{rs2, 5},
			Bits{rs1, 5},
			Bits{funct3[i.Name], 3},
			GetBits(imm, 4, 1),
			GetBits(imm, 11, 11),
			Bits{op, 7},
		).Uint32()
	case SLayout:
		return CatBits(
			GetBits(imm, 11, 5),
			Bits{rs2, 5},
			Bits{rs1, 5},
			Bits{funct3[i.Name], 3},
			GetBits(imm, 4, 0),
			Bits{op, 7},
		).Uint32()
	case ILayout:
		return CatBits(
			Bits{imm, 12},
			Bits{rs1, 5},
			Bits{funct3[i.Name], 3},
			Bits{rd, 5},
			Bits{op, 7},
		).Uint32()
	case FLayout:
	}
	panic("unimplemented")
}
