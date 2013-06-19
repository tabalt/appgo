package config

import (
    "encoding/json"
    "fmt"
    //"github.com/tabalt/golib/logger"
    "io/ioutil"
)

type Config struct {
    file string
    data map[string]interface{}
}

//从文件读取配置
func (conf *Config) File2Config(fileName string) bool {
    file, err := ioutil.ReadFile(fileName)
    if err != nil {
        //logger.Fatal(err.Error())
        fmt.Println(err.Error())
    }
    result := true
    if err := json.Unmarshal(file, &conf.data); err != nil {
        result = false
    }
    return result
}

//将配置写入文件
func (conf *Config) Config2File(fileName string) bool {
    configContent, err := json.Marshal(conf.data)
    if err != nil {
        //logger.Fatal(err.Error())
        fmt.Println(err.Error())
    }
    if err := ioutil.WriteFile(fileName, configContent, 0x777); err != nil {
        //logger.Fatal(err.Error())
        fmt.Println(err.Error())
    }
    return true
}

//初始化配置
func (conf *Config) Init(fileName string) bool {
    conf.file = fileName
    return conf.File2Config(conf.file)
}


//保存配置到文件
func (conf *Config) Save() bool {
    return conf.Config2File(conf.file)
}

//获取指定配置项 需自行类型转换
func (conf *Config) Get(key string) interface{} {
    var config interface{}
    config, err := conf.data[key]
    if !err {
        //logger.Fatal("配置项“" + key + "”不存在")
        fmt.Println("配置项“" + key + "”不存在")
    }
    return config
}

//获取指定全部配置
func (conf *Config) GetAll() interface{} {
    return conf.data
}

//设置指定配置项
func (conf *Config) Set(key string, value interface{}) bool {
    conf.data[key] = value
    _, err := conf.data[key]
    if !err {
        //logger.Fatal("配置项“" + key + "”不存在")
        fmt.Println("配置项“" + key + "”设置失败")
    }
    return true
}

//TODO　Set可选是否保存配置文件
