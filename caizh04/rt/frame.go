package rt

type Frame struct {
	lower            *Frame  //实现链表结构，下一个帧
	localVars        LocalVars  //保存局部变量指针
	operandStack     *OperandStack  //保存操作数栈指针
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}
func (self *Frame) OperandStack() *OperandStack {
	return  self.operandStack
}
/*
创建Frame实例
 */
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}