package comparisons

import(
	"../base"
	"../../rt"
)
/*
比较指令可以分为两类：一类将比较结果推入操作数栈顶，一类根据比较结果跳转。比较指令是编译器实现
if-else、for、while等语句的基石，共有19条,lcmp指令用于比较long变量
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