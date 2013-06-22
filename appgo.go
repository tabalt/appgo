package appgo

import (
    "fmt"
    "github.com/tabalt/appgo/config"
    "github.com/tabalt/appgo/logger"
)

// var (

// )

type App struct {
    ConfigFile string
    LogFile string
    Config config.Config
    Logger logger.Logger
    //...
}

func (this *App) Init() {
    fmt.Println("init appgo")
    this.Config.Init(this.ConfigFile, this.LogFile)
    this.Logger.Init(this.LogFile)
    //init other module
}
    