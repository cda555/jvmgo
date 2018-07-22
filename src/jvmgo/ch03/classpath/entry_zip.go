package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ZipEntry 压缩文件绝对路径
type ZipEntry struct {
	absPath string //存放ZIP或JAR文件的绝对路径
}

//返回绝对路径
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path) //获得绝对路径
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 从zip文件中提取class文件
// 	打开zip文件
//  遍历文件
//	如果找到class文件则返回，否则返回class not found
//	如果其中报错则返回错误信息
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath) //打开zip文件
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()            //当方法执行完成后关闭流
	for _, f := range r.File { //遍历文件
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

//String 此包的说明
func (self *ZipEntry) String() string {
	return self.absPath
}
