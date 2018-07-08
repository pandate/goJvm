package rt
/*
使用链表数据结构来实现java虚拟机栈，这样栈就可以按需使用内存空间，而且弹出的帧也可以及时被Go的垃圾收集器回收
 */
type Stack struct {
	maxSize     uint  //栈的容量
	size        uint  //栈当前大小
	_top        *Frame  //保存栈顶指针
}
/*
实例化栈
 */
func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
/*
把帧推入栈顶如果栈满了则抛出异常
 */
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
/*
把栈顶帧弹出
 */
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

/*
返回栈顶帧但不弹出
 */
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}