package rvasm

import "strconv"

func Reg2Str(r uint32) string {
	return "x" + strconv.Itoa(int(r))
}

var Csr = map[string]uint32{
	"ustatus":  0x000,
	"uie":      0x004,
	"utvec":    0x005,
	"uscratch": 0x040,
	"uepc":     0x041,
	"ucause":   0x042,
	"utval":    0x043,
	"uip":      0x044,

	"cycle":    0xC00,
	"time":     0xC01,
	"instret":  0xC02,
	"cycleh":   0xC80,
	"timeh":    0xC81,
	"instreth": 0xC82,
}

var Reg = map[string]uint32{
	"x0":  0,
	"x1":  1,
	"x2":  2,
	"x3":  3,
	"x4":  4,
	"x5":  5,
	"x6":  6,
	"x7":  7,
	"x8":  8,
	"x9":  9,
	"x10": 10,
	"x11": 11,
	"x12": 12,
	"x13": 13,
	"x14": 14,
	"x15": 15,
	"x16": 16,
	"x17": 17,
	"x18": 18,
	"x19": 19,
	"x20": 20,
	"x21": 21,
	"x22": 22,
	"x23": 23,
	"x24": 24,
	"x25": 25,
	"x26": 26,
	"x27": 27,
	"x28": 28,
	"x29": 29,
	"x30": 30,
	"x31": 31,

	"zero": 0,
	"ra":   1,
	"sp":   2,
	"gp":   3,
	"tp":   4,
	"t0":   5,
	"t1":   6,
	"t2":   7,
	"s0":   8,
	"s1":   9,
	"a0":   10,
	"a1":   11,
	"a2":   12,
	"a3":   13,
	"a4":   14,
	"a5":   15,
	"a6":   16,
	"a7":   17,
	"s2":   18,
	"s3":   19,
	"s4":   20,
	"s5":   21,
	"s6":   22,
	"s7":   23,
	"s8":   24,
	"s9":   25,
	"s10":  26,
	"s11":  27,
	"t3":   28,
	"t4":   29,
	"t5":   30,
	"t6":   31,

	"fp": 8,
}
