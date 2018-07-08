package stores

import "../../instructions/base"
import "../../rt"

/*
istore 将栈顶int型数值存入指定本地变量
istore_0 将栈顶int型数值存入本地第一个变量
istore_1 将栈顶int型数值存入本地第二个变量
istore_2 将栈顶int型数值存入本地第三个变量
istore_3 将栈顶int型数值存入本地第四个变量
*/
type ISTORE struct{ base.Index8Instruction }
type ISTORE_0 struct{ base.NoOperandsInstruction }
type ISTORE_1 struct{ base.NoOperandsInstruction }
type ISTORE_2 struct{ base.NoOperandsInstruction }
type ISTORE_3 struct{ base.NoOperandsInstruction }

func (self *ISTORE) Execute(frame *rt.Frame) {
	_istore(frame, uint(self.Index))
}


func (self *ISTORE_0) Execute(frame *rt.Frame) {
	_istore(frame, 0)
}


func (self *ISTORE_1) Execute(frame *rt.Frame) {
	_istore(frame, 1)
}


func (self *ISTORE_2) Execute(frame *rt.Frame) {
	_istore(frame, 2)
}


func (self *ISTORE_3) Execute(frame *rt.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}
