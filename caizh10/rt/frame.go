package rt

import "./heap"

type Frame struct {
	lower            *Frame  //实现链表结构，下一个帧
	localVars        LocalVars  //保存局部变量指针
	operandStack     *OperandStack  //保存操作数栈指针
	thread           *Thread
	method           *heap.Method
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
func (self *Frame) Method() *heap.Method {
	return self.method
}
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}


/*
创建Frame实例
 */
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:           thread,
		method:           method,
		localVars:        newLocalVars(method.MaxLocals()),
		operandStack:     newOperandStack(method.MaxStack()),
	}
}