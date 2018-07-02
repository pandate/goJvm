package references

import "../../instructions/base"
import "../../rt"

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// hack!
func (self *INVOKE_SPECIAL) Execute(frame *rt.Frame) {
	frame.OperandStack().PopRef()
}
