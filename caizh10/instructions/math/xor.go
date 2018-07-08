package math

import "../../instructions/base"
import "../../rt"

/*
按位异或指令
ixor int类型按位异或
lxor long类型按位异或
*/
type IXOR struct{ base.NoOperandsInstruction }
type LXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// Boolean XOR long

func (self *LXOR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
