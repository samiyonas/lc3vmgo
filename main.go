package main

const MEMORY_MAX = 1 << 16;
var memory[MEMORY_MAX]uint16;

const (
    R_R0 = iota,
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
