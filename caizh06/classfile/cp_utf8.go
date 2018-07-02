package classfile
type ConstantUtf8Info struct {
	str string
}
/*
读取出[]byte 然后解码成go字符串
 */
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}
/*
简化版 不包含Null字符或补充字符
 */
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
