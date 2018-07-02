package references

import "../../instructions/base"
import "../../rt"
import "../../rt/heap"

/*
instanceof指令判断对象是否是某个类的实例（或者对象的类
是否实现了某个接口），并把结果推入操作数栈。
 */
type INSTANCE_OF struct{ base.Index16Instruction }


/*
先弹出对象引用，如果是null，则把0推入操作数栈。用Java代码解释就是，
如果引用obj是null的话，不管ClassYYY是哪种类型，判断都是false
 */
func (self *INSTANCE_OF) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
