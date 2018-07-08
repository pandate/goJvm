package heap

import "../../classfile"
/*
字段和方法符号引用共有的信息
 */

type MemberRef struct {
	SymRef
	name       string  //成员名称
	descriptor string  //字段描述符  在虚拟机规范里一个类可以有多个同名字段，只要类型不同就行
}
/*
从class文件内存储的字段或方法常量中提取数据
 */
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
