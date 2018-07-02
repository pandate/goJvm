package heap

import (
	"../../classpath"
	"fmt"
	"../../classfile"
)
/*
类加载器
类的加载大致可以分为三个步骤：
首先找到class文件并把数据读取到内存；
然后解析class文件，生成虚拟机可以使用的类数据，并放入方法区；
最后进行链接。
 */
type ClassLoader struct {
	cp       *classpath.Classpath  //classpath指针
	classMap map[string]*Class // 已经加载的类数据，key是类全名。方法区的具体实现
}

func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:           cp,
		classMap:    make(map[string]*Class),
	}
}
/*
把类数据加载到方法区
先查找classMap，看类是否已经被加载。如果是，直接返回类数据，否则调用
loadNonArrayClass（）方法加载类。数组类和普通类有很大的不同，它的数据
并不是来自class文件，而是由Java虚拟机在运行期间生成。
 */
func (self *ClassLoader) LoadClass(name string) *Class {
	if class, ok := self.classMap[name]; ok {
		return class // 类 已 经 加 载
	}
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}
/*
调用了Classpath的ReadClass（）方法，并进行了错误处理。需要解释一下它的返回值。为了
打印类加载信息，把最终加载class文件的类路径项也返回给了调用者。
 */
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}
/*
首先调用parseClass（）函数把class文件数据转换成Class结构体。Class结构体的superClass
和interfaces字段存放超类名和直接接口表，这些类名其实都是符号引用。
调用resolveSuperClass（）和resolveInterfaces（）函数解析这些类符号引用。
 */
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}
/*
递归调用loadClass方法加载超类
 */
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
/*
递归调用loadClass方法加载直接接口
 */
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)
	prepare(class)
}
/*
类验证
 */
func verify(class *Class) {
	// todo
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}
/*
计算实例字段的个数，同时给它们编号
 */
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}
/*
计算静态字段的个数，同时给他们编号
 */
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}
/*
给类变量分配空间，让后给它们赋予初始值
 */
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

/*
如果静态变量属于基本类型或String类型，有final修饰符，且它的值在编译期已知，则
该值存储在class文件常量池中
从常量池中加载常量值，然后给静态变量赋值
 */
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}