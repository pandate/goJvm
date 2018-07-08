package comparisons

import(
	"../base"
	"../../rt"
)
/*
ifeq指令把操作数栈顶的int变量弹出，等于0时跳转
ifne指令把操作数栈顶的int变量弹出，不等于0时跳转
iflt指令把操作数栈顶的int变量弹出，小于0时跳转
ifle指令把操作数栈顶的int变量弹出，小于等于0时跳转
ifgt指令把操作数栈顶的int变量弹出，大于0时跳转
ifge指令把操作数栈顶的int变量弹出，大于等于0时跳转
 */
type IFEQ struct{ base.BranchInstruction }
type IFNE struct{ base.BranchInstruction }
type IFLT struct{ base.BranchInstruction }
type IFLE struct{ base.BranchInstruction }
type IFGT struct{ base.BranchInstruction }
type IFGE struct{ base.BranchInstruction }

func (self *IFEQ) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) Execute(frame *rt.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, self.Offset)
	}
}