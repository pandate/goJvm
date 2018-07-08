package math

import "../../instructions/base"
import "../../rt"

/*
dmul 将栈顶两double型数值相乘并将结果压入栈顶
fmul 将栈顶两float型数值相乘并将结果压入栈顶
imul 将栈顶两int型数值相乘并将结果压入栈顶
lmul 将栈顶两long型数值相乘并将结果压入栈顶
*/
type DMUL struct{ base.NoOperandsInstruction }
type FMUL struct{ base.NoOperandsInstruction }
type IMUL struct{ base.NoOperandsInstruction }
type LMUL struct{ base.NoOperandsInstruction }

func (self *DMUL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

func (self *FMUL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

func (self *IMUL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

func (self *LMUL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}
