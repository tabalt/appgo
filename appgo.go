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
	Config     *config.Config
	Logger     *logger.Logger
)

type App struct {

	//...
}

func init() {
	fmt.Println("init appgo")
	ConfigPath = "conf/config.json"
	LogPath = "log/app.log"
	Config = config.Instance(ConfigPath)
	Logger = logger.Instance(LogPath)
}

type Base struct {
}
