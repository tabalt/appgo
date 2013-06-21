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
func (log *Logger) Init(fileName string) bool {
    allowTypeList := []string{"Notice", "Fatal", "Error", "Default"}
    log.typeList = allowTypeList
    log.File = file.File{}
    return log.File.Init(fileName)
}

func (log *Logger) saveLog(logType string, logContent string) {
    trueType := "Default"
    for _, typeName := range log.typeList {
        if typeName == logType {
            trueType = logType
        }
    }
    content := fmt.Sprintf("%s : %s\n", trueType, logContent)
    log.File.WriteFile(content)
}

func (log *Logger) Notice(logContent string) {
    log.saveLog("Notice", logContent)
}

func (log *Logger) Fatal(logContent string) {
    log.saveLog("Fatal", logContent)
    os.Exit(1)
}

func (log *Logger) Error(logContent string) {
    log.saveLog("Error", logContent)
}

func (log *Logger) Default(logContent string) {
    log.saveLog("Default", logContent)
}
