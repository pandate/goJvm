package rt

import (
	"./heap"
	"fmt"
)
/*
线程
 */
type Thread struct {
	pc     int //pc寄存器
	stack *Stack //虚拟机栈指针
}
/*
创建栈结构体实例，它的参数表示要创建的栈可以容纳多少帧
 */
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
/*
向栈中写入帧
 */
func (self *Thread) PushFrame(frame *Frame) {
	fmt.Println("向栈中写入帧:",frame.method.Name())
	self.stack.push(frame)
}
/*
从栈中获取下一帧
 */
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
/*
返回当前帧
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}


func (self *Thread) TopFrame() *Frame {
	return self.stack.top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}
func (self *Thread) SetPC(i int) {
	self.pc = i
}
func (self *Thread) PC() int {
	return self.pc
}