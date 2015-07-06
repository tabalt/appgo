package web

import (
	"net/http"
	"reflect"
	"strings"
)

const (
	DEFAULT_CONTROLLER_NAME = "Index"
	DEFAULT_ACTION_NAME     = "Index"
	NOTFOUND_ACTION_NAME    = "NotFound"

	PATH_TRIM_STRING  = "/ "
	PATH_SPLIT_STRING = "/"

	EMPTY_STRING = ""

	AUTO_CALL_METHOD_INIT   = "Init"
	AUTO_CALL_METHOD_BEFORE = "Before"
	AUTO_CALL_METHOD_AFTER  = "After"
	AUTO_CALL_METHOD_UNINITIALIZE = "UnInitialize"
)

var (
	ControllerManager map[string]ControllerInterface
)

func init() {
	ControllerManager = make(map[string]ControllerInterface)
}

//------------------------------
// type Router
//------------------------------
type Router struct {
	ControllerName string
	ActionName     string
}

func NewRouter() *Router {
	return &Router{}
}

// init controller name and action name from url path
func (this *Router) initControllerNameAndActionNameFromPath(path string) {
	this.ControllerName = DEFAULT_CONTROLLER_NAME
	this.ActionName = DEFAULT_ACTION_NAME

	path = strings.Trim(path, PATH_TRIM_STRING)

	if EMPTY_STRING != path {

		pathList := strings.Split(path, PATH_SPLIT_STRING)
		pathLength := len(pathList)

		if pathLength == 2 {
			this.ControllerName = strings.Title(pathList[0])
			this.ActionName = strings.Title(pathList[1])
		} else if pathLength == 1 {
			this.ControllerName = strings.Title(pathList[0])
		}

	}
}

// register controller into ControllerManager
func (this *Router) RegisterController(name string, c ControllerInterface) {
	name = strings.ToLower(name)
	ControllerManager[name] = c
}

// get controller by name from ControllerManager
func (this *Router) GetController(name string) (c ControllerInterface, ok bool) {
	name = strings.ToLower(name)
	c, ok = ControllerManager[name]
	return

}

// get current controller, return &NotFoundController{} if it not exist
func (this *Router) GetCurrentController() (c ControllerInterface) {
	c, ok := this.GetController(this.ControllerName)
	if !ok {
		c = &NotFoundController{}
	}
	return c
}

// dispatch request to matched action of controller
// call method: Init -> Before -> ActionName -> After -> UnInitialize
func (this *Router) Dispatch(w http.ResponseWriter, r *http.Request) {

	context := &Context{
		Request:        r,
		ResponseWriter: w,

		DisableView: true,
		ViewName:    strings.ToLower(this.ControllerName + "_" + this.ActionName),
		ViewData:    make(map[string]interface{}),
	}

	currentController := this.GetCurrentController()
	controllerReflect := reflect.ValueOf(currentController)

	param := []reflect.Value{reflect.ValueOf(context)}

	method := controllerReflect.MethodByName(this.ActionName)
	if (reflect.Value{} == method) {
		method = controllerReflect.MethodByName(NOTFOUND_ACTION_NAME)
	}

	// call action
	controllerReflect.MethodByName(AUTO_CALL_METHOD_INIT).Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_BEFORE).Call(param)
	method.Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_AFTER).Call(param)
	controllerReflect.MethodByName(AUTO_CALL_METHOD_UNINITIALIZE).Call(param)

	// show view
	if !context.DisableView {
		ShowView(context)
	}
}

// ServeHTTP method
func (this *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.initControllerNameAndActionNameFromPath(r.URL.Path)
	this.Dispatch(w, r)
}
