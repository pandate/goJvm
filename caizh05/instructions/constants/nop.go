package constants

import(
	"../base"
	"../../rt"
)

type NOP struct{ base.NoOperandsInstruction }
func (self *NOP) Execute(frame *rt.Frame) {
	// 什么也不用做
}