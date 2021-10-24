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
