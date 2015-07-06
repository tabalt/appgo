package web

import (
	"net/http"
)

//------------------------------
//type ControllerInterface
//------------------------------
type ControllerInterface interface {
	Init(context *Context)
	Before(context *Context)
	NotFound(context *Context)
	After(context *Context)
	UnInitialize(context *Context)
}

//------------------------------
// type Controller
//------------------------------
type Controller struct {
}

// TODO init controller
func (this *Controller) Init(context *Context) {

}

// Before runs before request function execution.
func (this *Controller) Before(context *Context) {

}

// not found action
func (this *Controller) NotFound(context *Context) {
	context.DisableView = true
	http.Error(context.ResponseWriter, "404 not found", http.StatusNotFound)
}
// After runs after request function execution.
func (this *Controller) After(context *Context) {

}

// TODO uninitialize controller
func (this *Controller) UnInitialize(context *Context) {

}
//------------------------------
// type NotFoundController
//------------------------------
type NotFoundController struct {
	Controller
}
