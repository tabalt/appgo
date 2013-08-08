package file

import (
	"fmt"
	"os"
)

type File struct {
	name string
}

// create File Object
func NewFile(fileName string) *File {
	return &File{name: fileName}
}

//read file content
func (this *File) Read() (content string) {
	f, err := os.Open(this.name)
	defer f.Close()
	if err != nil {
		fmt.Println("Read file Faild. ", err.Error())
		os.Exit(1)
	}
	buf := make([]byte, 1)
	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		} else {
			content += string(buf)
		}
	}
	return
}

//get file bytes
func (this *File) GetBytes() []byte {
	content := this.Read()
	buf := []byte(content)
	return buf
}

//write content to file
func (this *File) Write(content string) {
	f, err := os.OpenFile(this.name, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println("Write file Faild. ", err.Error())
		os.Exit(1)
	}
	defer f.Close()
	f.WriteString(content)
}

//删除文件 文件目录操作 读写加锁 etc.
