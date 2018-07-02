package heap

/*
类符号引用
 */
type SymRef struct {
	cp        *ConstantPool  //存放符号引用所在的运行时常量池指针
	className string  //类全名
	class     *Class  //类结构体指针
}

/*
类符号解析
 */
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

/*
通俗地讲，如果类D通过符号引用N引用类C的话，要解析N，先用D的类加载器加载C，
然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常。
 */
func (self *SymRef) resolveClassRef() {
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
