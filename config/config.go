package config

import (
	"encoding/json"
	"fmt"
	"github.com/tabalt/appgo/file"
	"github.com/tabalt/appgo/logger"
)

var (
	loggerObj *logger.Logger
	fileObj   *file.File
)

func init() {
	loggerObj = logger.NewLogger("log/app.log")
}

type ConfigData struct {
	boolList    map[string]bool
	intList     map[string]int
	float32List map[string]float32
	float64List map[string]float64
	stringList  map[string]string
	arrayList   map[string][]string
}

type Config struct {
	file string
	data ConfigData
}

func NewConfig(fileName string) *Config {
	fileObj = file.NewFile(fileName)
	config := &Config{file: fileName}
	config.Init()
	return config
}

func (this *Config) Init() {
	jsonBytes := fileObj.GetBytes()
	//fmt.Println(jsonBytes)
	//fmt.Println(len(fileObj.Read()))

	var jsonObject interface{}
	if err := json.Unmarshal(jsonBytes[:len(jsonBytes)], &jsonObject); err != nil {
		loggerObj.Fatal(err.Error())
	}
	jsonMap, ok := jsonObject.(map[string]interface{})
	if ok {
		//config.data = jsonMap
		for k, v := range jsonMap {
			switch v.(type) {
			case bool:
				this.data.boolList[k] = v.(bool)
				fmt.Println(k)
			case int:
				this.data.intList[k] = v.(int)
				fmt.Println(k)
			case float32:
				this.data.float32List[k] = v.(float32)
				fmt.Println(k)
			case float64:
				this.data.float64List[k] = v.(float64)
				fmt.Println(k)
			case string:
				this.data.stringList[k] = v.(string)
				fmt.Println(k)
			case []string:
				this.data.arrayList[k] = v.([]string)
				fmt.Println(k)
			default:
				loggerObj.Warning("Unknown Config key :" + k + " , type : ")
			}
		}
	} else {
		loggerObj.Fatal("Config file error")
	}
}

//保存配置到文件
func (this *Config) Save() bool {
	return true
}

/*
//获取指定配置项 需自行类型转换
func (this *Config) Get(key string) interface{} {
	var config interface{}
	config, err := this.data[key]
	if !err {
		loggerObj.Fatal("config key " + key + " no exists")
	}
	return config
}


//设置指定配置项
func (this *Config) Set(key string, value interface{}) bool {
	this.data[key] = value
	_, err := this.data[key]
	if !err {
		loggerObj.Fatal("config key " + key + " no exists")
	}
	return true
}
*/

/*
//统一错误处理函数
func (this *Config) CheckError(err error, key string) ret bool {
	if !err {
		this.Logger.Fatal("config key " + key + " no exists")
		ret = false
	}
	ret = true
	return
}
*/

//TODO　Set可选是否保存配置文件

// Bool returns the boolean value for a given key.
func (this *Config) Bool(key string) (bool, error) {
	v, err := this.data.boolList[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Int returns the integer value for a given key.
func (this *Config) Int(key string) (int, error) {
	v, err := this.data.intList[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Float32 returns the float value for a given key.
func (this *Config) Float32(key string) (float32, error) {
	v, err := this.data.float32List[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Float64 returns the float value for a given key.
func (this *Config) Float64(key string) (float64, error) {
	v, err := this.data.float64List[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// String returns the string value for a given key.
func (this *Config) String(key string) (string, error) {
	v, err := this.data.stringList[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// array returns the string value for a given key.
func (this *Config) Array(key string) ([]string, error) {
	v, err := this.data.arrayList[key]
	if !err {
		loggerObj.Fatal("config key " + key + " not exists")
	}
	return v, nil
}
