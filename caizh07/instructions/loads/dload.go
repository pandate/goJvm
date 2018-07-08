package loads

import "../../instructions/base"
import "../../rt"

/*
dload 根据局部变量表索引从局部变量中获取一个double数据并推入栈顶
dload_0 从局部变量中获取第0个位置的double数据并推入栈顶
dload_1 从局部变量中获取第1个位置的double数据并推入栈顶
dload_2 从局部变量中获取第2个位置的double数据并推入栈顶
dload_3 从局部变量中获取第3个位置的double数据并推入栈顶
*/
type DLOAD struct{ base.Index8Instruction }
type DLOAD_0 struct{ base.NoOperandsInstruction }
type DLOAD_1 struct{ base.NoOperandsInstruction }
type DLOAD_2 struct{ base.NoOperandsInstruction }
type DLOAD_3 struct{ base.NoOperandsInstruction }
func (self *DLOAD) Execute(frame *rt.Frame) {
	_dload(frame, self.Index)
}

func (self *DLOAD_0) Execute(frame *rt.Frame) {
	_dload(frame, 0)
}

func (self *DLOAD_1) Execute(frame *rt.Frame) {
	_dload(frame, 1)
}

func (self *DLOAD_2) Execute(frame *rt.Frame) {
	_dload(frame, 2)
}

func (self *DLOAD_3) Execute(frame *rt.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rt.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
