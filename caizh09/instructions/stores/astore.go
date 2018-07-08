package stores

import "../../instructions/base"
import "../../rt"

/*
astore 将栈顶引用型数值存入指定本地变量
astore_0 将栈顶引用型数值存入第一个本地变量
astore_1 将栈顶引用型数值存入第二个本地变量
astore_2 将栈顶引用型数值存入第三个本地变量
astore_3 将栈顶引用型数值存入第四个本地变量
*/
type ASTORE struct{ base.Index8Instruction }
type ASTORE_0 struct{ base.NoOperandsInstruction }
type ASTORE_1 struct{ base.NoOperandsInstruction }
type ASTORE_2 struct{ base.NoOperandsInstruction }
type ASTORE_3 struct{ base.NoOperandsInstruction }

func (self *ASTORE) Execute(frame *rt.Frame) {
	_astore(frame, uint(self.Index))
}


func (self *ASTORE_0) Execute(frame *rt.Frame) {
	_astore(frame, 0)
}


func (self *ASTORE_1) Execute(frame *rt.Frame) {
	_astore(frame, 1)
}


func (self *ASTORE_2) Execute(frame *rt.Frame) {
	_astore(frame, 2)
}


func (self *ASTORE_3) Execute(frame *rt.Frame) {
	_astore(frame, 3)
}

func _astore(frame *rt.Frame, index uint) {
	ref := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, ref)
}
