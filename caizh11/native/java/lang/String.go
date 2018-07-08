package lang

import "../../../native"
import "../../../rt"
import "../../../rt/heap"

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// 如果字符串还没有入池，把它放入并返回该字符串，否则找到
//已入池字符串并返回。
func intern(frame *rt.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
