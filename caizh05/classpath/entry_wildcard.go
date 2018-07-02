package classpath

import (
	"strings"
	"os"
	"path/filepath"
)



func newWildcardEntry(path string) CompositeEntry{
	baseDir := path[:len(path)-1]
	compositeEntry :=[]Entry{}
	//根据后缀名选出jar文件，并返回skipDir跳过子目录
 	walkFn := func(path string,info os.FileInfo,err error) error{
 		if err != nil {
 			return err
		}
		if info.IsDir() && path != baseDir{
			return filepath.SkipDir
		}
		if strings.HasSuffix(path,".jar")||strings.HasSuffix(path,".JAR"){
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry,jarEntry)
		}
		return nil
	}
 	filepath.Walk(baseDir,walkFn)
	return compositeEntry
}


