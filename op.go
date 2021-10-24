package rvasm

var opinsn = map[string]uint32{
	"lui":   0b0110111,
	"auipc": 0b0010111,
	"jalr":  0b1100111,
}

var optype = map[InsnType]uint32{
	RType: 0b0110011,
	IType: 0b0010011,
	SType: 0b0100011,
	LType: 0b0000011,
	BType: 0b1100011,
	EType: 0b1110011,
	FType: 0b0001111,
	CType: 0b1110011,
	JType: 0b1101111,
	PType: 0b1110011,
}

func (i Insn) Op() uint32 {
	if op, ok := opinsn[i.Name]; ok {
		return op
	}
	return optype[i.Type]
}
