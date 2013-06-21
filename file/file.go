package file

import (
    "fmt"
    "os"
)

type File struct {
    fileName string

}

func (file *File) Init(fileName string) bool {
    file.fileName = fileName
    //init file...
    return true
}

//read file content
func (file *File)  ReadFile() (content string) {
    content = string(file.GetFileByte())
    return
}

//get file bytes
func (file *File)  GetFileByte() []byte {
    f, err := os.Open(file.fileName)
    defer f.Close()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    buf := make([]byte, 1024)
    for {
        n, _ := f.Read(buf)
        if 0 == n {
            break
        }
    }
    return buf
}

//write content to file
func (file *File)  WriteFile(content string) {
    f, err := os.OpenFile(file.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
    if err != nil {
        panic(err)
        return
    }
    defer f.Close()
    f.WriteString(content)
}

//删除文件 文件目录操作 读写加锁 etc.
