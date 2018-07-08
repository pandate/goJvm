package math

import "../../instructions/base"
import "../../rt"

/*
位移指令
ishl 将int型数值左移位指定位数并将结果压入栈顶
ishr 将int型数值右（符号）移位指定位数并将结果压入栈顶
iushr 将int型数值右（无符号）移位指定位数并将结果压入栈顶
lshl 将long型数值左移位指定位数并将结果压入栈顶
lshr 将long型数值右（符号）移位指定位数并将结果压入栈顶
lushr 将long型数值右（无符号）移位指定位数并将结果压入栈顶
*/
type ISHL struct{ base.NoOperandsInstruction }
type ISHR struct{ base.NoOperandsInstruction }
type IUSHR struct{ base.NoOperandsInstruction }
type LSHL struct{ base.NoOperandsInstruction }
type LSHR struct{ base.NoOperandsInstruction }
type LUSHR struct{ base.NoOperandsInstruction }

func (self *ISHL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 << s
	stack.PushInt(result)
}

// Arithmetic shift right int

func (self *ISHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// Logical shift right int

func (self *IUSHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

// Shift left long

func (self *LSHL) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

// Arithmetic shift right long

func (self *LSHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Logical shift right long

func (self *LUSHR) Execute(frame *rt.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
