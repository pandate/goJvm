package base

/*
字节码读取结构体
 */
type BytecodeReader struct {
	code    []byte //字节码
	pc    int  //记录读取的指针
}
/*
重置字节码结构体属性
复用对象可以减少对象创建的开销
 */
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}
/*
读取一个字节的数据返回，指针加1
 */
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}
/*
读取一个无符号的字节码并转成有符号的int型返回
 */
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

/*
读取两个字节码并拼接成一个uint16的数据返回
 */
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

/*
读取两个字节码并拼接成一个uint16的数据并转换为一个有符号的int16型返回
 */
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

/*
读取四个字节码并拼接成一个int32的数据返回
 */
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
/*
tableswitch指令操作码的后面有0~3字节的padding，以保证defaultOffset在字节码中的地址是4的倍数。
 */
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}
/*
传入需要读取的int32字节数 返回读取的int32数组
 */
func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}
/*
返回当前的指针位置
 */
func (self *BytecodeReader) PC() int {
	return self.pc
}