package appgo

import (
    "fmt"
    "github.com/tabalt/appgo/config"
    "github.com/tabalt/appgo/logger"
)

// var (

// )
// ui


//ConfigFile : "config.json", 
//LogFile : "log/app.log",

type App struct {
    ConfigFile string
    LogFile string
    Config config.Config
    Logger logger.Logger
    //...
}

func (this *App) Init() {
    fmt.Println("init appgo")
    fmt.Println("init config")
    //this.Config.Init(this.ConfigFile, this.LogFile)
    fmt.Println("init logger")
    //this.Logger.Init(this.LogFile)
    //init other module
}
    