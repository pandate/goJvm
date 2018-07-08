package math

import "../../instructions/base"
import "../../rt"

/*
dneg 将栈顶double型数值取负并将结果压入栈顶
fneg 将栈顶float型数值取负并将结果压入栈顶
ineg 将栈顶int型数值取负并将结果压入栈顶
lneg 将栈顶long型数值取负并将结果压入栈顶
*/
type DNEG struct{ base.NoOperandsInstruction }
type FNEG struct{ base.NoOperandsInstruction }
type INEG struct{ base.NoOperandsInstruction }
type LNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

func (self *FNEG) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

func (self *INEG) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

func (self *LNEG) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
