package base

import "../../rt"

type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rt.Frame)
}
/*
没有操作数的指令
 */
type NoOperandsInstruction struct {}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}
/*
跳转指令
 */
type BranchInstruction struct {
	Offset int  //跳转偏移量
}

func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}
/*
存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
 */
type Index8Instruction struct {
	Index uint //局部变量表索引
}

func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}
/*
访问运行时常量池，常量池索引由两字节操作数给出
 */
type Index16Instruction struct {
	Index uint
}

func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}