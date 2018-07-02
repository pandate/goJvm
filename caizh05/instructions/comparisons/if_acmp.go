package comparisons

import(
	"../base"
	"../../rt"
)
/*
if_acmpeq和if_acmpne指令把栈顶的两个引用弹出，根据引用是否相同进行跳转。
 */
type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}