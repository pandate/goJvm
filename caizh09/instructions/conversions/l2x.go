package conversions

import "../../instructions/base"
import "../../rt"

/*
l2d 从栈中取出一个long类型数据转成double类型
l2f 从栈中取出一个long类型数据转成float类型
l2i 从栈中取出一个long类型数据转成int类型
*/
type L2D struct{ base.NoOperandsInstruction }
type L2F struct{ base.NoOperandsInstruction }
type L2I struct{ base.NoOperandsInstruction }

func (self *L2D) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

func (self *L2F) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

func (self *L2I) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
