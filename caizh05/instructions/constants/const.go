package constants

import(
	"../base"
	"../../rt"
)

/*
把隐含在操作码中的常量值推入操作数栈顶
 */

type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

/*
aconst_null 指令把null引用推入操作数栈顶
 */
func (self *ACONST_NULL) Execute(frame *rt.Frame) {
	frame.OperandStack().PushRef(nil)
}
/*
dconst_0指令把double型0推入操作数栈顶
 */
func (self *DCONST_0) Execute(frame *rt.Frame) {
	frame.OperandStack().PushDouble(0.0)
}
/*
iconst_ml指令把int型-1推入操作数栈顶
 */
func (self *ICONST_M1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(-1)
}