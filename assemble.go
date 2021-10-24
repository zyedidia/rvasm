package rvasm

import (
	"fmt"
	"io"
)

type Program []uint32

func Assemble(fname string, r io.Reader, base uint32) (Program, error) {
	ast, err := Parse(fname, r)
	if err != nil {
		return nil, err
	}
	return AssembleAST(ast, base), nil
}

func AssembleAST(ast *RV32i, base uint32) Program {
	labels := make(map[string]uint32)

	// resolve labels
	var addr uint32
	for _, op := range ast.Ops {
		if op.Label != "" {
			labels[op.Label] = addr
		} else {
			addr += isize(op.Inst)
		}
	}

	// resolve pseudo-instructions
	newops := make([]*Operation, 0, len(ast.Ops))
	for _, op := range ast.Ops {
		if op.Label != "" {
			continue
		}
		if IsPseudo(op.Inst) {
			newops = append(newops, FromPseudo(op.Inst, labels, base)...)
		} else {
			newops = append(newops, op)
		}
	}

	addr = 0
	prog := make(Program, 0, len(ast.Ops))
	for _, op := range newops {
		if op.Label != "" {
			continue
		}
		var rd, rs1, rs2, imm uint32
		insn := Insns[op.Inst.Name]
		switch insn.Type {
		case BType:
			rs1 = op.Inst.Args[0].Na.Reg()
			rs2 = op.Inst.Args[1].Na.Reg()
			imm = op.Inst.Args[2].Na.Imm(labels, addr)
		case SType:
			rs2 = op.Inst.Args[0].Na.Reg()
			rs1 = Reg[op.Inst.Args[1].Ma.Reg]
			imm = uint32(op.Inst.Args[1].Ma.Imm)
		case LType:
			rd = op.Inst.Args[0].Na.Reg()
			rs1 = Reg[op.Inst.Args[1].Ma.Reg]
			imm = uint32(op.Inst.Args[1].Ma.Imm)
		case RType:
			rd = op.Inst.Args[0].Na.Reg()
			rs1 = op.Inst.Args[1].Na.Reg()
			rs2 = op.Inst.Args[2].Na.Reg()
		case CType:
			rd = op.Inst.Args[0].Na.Reg()
			rs1 = op.Inst.Args[2].Na.Reg()
			if csr, ok := Csr[op.Inst.Args[1].Na.Val]; ok {
				imm = csr
			} else {
				imm = op.Inst.Args[1].Na.Imm(labels, 0)
			}
		case IType:
			rd = op.Inst.Args[0].Na.Reg()
			rs1 = op.Inst.Args[1].Na.Reg()
			imm = op.Inst.Args[2].Na.Imm(labels, 0)
		case UType:
			rd = op.Inst.Args[0].Na.Reg()
			imm = op.Inst.Args[1].Na.Imm(labels, 0)
		case JType:
			rd = op.Inst.Args[0].Na.Reg()
			imm = (op.Inst.Args[1].Na.Imm(labels, 0) + base) / 8
		case EType:
			imm = funct12[op.Inst.Name]
		case PType:
			rs2 = prs2[op.Inst.Name]
		}
		if shifti[op.Inst.Name] {
			rs2 = imm
		}
		prog = append(prog, insn.Encode(rd, rs1, rs2, imm))
		addr += isize(op.Inst)
	}

	return prog
}

func (p Program) EncodeToHex() []string {
	s := make([]string, len(p))
	for i, insn := range p {
		s[i] = fmt.Sprintf("%08x", insn)
	}
	return s
}
