package appgo

import (
    "fmt"
    "github.com/tabalt/appgo/config"
    "github.com/tabalt/appgo/logger"
)

// var (

// )

type App struct {
    Config config.Config
    Logger logger.Logger
    //...
}

func (app *App) Init() {
    app.Config.Init("config.json")
    app.Logger.Init("log/app.log")
    //init other module
    fmt.Println("init appgo")
}
    