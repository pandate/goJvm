package classfile
/*
属性表
各种属性表达的信息各不相同，因此无法用统一的结构来定义。不同之处在于，常量是由Java虚拟机规范严格定义的
属性是可以扩展的，不同的虚拟机实现可以定义自己的属性类型。由于这个原因，java虚拟机规范没有使用tag，而是使用属性名来区别不同属性。属性数据放在属性名之后的u1表中，这样java虚拟机实现就可以跳过自己无法识别的属性。属性表中实际存放的是常量池索引
 */
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

/*
读取属性表
 */
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}
/*
读取单个属性
 */
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32,
	cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code": return &CodeAttribute{cp: cp}
	case "ConstantValue": return &ConstantValueAttribute{}
	case "Deprecated": return &DeprecatedAttribute{}
	case "Exceptions": return &ExceptionsAttribute{}
	case "LineNumberTable": return &LineNumberTableAttribute{}
	case "LocalVariableTable": return &LocalVariableTableAttribute{}
	case "SourceFile": return &SourceFileAttribute{cp: cp}
	case "Synthetic": return &SyntheticAttribute{}
	default: return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
