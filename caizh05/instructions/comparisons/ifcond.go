package comparisons

import(
	"../base"
	"../../rt"
)
/*
if<cond>指令把操作数栈顶的int变量弹出，然后跟0进行比较，满足条件则跳转
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