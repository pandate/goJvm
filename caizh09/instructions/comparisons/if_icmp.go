package comparisons

import(
	"../base"
	"../../rt"
)


/*
if_icmpeq比较栈顶两int型数值大小，当结果等于0时跳转
if_icmpne比较栈顶两int型数值大小，当结果不等于0时跳转
if_icmplt比较栈顶两int型数值大小，当结果小于0时跳转
if_icmple比较栈顶两int型数值大小，当结果小于等于0时跳转
if_icmpgt比较栈顶两int型数值大小，当结果大于0时跳转
if_icmpge比较栈顶两int型数值大小，当结果大于等于0时跳转
*/
type IF_ICMPEQ struct{ base.BranchInstruction }
type IF_ICMPNE struct{ base.BranchInstruction }
type IF_ICMPLT struct{ base.BranchInstruction }
type IF_ICMPLE struct{ base.BranchInstruction }
type IF_ICMPGT struct{ base.BranchInstruction }
type IF_ICMPGE struct{ base.BranchInstruction }

func (self *IF_ICMPEQ) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 == val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 != val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 < val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 <= val2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 > val2 {
		base.Branch(frame, self.Offset)
	}
}


func (self *IF_ICMPGE) Execute(frame *rt.Frame) {
	if val1, val2 := _icmpPop(frame); val1 >= val2 {
		base.Branch(frame, self.Offset)
	}
}
/*
弹出栈顶的两个int数据
 */
func _icmpPop(frame *rt.Frame) (val1, val2 int32) {
	stack := frame.OperandStack()
	val2 = stack.PopInt()
	val1 = stack.PopInt()
	return
}