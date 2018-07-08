package references

import "../../instructions/base"
import "../../rt"
import "../../rt/heap"

/*
检验类型转换指令
heckcast指令和instanceof指令很像，区别在于：instanceof指令会改变操作数栈（弹出对象
引用，推入判断结果）；checkcast则不改变操作数栈（如果判断失败，直接抛出
ClassCastException异常）。
 */
type CHECK_CAST struct{ base.Index16Instruction }


/*
先从操作数栈中弹出对象引用，再推回去，这样就不会改变操作数栈的状态。如果引用是null
，则指令执行结束。也就是说，null引用可以转换成任何类型，否则解析类符号引用，判断对
象是否是类的实例。如果是的话，指令执行结束，否则抛出ClassCastException。
 */
func (self *CHECK_CAST) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
