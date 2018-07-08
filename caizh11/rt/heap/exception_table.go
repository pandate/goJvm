package heap

import "../../classfile"

type ExceptionTable []*ExceptionHandler

type ExceptionHandler struct {
	startPc int
	endPc int
	handlerPc int
	catchType *ClassRef
}
/*
newExceptionTable（）函数把class文件中的异常处理表转换成
ExceptionTable类型。有一点需要特别说明：异常处理项的catchType
有可能是0。我们知道0是无效的常量池索引，但是在这里0并非表
示catch-none，而是表示catch-all
 */
func newExceptionTable(entries []*classfile.ExceptionTableEntry,
	cp *ConstantPool) ExceptionTable {
	table := make([]*ExceptionHandler, len(entries))
	for i, entry := range entries {
		table[i] = &ExceptionHandler{
			startPc: int(entry.StartPc()),
			endPc: int(entry.EndPc()),
			handlerPc: int(entry.HandlerPc()),
			catchType: getCatchType(uint(entry.CatchType()), cp),
		}
	}
	return table
}
/*
从运行时常量池中查找类符号引用
 */
func getCatchType(index uint, cp *ConstantPool) *ClassRef {
	if index == 0 {
		return nil
	}
	return cp.GetConstant(index).(*ClassRef)
}

func (self ExceptionTable) findExceptionHandler(exClass *Class,
	pc int) *ExceptionHandler {
	for _, handler := range self {
		if pc >= handler.startPc && pc < handler.endPc {
			if handler.catchType == nil {
				return handler // catch-all
			}
			catchClass := handler.catchType.ResolvedClass()
			if catchClass == exClass || catchClass.IsSuperClassOf(exClass) {
				return handler
			}
		}
	}
	return nil
}