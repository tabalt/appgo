package json

import (
	"encoding/json"
	"github.com/tabalt/appgo/logger"
	"io/ioutil"
)

// create Json Object
func NewJson(fileName string) *Json {
	return &Json{fileName: fileName, Logger: logger.NewLogger("log/app.log")}
}

type Json struct {
	fileName string
	Logger   *logger.Logger
	object   interface{}
}

//Read Json Object from file
func (this *Json) ReadFile(fileName string) interface{} {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		this.Logger.Fatal(err.Error())
	}
	var jsonObject interface{}
	if err := json.Unmarshal(file, &jsonObject); err != nil {
		this.Logger.Fatal(err.Error())
	}
	return jsonObject
}

//Write Json Object to file
func (this *Json) WriteFile(fileName string, jsonObject interface{}) bool {
	content, err := json.Marshal(jsonObject)
	if err != nil {
		this.Logger.Fatal(err.Error())
	}
	if err := ioutil.WriteFile(fileName, content, 0x777); err != nil {
		this.Logger.Fatal(err.Error())
	}
	return true
}

//Read Json Object
func (this *Json) Read() bool {
	this.object = this.ReadFile(this.fileName)
	return true
}

//Save Json Object
func (this *Json) Save() bool {
	return this.WriteFile(this.fileName, this.object)
}

//Get json object
func (this *Json) Get() interface{} {
	return this.object
}

//Set json object
func (this *Json) Set(jsonObject interface{}) bool {
	this.object = jsonObject
	return true
}

//TODO　Set可选是否保存配置文件
