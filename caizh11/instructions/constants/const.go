package constants

import "../../instructions/base"
import "../../rt"

/*
aconst_null null进栈
dconst_0    double型常量值0进栈
dconst_1    double型常量值1进栈
fconst_0    float型常量值0进栈
fconst_1    float型常量值1进栈
fconst_2    float型常量值2进栈
iconst_m1   int型常量值-1进栈
iconst_0    int型常量值0进栈
iconst_1   int型常量值1进栈
iconst_2   int型常量值1进栈
iconst_3   int型常量值1进栈
iconst_4   int型常量值1进栈
iconst_5   int型常量值1进栈
lconst_0   long型常量值0进栈
lconst_1   long型常量值1进栈
 */
type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }


func (self *ACONST_NULL) Execute(frame *rt.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Push double

func (self *DCONST_0) Execute(frame *rt.Frame) {
	frame.OperandStack().PushDouble(0.0)
}


func (self *DCONST_1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Push float

func (self *FCONST_0) Execute(frame *rt.Frame) {
	frame.OperandStack().PushFloat(0.0)
}


func (self *FCONST_1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushFloat(1.0)
}


func (self *FCONST_2) Execute(frame *rt.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Push int constant

func (self *ICONST_M1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(-1)
}


func (self *ICONST_0) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(0)
}


func (self *ICONST_1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(1)
}


func (self *ICONST_2) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(2)
}


func (self *ICONST_3) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(3)
}


func (self *ICONST_4) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(4)
}


func (self *ICONST_5) Execute(frame *rt.Frame) {
	frame.OperandStack().PushInt(5)
}

// Push long constant

func (self *LCONST_0) Execute(frame *rt.Frame) {
	frame.OperandStack().PushLong(0)
}


func (self *LCONST_1) Execute(frame *rt.Frame) {
	frame.OperandStack().PushLong(1)
}
