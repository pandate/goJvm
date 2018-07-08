package loads
import(
	"../base"
	"../../rt"
)
/*
加载指令从局部变量表获取变量，然后推入操作数栈顶,int变量操作
iload 根据局部变量表索引从局部变量中获取一个Int数据并推入栈顶
iload_0 从局部变量中获取第0个位置的Int数据并推入栈顶
iload_1 从局部变量中获取第1个位置的Int数据并推入栈顶
iload_2 从局部变量中获取第2个位置的Int数据并推入栈顶
iload_3 从局部变量中获取第3个位置的Int数据并推入栈顶
 */

type ILOAD struct{ base.Index8Instruction }
type ILOAD_0 struct{ base.NoOperandsInstruction }
type ILOAD_1 struct{ base.NoOperandsInstruction }
type ILOAD_2 struct{ base.NoOperandsInstruction }
type ILOAD_3 struct{ base.NoOperandsInstruction }


func _iload(frame *rt.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}


func (self *ILOAD) Execute(frame *rt.Frame) {
	_iload(frame, uint(self.Index))
}

func (self *ILOAD_0) Execute(frame *rt.Frame) {
	_iload(frame, 0)
}

func (self *ILOAD_1) Execute(frame *rt.Frame) {
	_iload(frame, 1)
}

func (self *ILOAD_2) Execute(frame *rt.Frame) {
	_iload(frame, 2)
}

func (self *ILOAD_3) Execute(frame *rt.Frame) {
	_iload(frame, 3)
}