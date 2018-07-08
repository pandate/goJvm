package references

import (
	"../../instructions/base"
	"../../rt"
	"../../rt/heap"
)


/*
new指令专门用来创建类实例。
 */
type NEW struct{ base.Index16Instruction }


/*
new指令的操作数是一个uint16索引，来自字节码。通过这个索引，可以从当前类的运
行时常量池中找到一个类符号引用。解析这个类符号引用，拿到类数据，然后创建对象，
并把对象引用推入栈顶，new指令的工作就完成了。
 */
func (self *NEW) Execute(frame *rt.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
