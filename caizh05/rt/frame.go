package rt

type Frame struct {
	lower            *Frame  //实现链表结构，下一个帧
	localVars        LocalVars  //保存局部变量指针
	operandStack     *OperandStack  //保存操作数栈指针
	thread           *Thread
	nextPC           int
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return  self.operandStack
}
func (self *Frame) NextPC() int {
	return self.nextPC
}
func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}
func (self *Frame) Thread() *Thread{
	return self.thread
}

/*
创建Frame实例
 */
func newFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}