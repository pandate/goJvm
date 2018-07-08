package conversions

import "../../instructions/base"
import "../../rt"

/*
f2d 从栈中取出一个float类型数据转成double类型
f2i 从栈中取出一个float类型数据转成int类型
f2l 从栈中取出一个float类型数据转成long类型
*/
type F2D struct{ base.NoOperandsInstruction }
type F2I struct{ base.NoOperandsInstruction }
type F2L struct{ base.NoOperandsInstruction }

func (self *F2D) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	d := float64(f)
	stack.PushDouble(d)
}

func (self *F2I) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

func (self *F2L) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	l := int64(f)
	stack.PushLong(l)
}
