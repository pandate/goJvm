package references

import "../../instructions/base"
import "../../rt"
import (
	"../../rt/heap"
)

/*
调用超类构造方法，实例初始化方法，私有方法
因为对象是需要初始化的，所以每个类都至少有一个构造函数。即使用户自己不定义，
编译器也会自动生成一个默认构造函数。在创建类实例时，编译器会在new指令的后面
加入invokespecial指令来调用构造函数初始化对象。
*/
type INVOKE_SPECIAL struct{ base.Index16Instruction }


func (self *INVOKE_SPECIAL) Execute(frame *rt.Frame) {
	//先拿到当前类、当前常量池、方法符号引用，然后解析符号引
	//用，拿到解析后的类和方法。
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//从操作数栈中弹出this引用，如果该引用是null，抛出
	//NullPointerException异常。注意，在传递参数之前，不能破坏操作
	//数栈的状态
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount()-1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	//判断确保protected方法只能被声明该方法的类或子类
	//调用。如果违反这一规定，则抛出IllegalAccessError异常。接着往下
	//看
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	//上面这段代码比较难懂，把它翻译成更容易理解的语言：如果
	//调用的中超类中的函数，但不是构造函数，且当前类的
	//ACC_SUPER标志被设置，需要一个额外的过程查找最终要调用的
	//方法；否则前面从方法符号引用中解析出来的方法就是要调用的方
	//法。
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}
	//如果查找过程失败，或者找到的方法是抽象的，抛出
	//AbstractMethodError异常。最后，如果一切正常，就调用方法。
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}
