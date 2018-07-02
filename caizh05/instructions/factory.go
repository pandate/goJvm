package instructions

import (
	"./base"
	"./constants"
	"fmt"
)

var (
	nop = &constants.NOP{}
	aconst_null = &constants.ACONST_NULL{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00: return nop
	case 0x01: return aconst_null
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}