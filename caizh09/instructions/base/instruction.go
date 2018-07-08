package base

import "../../rt"
/*
指令接口所有的指令都需要实现这两个方法
 */
type Instruction interface {
	//从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	//执行指令
	Execute(frame *rt.Frame)
}
/*
没有操作数的指令
实现一个没有操作数的指令作为具体操作数的组成部分，使用了组合的方式减少了后续没有操作数的指令的方法
 */
type NoOperandsInstruction struct {}
/*
没有操作数获取时不需要做任何操作
 */
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing to do
}


/*
表示跳转类型的指令抽象，Offset字段存放跳转偏移量。
 */
type BranchInstruction struct {
	Offset int  //跳转偏移量
}
/*
从字节码中读取一个uint16整数，转成int后赋给Offset字段。
 */
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出
 */
type Index8Instruction struct {
	Index uint //局部变量表索引
}
/*
读取字节码获取一个索引地址后赋值给局部变量索引
 */
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}
/*
访问运行时常量池，常量池索引由两字节操作数给出
 */
type Index16Instruction struct {
	Index uint
}
/*
读取字节码获取一个常量池地址后赋值给局部变量索引
 */
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint16())
}