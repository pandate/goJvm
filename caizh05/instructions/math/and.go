package math
import(
	"../base"
	"../../rt"
)
/*
布尔运算指令与
 */
type IAND struct{ base.NoOperandsInstruction }
type LAND struct{ base.NoOperandsInstruction }

func (self *IAND) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}