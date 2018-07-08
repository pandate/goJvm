package base

import "../../rt"
import (
	"../../rt/heap"
	"fmt"
)

/*
初始化类函数
*/
func InitClass(thread *rt.Thread, class *heap.Class) {
	fmt.Println("init class:",class)
	class.StartInit()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}
/*
获取初始化函数并加入栈帧中
 */
func scheduleClinit(thread *rt.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}
/*
初始化超类
 */
func initSuperClass(thread *rt.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		fmt.Println("init SuperClass :",superClass)
		if superClass != nil && !superClass.InitStarted() {
			InitClass(thread, superClass)
		}
	}
}
