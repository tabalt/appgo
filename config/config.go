package config

import (
    "github.com/tabalt/appgo/logger"
    "github.com/tabalt/appgo/encoding/json"
)

type Config struct {
    file string
    Json *json.Json
    Logger *logger.Logger
    data map[string] interface{}
}

func NewConfig(fileName string) *Config {
    config := &Config{ file : fileName, Logger : logger.NewLogger("log/app.log"), Json : json.NewJson("")}
    jsonObject := config.Json.ReadFile(fileName)
    jsonMap, ok := jsonObject.(map[string] interface{})
    if ok {
        config.data = jsonMap
    } else {
        config.Logger.Fatal("Config file error")
    }
    return config
}

//保存配置到文件
func (this *Config) Save() bool {
    return this.Json.Save()
}

//获取指定配置项 需自行类型转换
func (this *Config) Get(key string) interface{} {
    var config interface{}
    config, err := this.data[key]
    if !err {
        this.Logger.Fatal("config key " + key + " no exists")
    }
    return config
}

//获取指定全部配置
func (this *Config) GetAll() interface{} {
    return this.Json.Get()
}

//设置指定配置项
func (this *Config) Set(key string, value interface{}) bool {
    this.data[key] = value
    _, err := this.data[key]
    if !err {
        this.Logger.Fatal("config key " + key + " no exists")
    }
    return true
}

//TODO　Set可选是否保存配置文件
