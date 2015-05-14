package web

import (
	"net/http"
)

//------------------------------
// type Application
//------------------------------
type Application struct {
	Port   string
	Router *Router
}

func NewApplication(port string) *Application {

	router := NewRouter()

	return &Application{Port: port, Router: router}
}

func (this *Application) SetPort(port string) {
	this.Port = port
}

func (this *Application) RegisterController(name string, c ControllerInterface) {
	this.Router.RegisterController(name, c)
}

func (this *Application) Run() {
	http.ListenAndServe(this.Port, this.Router)
}
