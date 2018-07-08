package heap

import (
	"../../classfile"
)

type Method struct {
	ClassMember
	maxStack      uint  //操作数栈
	maxLocals     uint  //局部变量表大小
	code          []byte  //方法字节码
	argSlotCount  uint //参数个数
	exceptionTable ExceptionTable //异常信息表
	lineNumberTable *classfile.LineNumberTableAttribute
}

func newMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	//先计算argSlotCount字段，如果是本地
	//方法，则注入字节码和其他信息。
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}


func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {

	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(),
			self.class.constantPool)
	}
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}
/*
FindExceptionHandler（）方法调用
ExceptionTable.findExceptionHandler（）方法搜索异常处理表，如果
能够找到对应的异常处理项，则返回它的handlerPc字段，否则返
回–1。
 */
func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}


func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}
func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}
func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}
func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}
func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

func (self *Method) ArgSlotCount() uint {
	//fmt.Println("参数个数：",self.argSlotCount)
	return self.argSlotCount
}


func (self *Method) MaxStack() uint {
	return self.maxStack
}
func (self *Method) MaxLocals() uint {
	return self.maxLocals
}
func (self *Method) Code() []byte {
	return self.code
}
/*
计算变量数量占用的操作数
 */
func (self *Method) calcArgSlotCount(param []string) {
	for _, paramType := range param {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

/*
本地方法在class文件中没有Code属性，所以需要给maxStack和
maxLocals字段赋值。本地方法帧的操作数栈至少要能容纳返回值，
为了简化代码，暂时给maxStack字段赋值为4。因为本地方法帧的
局部变量表只用来存放参数值，所以把argSlotCount赋给maxLocals
字段刚好。至于code字段，也就是本地方法的字节码，第一条指令
都是0xFE，第二条指令则根据函数的返回值选择相应的返回指令。
 */
func (self *Method) injectCodeAttribute(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V': self.code = []byte{0xfe, 0xb1} // return
	case 'D': self.code = []byte{0xfe, 0xaf} // dreturn
	case 'F': self.code = []byte{0xfe, 0xae} // freturn
	case 'J': self.code = []byte{0xfe, 0xad} // lreturn
	case 'L', '[': self.code = []byte{0xfe, 0xb0} // areturn
	default: self.code = []byte{0xfe, 0xac} // ireturn
	}
}
