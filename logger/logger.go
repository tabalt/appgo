package logger

import (
    "fmt"
    "github.com/tabalt/appgo/file"
    "os"
)
type Logger struct {
    typeList []string
    File file.File
}

//init logger
func (this *Logger) Init(fileName string) bool {
    allowTypeList := []string{"Notice", "Fatal", "Error", "Default"}
    this.typeList = allowTypeList
    return this.File.Init(fileName)
}

func (this *Logger) saveLog(logType string, logContent string) {
    trueType := "Default"
    for _, typeName := range this.typeList {
        if typeName == logType {
            trueType = logType
        }
    }
    content := fmt.Sprintf("%s : %s\n", trueType, logContent)
    this.File.WriteFile(content)
}

func (this *Logger) Notice(logContent string) {
    this.saveLog("Notice", logContent)
}

func (this *Logger) Fatal(logContent string) {
    this.saveLog("Fatal", logContent)
    os.Exit(1)
}

func (this *Logger) Error(logContent string) {
    this.saveLog("Error", logContent)
}

func (this *Logger) Default(logContent string) {
    this.saveLog("Default", logContent)
}
