package stores
import(
	"../base"
	"../../rt"
)
/*
存储指令把变量从操作数栈顶弹出，然后存入变量表
 */

type LSTORE struct{ base.Index8Instruction }
type LSTORE_0 struct{ base.NoOperandsInstruction }
type LSTORE_1 struct{ base.NoOperandsInstruction }
type LSTORE_2 struct{ base.NoOperandsInstruction }
type LSTORE_3 struct{ base.NoOperandsInstruction }

func _lstore(frame *rt.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (self *LSTORE) Execute(frame *rt.Frame) {
	_lstore(frame, uint(self.Index))
}

func (self *LSTORE_0) Execute(frame *rt.Frame) {
	_lstore(frame, 0)
}

func (self *LSTORE_1) Execute(frame *rt.Frame) {
	_lstore(frame, 1)
}

func (self *LSTORE_2) Execute(frame *rt.Frame) {
	_lstore(frame, 2)
}

func (self *LSTORE_3) Execute(frame *rt.Frame) {
	_lstore(frame, 3)
}