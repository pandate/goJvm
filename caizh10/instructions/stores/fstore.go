package stores

import "../../instructions/base"
import "../../rt"

/*
fstore 将栈顶float型数值存入指定本地变量
fstore_0 将栈顶float型数值存入本地第一个变量
fstore_1 将栈顶float型数值存入本地第二个变量
fstore_2 将栈顶float型数值存入本地第三个变量
fstore_3 将栈顶float型数值存入本地第四个变量
*/
type FSTORE struct{ base.Index8Instruction }
type FSTORE_0 struct{ base.NoOperandsInstruction }
type FSTORE_1 struct{ base.NoOperandsInstruction }
type FSTORE_2 struct{ base.NoOperandsInstruction }
type FSTORE_3 struct{ base.NoOperandsInstruction }

func (self *FSTORE) Execute(frame *rt.Frame) {
	_fstore(frame, uint(self.Index))
}


func (self *FSTORE_0) Execute(frame *rt.Frame) {
	_fstore(frame, 0)
}


func (self *FSTORE_1) Execute(frame *rt.Frame) {
	_fstore(frame, 1)
}


func (self *FSTORE_2) Execute(frame *rt.Frame) {
	_fstore(frame, 2)
}


func (self *FSTORE_3) Execute(frame *rt.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
