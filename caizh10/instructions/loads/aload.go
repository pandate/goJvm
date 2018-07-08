package loads

import "../../instructions/base"
import "../../rt"

/*
aload 根据局部变量表索引从局部变量中获取一个引用数据并推入栈顶
aload_0 从局部变量中获取第0个位置的引用数据并推入栈顶
aload_1 从局部变量中获取第1个位置的引用数据并推入栈顶
aload_2 从局部变量中获取第2个位置的引用数据并推入栈顶
aload_3 从局部变量中获取第3个位置的引用数据并推入栈顶
*/
type ALOAD struct{ base.Index8Instruction }

func (self *ALOAD) Execute(frame *rt.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct{ base.NoOperandsInstruction }

func (self *ALOAD_0) Execute(frame *rt.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct{ base.NoOperandsInstruction }

func (self *ALOAD_1) Execute(frame *rt.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct{ base.NoOperandsInstruction }

func (self *ALOAD_2) Execute(frame *rt.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct{ base.NoOperandsInstruction }

func (self *ALOAD_3) Execute(frame *rt.Frame) {
	_aload(frame, 3)
}

func _aload(frame *rt.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
