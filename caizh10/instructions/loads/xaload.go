package loads

import "../../instructions/base"
import "../../rt"
import "../../rt/heap"

/*
iastore指令的三个操作数分别是：要赋给数组元素的值、数组
索引、数组引用，依次从操作数栈中弹出。如果数组引用是null，则
抛出NullPointerException。如果数组索引小于0或者大于等于数组
长度，则抛出ArrayIndexOutOfBoundsException异常。这两个检查和
<t>aload系列指令一样。如果一切正常，则按索引给数组元素赋值。
*/
type AALOAD struct{ base.NoOperandsInstruction }
type BALOAD struct{ base.NoOperandsInstruction }
type CALOAD struct{ base.NoOperandsInstruction }
type DALOAD struct{ base.NoOperandsInstruction }
type FALOAD struct{ base.NoOperandsInstruction }
type IALOAD struct{ base.NoOperandsInstruction }
type LALOAD struct{ base.NoOperandsInstruction }
type SALOAD struct{ base.NoOperandsInstruction }


/*
首先从操作数栈中弹出第一个操作数：数组索引，然后弹出第
二个操作数：数组引用。如果数组引用是null，则抛出
NullPointerException异常。
如果一切正常，则按索引取出数组元素，推入操作数栈顶。
 */
func (self *AALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
}

// Load byte or boolean from array

func (self *BALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	bytes := arrRef.Bytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
}

// Load char from array

func (self *CALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	chars := arrRef.Chars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
}

// Load double from array

func (self *DALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	doubles := arrRef.Doubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
}

// Load float from array

func (self *FALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	floats := arrRef.Floats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
}

// Load int from array

func (self *IALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
}

// Load long from array

func (self *LALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	longs := arrRef.Longs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
}

// Load short from array

func (self *SALOAD) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	shorts := arrRef.Shorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
}

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
/*
如果数组索引小于0，或者大于等于数组长度，则抛出
ArrayIndexOutOfBoundsException。
 */
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}
