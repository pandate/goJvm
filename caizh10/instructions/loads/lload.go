package loads

import "../../instructions/base"
import "../../rt"

/*
lload 根据局部变量表索引从局部变量中获取一个long数据并推入栈顶
lload_0 从局部变量中获取第0个位置的long数据并推入栈顶
lload_1 从局部变量中获取第1个位置的long数据并推入栈顶
lload_2 从局部变量中获取第2个位置的long数据并推入栈顶
lload_3 从局部变量中获取第3个位置的long数据并推入栈顶
*/
type LLOAD struct{ base.Index8Instruction }
type LLOAD_0 struct{ base.NoOperandsInstruction }
type LLOAD_1 struct{ base.NoOperandsInstruction }
type LLOAD_2 struct{ base.NoOperandsInstruction }
type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD) Execute(frame *rt.Frame) {
	_lload(frame, self.Index)
}

func (self *LLOAD_0) Execute(frame *rt.Frame) {
	_lload(frame, 0)
}

func (self *LLOAD_1) Execute(frame *rt.Frame) {
	_lload(frame, 1)
}

func (self *LLOAD_2) Execute(frame *rt.Frame) {
	_lload(frame, 2)
}

func (self *LLOAD_3) Execute(frame *rt.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rt.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
