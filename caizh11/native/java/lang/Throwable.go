package lang

import (
	"../../../native"
	"../../../rt"
	"../../../rt/heap"
)


/*
从字节码来看，如果没有异常抛出，则会直接goto到return指
令，方法正常返回。那么如果有异常抛出，goto和return之间的指令
是如何执行的呢？答案是查找方法的异常处理表。异常处理表是
Code属性的一部分，它记录了方法是否有能力处理某种异常。
异常处理表的每一项都包含3个信息：处理哪部分代码抛出的
异常、哪类异常，以及异常处理代码在哪里。具体来说，start_pc和
end_pc可以锁定一部分字节码，这部分字节码对应某个可能抛出异
常的try{}代码块。catch_type是个索引，通过它可以从运行时常量池
中查到一个类符号引用，解析后的类是个异常类，假定这个类是
X。如果位于start_pc和end_pc之间的指令抛出异常x，且x是X（或者
X的子类）的实例，handler_pc就指出负责异常处理的catch{}块在哪
里。
当tryItOut（）方法通过athrow指令抛出TestExc异常时，Java虚拟
机首先会查找tryItOut（）方法的异常处理表，看它能否处理该异常。
如果能，则跳转到相应的字节码开始异常处理。假设tryItOut（）方法
无法处理异常，Java虚拟机会进一步查看它的调用者，也就是
catchOne（）方法的异常处理表。catchOne（）方法刚好可以处理
TestExc异常，使catch{}块得以执行。
假设catchOne（）方法也无法处理TestExc异常，Java虚拟机会继
续查找catchOne（）的调用者的异常处理表。这个过程会一直继续下
去，直到找到某个异常处理项，或者到达Java虚拟机栈的底部。
 */
func init() {
	native.Register("java/lang/Throwable", "fillInStackTrace",
		"(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName string  //文件名
	className string  //类型
	methodName string  //方法名
	lineNumber int //帧正在执行哪行代码
}

func fillInStackTrace(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushRef(this)
	stes := createStackTraceElements(this, frame.Thread())
	this.SetExtra(stes)
}
/*
由于栈顶两帧正在执行
fillInStackTrace（int）和fillInStackTrace（）方法，所以需要跳过这两
帧。这两帧下面的几帧正在执行异常类的构造函数，所以也要跳
过，具体要跳过多少帧数则要看异常类的继承层次。
distanceToObject（）函数计算所需跳过的帧数
计算好需要跳过的帧之后，调用Thread结构体的GetFrames（）
方法拿到完整的Java虚拟机栈，然后reslice一下就是真正需要的帧。
 */
func createStackTraceElements(tObj *heap.Object, thread *rt.Thread)[]*StackTraceElement {
	skip := distanceToObject(tObj.Class()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
	stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.SuperClass(); c != nil; c = c.SuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rt.Frame) *StackTraceElement {
	method := frame.Method()
	class := method.Class()
	return &StackTraceElement{
		fileName: class.SourceFile(),
		className: class.JavaName(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

