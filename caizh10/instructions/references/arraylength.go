package references

import "../../instructions/base"
import "../../rt"

// arraylength指令用于获取数组长度。
type ARRAY_LENGTH struct{ base.NoOperandsInstruction }

/*
arraylength指令只需要一个操作数，即从操作数栈顶弹出的数
组引用。Execute（）方法把数组长度推入操作数栈顶，
 */
func (self *ARRAY_LENGTH) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
