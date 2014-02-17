package appgo

import (
	"fmt"
	"github.com/tabalt/appgo/config"
	"github.com/tabalt/appgo/logger"
    //"strconv"
)

const VERSION = "0.0.2"

var (
	AppName    string
	AppPath    string
	ConfigPath string
	LogPath    string
	AppConfig     *config.Config
	AppLogger     *logger.Logger
)

func init() {
    fmt.Println("init appgo")
    ConfigPath = "conf/config.json"
    LogPath = "log/app.log"
    AppConfig = config.Instance(ConfigPath)
    AppLogger = logger.Instance(LogPath)
}

type App struct {

	//...
}


type Model interface {
	Add() 
    Edit()
    Find()
    Exist()
    List()
    Del()
}

func HasError(err error) (result bool){
    result = false
    if err != nil {
        result = true
        AppLogger.Error(err.Error())
    }
    return
}

