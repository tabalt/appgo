package file

import (
    "fmt"
    "os"
)

type File struct {
    fileName string
}

func (this *File) Init(fileName string) bool {
    this.fileName = fileName
    //init file...
    return true
}

//read file content
func (this *File)  ReadFile() (content string) {
    content = string(this.GetFileByte())
    return
}

//get file bytes
func (this *File)  GetFileByte() []byte {
    f, err := os.Open(this.fileName)
    defer f.Close()
    if err != nil {
        fmt.Println("Open file1 " + this.fileName + " Faild : ", err)
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
func (this *File)  WriteFile(content string) {
    f, err := os.OpenFile(this.fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
    if err != nil {
        fmt.Println("Open file2 " + this.fileName + " Faild : ", err)
        os.Exit(1)
    }
    defer f.Close()
    f.WriteString(content)
}

//删除文件 文件目录操作 读写加锁 etc.
