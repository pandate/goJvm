package math

import "../../instructions/base"
import "../../rt"

/*
ddiv 将栈顶两double型数值相除并将结果压入栈顶
fdiv 将栈顶两float型数值相除并将结果压入栈顶
idiv 将栈顶两int型数值相除并将结果压入栈顶
ldiv 将栈顶两long型数值相除并将结果压入栈顶
*/
type DDIV struct{ base.NoOperandsInstruction }
type FDIV struct{ base.NoOperandsInstruction }
type IDIV struct{ base.NoOperandsInstruction }
type LDIV struct{ base.NoOperandsInstruction }

func (self *DDIV) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

func (self *FDIV) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

func (self *IDIV) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushInt(result)
}

func (self *LDIV) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := v1 / v2
	stack.PushLong(result)
}
