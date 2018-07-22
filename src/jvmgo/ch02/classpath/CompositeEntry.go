package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

// 参数（路径列表）按分割符分成小路径，然后把每个小路径 都转换成具体的Entry实例
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 依次调用每一个子路径的readClass()方法，如果成功读取到class数据 ，返回数据；如果收到错误信息，则继续；如果遍历完所有的子路经还没有找到class文件，则返回错误。
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
