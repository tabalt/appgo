package appgo

import (
	"fmt"
	"github.com/tabalt/appgo/config"
	"github.com/tabalt/appgo/logger"
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


