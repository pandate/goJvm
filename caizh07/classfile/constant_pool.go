package classfile

import "fmt"

/*
常量池实际上也是一个表，有三个点需要注意
1、表头给出的常量池大小比实际大1
2、有效的常量池索引是1~n-1。0是无效索引表示不指向任何常量
3、CONSTANT_Long_info和CONSTANT_Double_info各占两个位置。如果常量池中存在这两种常量，实际的常量数量比n-1少
 */
type ConstantPool []ConstantInfo


/*
读取常量池
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ { //
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ //
		}
	}
	return cp
}

/*
按索引查找常量
 */

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Errorf("Invalid constant pool index: %v!", index))
}
/*
从常量池查找字段或方法的名字和描述符
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/*
从常量池查找类名
 */
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

/*
从常量池查找UTF-8字符串
 */
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}