package misc

import (
	"../../../native"
	"../../../rt"
	"../../../instructions/base"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}



func initialize(frame *rt.Frame) {
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}