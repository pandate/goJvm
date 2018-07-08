package comparisons
import(
	"../base"
	"../../rt"
)
/*
dcmpg
比较栈顶两double型数值大小，并且结果（1，0，-1）进栈；当其中一个数值为NaN时，1进栈
dcmpl
比较栈顶两double型数值大小，并且结果（1，0，-1）进栈；当其中一个数值为NaN时，-1进栈
 */
type DCMPG struct{ base.NoOperandsInstruction }
type DCMPL struct{ base.NoOperandsInstruction }

/*
抽象比较函数减少重复代码，从
 */
func _dcmp(frame *rt.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (self *DCMPG) Execute(frame *rt.Frame) {
	_dcmp(frame, true)
}
func (self *DCMPL) Execute(frame *rt.Frame) {
	_dcmp(frame, false)
}