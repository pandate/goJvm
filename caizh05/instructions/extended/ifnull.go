package extended
import(
	"../base"
	"../../rt"
)
/*
根据引用是否是null进行跳转，ifnull和ifnonnull指令把栈顶的引用弹出。
 */
type IFNULL struct{ base.BranchInstruction } // Branch if reference is null
type IFNONNULL struct{ base.BranchInstruction } // Branch if reference not null

func (self *IFNULL) Execute(frame *rt.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}