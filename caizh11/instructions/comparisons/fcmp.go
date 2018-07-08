package comparisons
import(
	"../base"
	"../../rt"
)
/*
fcmpg比较栈顶两float型数值大小，并且结果（1，0，-1）进栈；当其中一个数值为NaN时， 1进栈
fcmpl比较栈顶两float型数值大小，并且结果（1，0，-1）进栈；当其中一个数值为NaN时，-1进栈
 */
type FCMPG struct{ base.NoOperandsInstruction }
type FCMPL struct{ base.NoOperandsInstruction }

func _fcmp(frame *rt.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

func (self *FCMPG) Execute(frame *rt.Frame) {
	_fcmp(frame, true)
}
func (self *FCMPL) Execute(frame *rt.Frame) {
	_fcmp(frame, false)
}