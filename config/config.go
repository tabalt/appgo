package config

import (
    "encoding/json"
    "github.com/tabalt/appgo/logger"
    "io/ioutil"
)

type Config struct {
    file string
    Logger  logger.Logger
    data map[string]interface{}
}

//从文件读取配置
func (this *Config) File2Config(fileName string) bool {
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        this.Logger.Fatal(err.Error())
    }
    result := true
    if err := json.Unmarshal(file, &this.data); err != nil {
        result = false
    }
    return result
}

//将配置写入文件
func (this *Config) Config2File(fileName string) bool {
    configContent, err := json.Marshal(this.data)
    if err != nil {
        this.Logger.Fatal(err.Error())
    }
    if err := ioutil.WriteFile(fileName, configContent, 0x777); err != nil {
        this.Logger.Fatal(err.Error())
    }
    return true
}

//初始化配置
func (this *Config) Init(configFile string, logFile string) bool {
    this.file = configFile
    this.Logger.Init(logFile)
    return this.File2Config(this.file)
}

//保存配置到文件
func (this *Config) Save() bool {
    return this.Config2File(this.file)
}

//获取指定配置项 需自行类型转换
func (this *Config) Get(key string) interface{} {
    var config interface{}
    config, err := this.data[key]
    if !err {
        this.Logger.Fatal("配置项“" + key + "”不存在")
    }
    return config
}

//获取指定全部配置
func (this *Config) GetAll() interface{} {
    return this.data
}

//设置指定配置项
func (this *Config) Set(key string, value interface{}) bool {
    this.data[key] = value
    _, err := this.data[key]
    if !err {
        this.Logger.Fatal("配置项“" + key + "”不存在")
    }
    return true
}

//TODO　Set可选是否保存配置文件
