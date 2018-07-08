package stores

import "../../instructions/base"
import "../../rt"

/*
dstore 将栈顶double型数值存入指定本地变量
dstore_0 将栈顶double型数值存入本地第一个变量
dstore_1 将栈顶double型数值存入本地第二个变量
dstore_2 将栈顶double型数值存入本地第三个变量
dstore_3 将栈顶double型数值存入本地第四个变量
*/
type DSTORE struct{ base.Index8Instruction }
type DSTORE_0 struct{ base.NoOperandsInstruction }
type DSTORE_1 struct{ base.NoOperandsInstruction }
type DSTORE_2 struct{ base.NoOperandsInstruction }
type DSTORE_3 struct{ base.NoOperandsInstruction }

func (self *DSTORE) Execute(frame *rt.Frame) {
	_dstore(frame, uint(self.Index))
}


func (self *DSTORE_0) Execute(frame *rt.Frame) {
	_dstore(frame, 0)
}


func (self *DSTORE_1) Execute(frame *rt.Frame) {
	_dstore(frame, 1)
}


func (self *DSTORE_2) Execute(frame *rt.Frame) {
	_dstore(frame, 2)
}


func (self *DSTORE_3) Execute(frame *rt.Frame) {
	_dstore(frame, 3)
}

func _dstore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}
