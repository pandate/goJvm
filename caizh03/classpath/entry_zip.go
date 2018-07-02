package classpath

import (
	"path/filepath"
	"archive/zip"
	"io/ioutil"
	"errors"
)

type ZipEntry struct {
	absPath string
}
/*
 先把参数转换成绝对路径，如果转换过程出现错误，则调用panic函数终止程序执行，否则创建DirEntry实例并返回
 */
func newZipEntry(path string) *ZipEntry{
	absDir,err :=filepath.Abs(path)
	if err !=nil{
		panic(err)
	}
	return &ZipEntry{absDir}
}

/*
首先打开zip文件，如果这步出错的话直接返回。然后遍历ZIP压缩包里的文件，看能否找到class文件。如果
能找到，则打开class文件，把内容读取并返回，如果找不到，或出现其他错误则返回错误信息。
 */
func (self *ZipEntry) readClass(className string)([]byte,Entry,error){
	r,err :=zip.OpenReader(self.absPath)
	if err != nil{
		return nil,nil,err
	}
	defer r.Close()
	for _, f:= range r.File{
		if f.Name == className {
			rc,err :=f.Open()
			if err != nil{
				return nil,nil,err
			}
			defer rc.Close()
			data,err := ioutil.ReadAll(rc)
			if err!= nil {
				return nil,nil,err
			}
			return data,self,nil
		}
	}
	return nil,nil,errors.New("class not found:"+className)
}

func (self *ZipEntry) String() string  {
	return self.absPath
}