package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}
/*
 先把参数转换成绝对路径，如果转换过程出现错误，则调用panic函数终止程序执行，否则创建DirEntry实例并返回
 */
func newDirEntry(path string) *DirEntry{
	absDir,err :=filepath.Abs(path)
	if err !=nil{
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string)([]byte,Entry,error){
	fileName :=filepath.Join(self.absDir,className)
	data,err:=ioutil.ReadFile(fileName)
	return data,self,err
}

func (self *DirEntry) String() string  {
	return self.absDir
}