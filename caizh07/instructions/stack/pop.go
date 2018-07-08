package stack

import(
	"../base"
	"../../rt"
)
/*
栈指令直接对操作数栈进行操作，共9条：pop和pop2指令将栈顶变量弹出，dup系列指令复制栈顶变量，swap指令交换栈顶的两个变量
 */

type POP struct{ base.NoOperandsInstruction }
type POP2 struct{ base.NoOperandsInstruction }
/*
弹出int等占用一个操作数栈位置的变量
 */
func (self *POP) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}
/*
弹出double等占用两个操作数栈位置的变量
 */
func (self *POP2) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}