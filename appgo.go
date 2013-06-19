package appgo

import (
	"github.com/tabalt/appgo/config"
    "fmt"
)

// var (

// )

type AppGo struct {
    Config config.Config

    //...
}

func (app *AppGo) Init() {
    app.Config.Init("config.json")
    fmt.Println("init appgo")    
}
