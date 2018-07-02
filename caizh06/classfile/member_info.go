package classfile

type MemberInfo struct {
	cp                  ConstantPool //常量池
	accessFlags         uint16  //访问标志
	nameIndex           uint16  //常量池索引给出字段名或方法名
	descriptorIndex     uint16  //常量池索引给出字段或方法的描述符
	attributes          []AttributeInfo  //属性表
}
/*
读取字段表或方法表
 */
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}
/*
读取字段或方法数据
 */
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:                  cp,
		accessFlags:         reader.readUint16(),
		nameIndex:           reader.readUint16(),
		descriptorIndex:     reader.readUint16(),
		attributes:          readAttributes(reader, cp),
	}
}
/*
从常量池查找字段或方法名
 */
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
/*
从常量池找字段或方法描述符
 */
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) ConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}
