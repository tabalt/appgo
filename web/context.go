package web

import (
	"net/http"
)

//------------------------------
// type Context
//------------------------------
type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter

	DisableView bool
	ViewName    string
	ViewData    map[string]interface{}
}

func (this *Context) AssignViewData(key string, value interface{}) {
	this.DisableView = false
	this.ViewData[key] = value
}

func (this *Context) SetViewName(name string) {
	this.DisableView = false
	this.ViewName = name
}
