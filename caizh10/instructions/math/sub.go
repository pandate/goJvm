package math

import "../../instructions/base"
import "../../rt"

/*
减法指令
dsub double相减
fsub float相减
isub int相减
lsub long相减
 */
type DSUB struct{ base.NoOperandsInstruction }
type FSUB struct{ base.NoOperandsInstruction }
type ISUB struct{ base.NoOperandsInstruction }
type LSUB struct{ base.NoOperandsInstruction }

func (self *DSUB) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

// Subtract float

func (self *FSUB) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

// Subtract int

func (self *ISUB) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

// Subtract long

func (self *LSUB) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
