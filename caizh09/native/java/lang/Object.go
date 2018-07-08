package lang

import (
	"../../../native"
	"../../../rt"
	"unsafe"
)
func init() {
	native.Register("java/lang/Object", "getClass", "()Ljava/lang/Class;", getClass)
	native.Register("java/lang/Object", "hashCode", "()I", hashCode)
	native.Register("java/lang/Object", "clone", "()Ljava/lang/Object;", clone)
}

/*
首先，从局部变量表中拿到this引用。GetThis（）方法其实就是调用
GetRef（0），不过为了提高代码的可读性，给LocalVars结构体添加了
这个方法。有了this引用后，通过Class（）方法拿到它的Class结构体
指针，进而又通过JClass（）方法拿到它的类对象。最后，把类对象推
入操作数栈顶。
 */
func getClass(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Class().JClass()
	frame.OperandStack().PushRef(class)
}

func hashCode(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func clone(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	cloneable := this.Class().Loader().LoadClass("java/lang/Cloneable")
	if !this.Class().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.OperandStack().PushRef(this.Clone())
}
