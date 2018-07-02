package classfile
/*
出字段或方法的名称和描述符
Java虚拟机规范定义了一种简单的语法来描述字段和方法，可以根据下面的规则生成描述符。
1）类型描述符。
①基本类型byte、short、char、int、long、float和double的描述符是单个字母，分别对应B、S、C、I、J、F和D。注意，long的描述符是J而不是L。
②引用类型的描述符是L＋类的完全限定名＋分号。 如Ljava.lang.Object;
③数组类型的描述符是[＋数组元素类型描述符。[I表示int[]   [[D表示double[][] [Ljava.lang.Object;表示java.lang.Object[]
2）字段描述符就是字段类型的描述符。
3）方法描述符是（分号分隔的参数类型描述符）+返回值类型描述符，其中void返回值由单个字母V表示
（）V  表示 void run()
()Ljava.lang.String; 表示String toString()
(Ljava.lang.String;)V 表示 void main(String[] args)
(FF)F float max(float x,floaty)
 */

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16 //字段或方法名索引
	descriptorIndex uint16 //字段或方法名的描述符索引
}
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}