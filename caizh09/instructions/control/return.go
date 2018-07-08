package control

import "../../instructions/base"
import "../../rt"

/*
方法返回指令不需要操作数
*/
type RETURN struct{ base.NoOperandsInstruction }
type ARETURN struct{ base.NoOperandsInstruction }
type DRETURN struct{ base.NoOperandsInstruction }
type FRETURN struct{ base.NoOperandsInstruction }
type IRETURN struct{ base.NoOperandsInstruction }
type LRETURN struct{ base.NoOperandsInstruction }

func (self *RETURN) Execute(frame *rt.Frame) {
	frame.Thread().PopFrame()
}

func (self *ARETURN) Execute(frame *rt.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	ref := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(ref)
}

func (self *DRETURN) Execute(frame *rt.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(val)
}

func (self *FRETURN) Execute(frame *rt.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(val)
}

func (self *IRETURN) Execute(frame *rt.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(val)
}

func (self *LRETURN) Execute(frame *rt.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	val := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(val)
}
