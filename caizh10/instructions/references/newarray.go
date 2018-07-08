package references

import "../../instructions/base"
import "../../rt"
import "../../rt/heap"

const (
	//数组类型
	AT_BOOLEAN = 4
	AT_CHAR    = 5
	AT_FLOAT   = 6
	AT_DOUBLE  = 7
	AT_BYTE    = 8
	AT_SHORT   = 9
	AT_INT     = 10
	AT_LONG    = 11
)

// newarray指令用来创建基本类型数组，包括boolean[]、byte[]、
//char[]、short[]、int[]、long[]、float[]和double[]8种。
type NEW_ARRAY struct {
	atype uint8
}

func (self *NEW_ARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}
func (self *NEW_ARRAY) Execute(frame *rt.Frame) {
	//newarray指令的第二个操作数是count，从操作数栈中弹出，表
	//示数组长度。Execute（）方法根据atype和count创建基本类型数组
	stack := frame.OperandStack()
	count := stack.PopInt()
	//如果count小于0，则抛出NegativeArraySizeException异常，否则
	//根据atype值使用当前类的类加载器加载数组类，然后创建数组对
	//象并推入操作数栈。
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}
