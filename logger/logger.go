package logger

import (
	"fmt"
	"github.com/tabalt/appgo/file"
	"os"
)

var (
	lFile    *file.File
	typeList []string
)

// init logger package
func init() {
	// Log type list
	typeList = []string{"Notice", "Fatal", "Error", "Default", "Warning"}
}

// create Logger Object
//TODO 可指定日志文件
func Instance(logPath string) *Logger {
	lFile = file.Instance(logPath)
	return &Logger{}
}

type Logger struct {
}

//add log to file
func (this *Logger) Log(logType string, content string) {
	trueType := "Default"
	for _, typeName := range typeList {
		if typeName == logType {
			trueType = logType
		}
	}
	//TODO　log format
	content = fmt.Sprintf("%s : %s\n", trueType, content)
	lFile.Write(content)
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

// add Error log
func (this *Logger) Warning(logContent string) {
	this.Log("Warning", logContent)
}

// add Default log
func (this *Logger) Default(logContent string) {
	this.Log("Default", logContent)
}
