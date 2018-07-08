package lang

import (
	"../../../native"
	"../../../rt"
	"../../../rt/heap"
)
func init() {
	native.Register("java/lang/Class", "getPrimitiveClass",
		"(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.Register("java/lang/Class", "desiredAssertionStatus0",
		"(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}


/*
先从局部变量表中拿到类名，
这是个Java字符串，需要把它转成Go字符串。基本类型的类已经加
载到了方法区中，直接调用类加载器的LoadClass（）方法获取即可。
最后，把类对象引用推入操作数栈顶。
 */
func getPrimitiveClass(frame *rt.Frame) {
	nameObj := frame.LocalVars().GetRef(0)
	name := heap.GoString(nameObj)
	loader := frame.Method().Class().Loader()
	class := loader.LoadClass(name).JClass()
	frame.OperandStack().PushRef(class)
}
/*
首先从局部变量表中拿到this引用，这是一个类对象引用，通
过Extra（）方法可以获得与之对应的Class结构体指针。然后拿到类
名，转成Java字符串并推入操作数栈顶。注意这里需要的是
java.lang.Object这样的类名，而非java/lang/Object。Class结构体的
JavaName（）方法返回转换后的类名，
 */
func getName0(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.Extra().(*heap.Class)
	name := class.JavaName()
	nameObj := heap.JString(class.Loader(), name)
	frame.OperandStack().PushRef(nameObj)
}

func desiredAssertionStatus0(frame *rt.Frame) {
	frame.OperandStack().PushBoolean(false)
}