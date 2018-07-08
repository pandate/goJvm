package loads

import "../../instructions/base"
import "../../rt"

/*
fload 根据局部变量表索引从局部变量中获取一个float数据并推入栈顶
fload_0 从局部变量中获取第0个位置的float数据并推入栈顶
fload_1 从局部变量中获取第1个位置的float数据并推入栈顶
fload_2 从局部变量中获取第2个位置的float数据并推入栈顶
fload_3 从局部变量中获取第3个位置的float数据并推入栈顶
*/
type FLOAD struct{ base.Index8Instruction }
type FLOAD_0 struct{ base.NoOperandsInstruction }
type FLOAD_1 struct{ base.NoOperandsInstruction }
type FLOAD_2 struct{ base.NoOperandsInstruction }
type FLOAD_3 struct{ base.NoOperandsInstruction }

func (self *FLOAD) Execute(frame *rt.Frame) {
	_fload(frame, self.Index)
}

func (self *FLOAD_0) Execute(frame *rt.Frame) {
	_fload(frame, 0)
}

func (self *FLOAD_1) Execute(frame *rt.Frame) {
	_fload(frame, 1)
}

func (self *FLOAD_2) Execute(frame *rt.Frame) {
	_fload(frame, 2)
}

func (self *FLOAD_3) Execute(frame *rt.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rt.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
