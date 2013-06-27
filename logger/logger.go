package logger

import (
    "fmt"
    "github.com/tabalt/appgo/file"
    "os"
)
type Logger struct {
    typeList []string
    File *file.File
}

// create Logger Object
//TODO 可指定日志文件
func NewLogger() *Logger {
    // log type list
    logTypeList := []string{"Notice", "Fatal", "Error", "Default"}
    //fileObj := file.NewFile("log/app.log")
    return &Logger{ typeList: logTypeList, File : file.NewFile("log/app.log")}
}

//add log to file
func (this *Logger) Log(logType string, content string) {
    trueType := "Default"
    for _, typeName := range this.typeList {
        if typeName == logType {
            trueType = logType
        }
    }
    //TODO　log format
    content = fmt.Sprintf("%s : %s\n", trueType, content)
    this.File.Write(content)
}

// add Notice log
func (this *Logger) Notice(logContent string) {
    this.Log("Notice", logContent)
}

// add Fatal log
func (this *Logger) Fatal(logContent string) {
    this.Log("Fatal", logContent)
    os.Exit(1)
}

// add Error log
func (this *Logger) Error(logContent string) {
    this.Log("Error", logContent)
}

// add Default log
func (this *Logger) Default(logContent string) {
    this.Log("Default", logContent)
}
