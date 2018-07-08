package classfile

import "fmt"

type CodeAttribute struct {
	cp                 ConstantPool
	maxStack           uint16  //操作数栈最大深度
	maxLocals          uint16  //局部变量大小
	code                  []byte  //字节码
	exceptionTable     []*ExceptionTableEntry  //异常处理表
	attributes         []AttributeInfo
}
type ExceptionTableEntry struct {
	startPc       uint16
	endPc         uint16
	handlerPc     uint16
	catchType     uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {

	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
	fmt.Println("readInfo maxStack:",self.maxStack," code:",self.code)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}

func (self *CodeAttribute) MaxLocals() uint {
	return uint(self.maxLocals)
}
func (self *CodeAttribute) MaxStack() uint {
	return uint(self.maxStack)
}
func (self *CodeAttribute) Code() []byte {
	return self.code
}

func (self *ExceptionTableEntry) StartPc() uint16 {
	return self.startPc
}
func (self *ExceptionTableEntry) EndPc() uint16 {
	return self.endPc
}
func (self *ExceptionTableEntry) HandlerPc() uint16 {
	return self.handlerPc
}
func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}
