package base

import "../../rt"

/*
跳转函数  获取帧的pc指针加上偏移量后赋值给帧的下一个pc指针中
*/
func Branch(frame *rt.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}