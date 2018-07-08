package conversions

import(
	"../base"
	"../../rt"
)
/*
类型转换指令对应java语言中的基本类型强制转换操作。类型转换指令共有15条
这里实现把double变量强制转换成其他类型
d2f 从栈中取出一个double类型数据转成float类型
d2i 从栈中取出一个double类型数据转成int类型
d2l 从栈中取出一个double类型数据转成long类型
 */

type D2F struct{ base.NoOperandsInstruction }
type D2I struct{ base.NoOperandsInstruction }
type D2L struct{ base.NoOperandsInstruction }

func (self *D2F) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}


func (self *D2I) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (self *D2L) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}


