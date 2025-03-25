package main

const MEMORY_MAX = 1 << 16;
var memory[MEMORY_MAX]uint16;

const (
    R_R0 = iota
    R_R1
    R_R2
    R_R3
    R_R4
    R_R5
    R_R6
    R_R7
    R_PC /* program counter */
    R_COND /* condition flag */
    R_COUNT
)

var reg[R_COUNT]uint16;

const (
    OP_BR = iota /* branch */
    OP_ADD /* add */
    OP_LD /* load */
    OP_ST /* store */
    OP_JSR /* jump register */
    OP_AND /* bitwise and */
    OP_LDR /* load register */
    OP_STR /* store register */
    OP_RTI /* unused */
    OP_NOT /* bitwise not */
    OP_LDI /* load indirect */
    OP_STI /* store indirect */
    OP_JMP /* jump */
    OP_RES /* reserved (unused) */
    OP_LEA /* load effective address */
    OP_TRAP /* execute trap */
)

const (
    FL_POS = 1 << 0
    FL_ZRO = 1 << 1
    FL_NEG = 1 << 2
)
