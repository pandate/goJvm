package heap

/*
在三种情况下，S类型的引用值可以赋值给T类型：S和T是同一类型；T是类且S是T的子类；
或者T是接口且S实现了T接口。这是简化版的判断逻辑，因为还没有实现数组
 */
func (self *Class) isAssignableFrom(other *Class) bool {
	s, t := other, self

	if s == t {
		return true
	}
	if !s.IsArray() {
		if !s.IsInterface() {
			if !t.IsInterface() {
				return s.IsSubClassOf(t)
			} else {
				return s.IsImplements(t)
			}
		} else {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isSuperInterfaceOf(s)
			}
		}
	} else {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.isJlObject()
			} else {
				return t.isJlCloneable() || t.isJioSerializable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
}

// 判断S是否是T的子类，实际上也就是判断T是否是S的（直接或间接）超类。
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// 判断S是否实现了T接口，就看S或S的（直接或间接）超类是否实现了某个接口T'
// ，T'要么是T，要么是T的子接口。
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// 与isSubClassOf类似但是用到了递归
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}



// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}

// iface extends self
func (self *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(self)
}

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}
func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}
func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}
