# rvasm

A simple RISC-V assembler. Supports the RV32i standard plus CSR and other
privileged instructions (`uret`, `sret`, `mret`, `wfi`). Also supports all
pseudoinstructions associated with the supported base instructions.

This is a toy assembler. It may be useful for small RISC-V simulator projects.

Caveats:

* Does not support the `fence` instruction.
* Does not support data sections/symbol lookups for load/store instructions.
* Not extensively tested (see `testdata` for tests).
* `li` pseudo-instruction does not support full 32-bit immediates.

rvasm also supports full RISC-V RV32gc disassembly via the
[deadsy/rvda](https://github.com/deadsy/rvda) package. Use the `-disas` flag to
disassemble.

# Usage

rvasm will print the instructions in hex format when given an assembly program.

```
$ rvasm prog.s
00500113
00c00193
ff718393
0023e233
0041f2b3
004282b3
...
```

Use the `-raw` flag to dump the bytes directly instead of using the hex
representation.

rvasm assumes the base address of the machine code to be 0, but you can change
this with the `-base` flag.

rvasm will also assemble input passed via stdin if there are no input files.
This can be quite useful for quickly generating machine code for individual
instructions:

```
$ echo "li t0, 42" | rvasm
02a00293
```
