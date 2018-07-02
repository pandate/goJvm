package stores

import "../../instructions/base"
import "../../rt"

// Store float into local variable
type FSTORE struct{ base.Index8Instruction }

func (self *FSTORE) Execute(frame *rt.Frame) {
	_fstore(frame, uint(self.Index))
}

type FSTORE_0 struct{ base.NoOperandsInstruction }

func (self *FSTORE_0) Execute(frame *rt.Frame) {
	_fstore(frame, 0)
}

type FSTORE_1 struct{ base.NoOperandsInstruction }

func (self *FSTORE_1) Execute(frame *rt.Frame) {
	_fstore(frame, 1)
}

type FSTORE_2 struct{ base.NoOperandsInstruction }

func (self *FSTORE_2) Execute(frame *rt.Frame) {
	_fstore(frame, 2)
}

type FSTORE_3 struct{ base.NoOperandsInstruction }

func (self *FSTORE_3) Execute(frame *rt.Frame) {
	_fstore(frame, 3)
}

func _fstore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}
