package comparisons

import(
	"../base"
	"../../rt"
)
/*
比较栈顶两long型数值大小，并且结果（1，0，-1）进栈
 */

type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}