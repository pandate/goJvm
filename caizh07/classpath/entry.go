package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/*
	路径之间用斜线分隔，文件名有.class后缀
	比如要读取java.lang.Object类，传入的参数应该是java/lang/Object.class。
	返回值是读取到的字节数据、最终定位到class文件的Entry，以及错误信息。
	 */
	readClass(className string)([]byte,Entry,error)
	/*
	作用相当于java中的toString()，用于返回变量的字符串表示
	 */
	String() string


}
func newEntry(path string) Entry{
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR")||
		strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP"){
		return newZipEntry(path)
	}
	return newDirEntry(path)
}

