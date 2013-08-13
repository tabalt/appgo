package config

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/tabalt/appgo/file"
	"github.com/tabalt/appgo/logger"
)

var (
	cLogger *logger.Logger
	cFile   *file.File
)

func init() {
	cLogger = logger.Instance("log/app.log")
}

// get config instance
func Instance(fileName string) *Config {
	cFile = file.Instance(fileName)
	cData := &ConfigData{boolList: make(map[string]bool), intList: make(map[string]int), float32List: make(map[string]float32), float64List: make(map[string]float64), stringList: make(map[string]string), arrayList: make(map[string][]interface{})}

	config := &Config{data: cData}
	config.Init()
	return config
}

// config data struct
type ConfigData struct {
	boolList    map[string]bool
	intList     map[string]int
	float32List map[string]float32
	float64List map[string]float64
	stringList  map[string]string
	arrayList   map[string][]interface{} 
}

type Config struct {
	data *ConfigData
}

func (this *Config) Init() {
	jsonBytes := cFile.GetBytes()

	var jsonObject interface{}
	if err := json.Unmarshal(jsonBytes[:len(jsonBytes)], &jsonObject); err != nil {
		cLogger.Fatal(err.Error())
	}
	jsonMap, ok := jsonObject.(map[string]interface{})
	if !ok {
		cLogger.Fatal("Config file error")
		os.Exit(1)
	}

	for k, v := range jsonMap {
		switch val := v.(type) {
			case bool:
				this.data.boolList[k] = val
			case int:
				this.data.intList[k] = val
			case float32:
				this.data.float32List[k] = val
			case float64:
				this.data.float64List[k] = val
			case string:
				this.data.stringList[k] = val
			case []interface{}:
				
				/*strList := make(map[int] string)
				fmt.Println(len(val))
				for i, iv := range val {
					fmt.Println(i, iv)
					//strList[i] = string(iv.String())
				}*/
				this.data.arrayList[k] = val
			default:
				cLogger.Warning("Ignore config key :" + k )
				fmt.Println("Ignore config key :" + k )
		}
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
		cLogger.Fatal("config key " + key + " no exists")
	}
	return config
}


//设置指定配置项
func (this *Config) Set(key string, value interface{}) bool {
	this.data[key] = value
	_, err := this.data[key]
	if !err {
		cLogger.Fatal("config key " + key + " no exists")
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
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Int returns the integer value for a given key.
func (this *Config) Int(key string) (int, error) {
	v, err := this.data.intList[key]
	if !err {
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Float32 returns the float value for a given key.
func (this *Config) Float32(key string) (float32, error) {
	v, err := this.data.float32List[key]
	if !err {
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// Float64 returns the float value for a given key.
func (this *Config) Float64(key string) (float64, error) {
	v, err := this.data.float64List[key]
	if !err {
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// String returns the string value for a given key.
func (this *Config) String(key string) (string, error) {
	v, err := this.data.stringList[key]
	if !err {
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}

// array returns the string value for a given key.
func (this *Config) Array(key string) ([]interface{}, error) {
	v, err := this.data.arrayList[key]
	if !err {
		cLogger.Fatal("config key " + key + " not exists")
	}
	return v, nil
}
