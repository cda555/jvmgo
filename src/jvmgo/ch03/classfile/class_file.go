package classfile

import (
	"fmt"
)

type ClassFile struct {
	minorVersion uint16       // 副版本号
	majorVersion uint16       // 主版本号
	constantPool ConstantPool //常量池
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

// Parse 把[]byte解析成ClassFile结构体
func Parse(classDate []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classDate}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// 解析class字节
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)               // 魔数
	self.readAndCheckVersion(reader)             // 版本号
	self.constantPool = readConstantPool(reader) // 解析常量池
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 字段和方法表 见 3.2.8
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool) // 属性表 见 3.4
}

// MajorVersion 主版本号
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

// ClassName 从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 从常量池查找超类
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //java.lang.Object没有超类
}

// InterfaceNames 从常量池中查找接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceName := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceName[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceName
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 检查版本号
func (self *ClassFile) readAndcheckversion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
