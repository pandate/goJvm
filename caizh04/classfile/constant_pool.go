package classfile
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
	panic("Invalid constant pool index!")
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

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer: return &ConstantIntegerInfo{}
	case CONSTANT_Float: return &ConstantFloatInfo{}
	case CONSTANT_Long: return &ConstantLongInfo{}
	case CONSTANT_Double: return &ConstantDoubleInfo{}
	case CONSTANT_Utf8: return &ConstantUtf8Info{}
	case CONSTANT_String: return &ConstantStringInfo{cp: cp}
	case CONSTANT_Class: return &ConstantClassInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case CONSTANT_NameAndType: return &ConstantNameAndTypeInfo{}
	//case CONSTANT_MethodType: return &ConstantMethodTypeInfo{}
	//case CONSTANT_MethodHandle: return &ConstantMethodHandleInfo{}
	//case CONSTANT_InvokeDynamic: return &ConstantInvokeDynamicInfo{}
	default: panic("java.lang.ClassFormatError: constant pool tag!")
	}
}