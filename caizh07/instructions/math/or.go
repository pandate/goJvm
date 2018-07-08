package math

import "../../instructions/base"
import "../../rt"

/*
ior	将栈顶两int型数值作“按位或”并将结果压入栈顶
lor 将栈顶两long型数值作“按位或”并将结果压入栈顶
*/
type IOR struct{ base.NoOperandsInstruction }
type LOR struct{ base.NoOperandsInstruction }

func (self *IOR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

func (self *LOR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}
