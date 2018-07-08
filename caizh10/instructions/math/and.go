package math
import(
	"../base"
	"../../rt"
)
/*
布尔运算指令与
iland 将栈顶两int型数值作“按位与”并将结果压入栈顶
land  将栈顶两long型数值作“按位与”并将结果压入栈顶
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

func (self *LAND) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}