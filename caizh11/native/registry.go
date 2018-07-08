package native

import "../rt"

/*
把本地方法定义成一个函数，参数是Frame结构体指针，没有
返回值。这个frame参数就是本地方法的工作空间，也就是连接Java
虚拟机和Java类库的桥梁
 */
type NativeMethod func( frame *rt.Frame)

var registry = map[string]NativeMethod{}
/*
类名、方法名和方法描述符加在一起才能唯一确定一个方法，
所以把它们的组合作为本地方法注册表的键，Register（）函数把前
述三种信息和本地方法实现关联起来。
 */
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

/*
根据类名、方法名和方法描述符查找
本地方法实现，如果找不到，则返回nil。
 */
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	return nil
}
func emptyNativeMethod(frame *rt.Frame) {
	// do nothing
}