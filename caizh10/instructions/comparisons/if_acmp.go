package comparisons

import(
	"../base"
	"../../rt"
)
/*
if_acmpeq 比较栈顶两引用型数值，当结果相等时跳转
if_acmpne 比较栈顶两引用型数值，当结果不相等时跳转
 */
type IF_ACMPEQ struct{ base.BranchInstruction }
type IF_ACMPNE struct{ base.BranchInstruction }

func (self *IF_ACMPEQ) Execute(frame *rt.Frame) {
	if _acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) Execute(frame *rt.Frame) {
	if !_acmp(frame) {
		base.Branch(frame, self.Offset)
	}
}
/*
判断栈顶引用是否相等
 */
func _acmp(frame *rt.Frame) bool {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	return ref1 == ref2 // todo
}