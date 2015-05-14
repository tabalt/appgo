package web

import (
	"io"
	"net/http"
)

//------------------------------
//type ControllerInterface
//------------------------------
type ControllerInterface interface {
}

//------------------------------
// type Controller
//------------------------------
type Controller struct {
}

func (this *Controller) Init(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "init Controller")
}

func (this *Controller) NotFound(writer http.ResponseWriter, request *http.Request) {

	// TODO set 404 not found page
	//io.WriteString(writer, "404 not found")
	io.WriteString(writer, "404 not found\n"+request.URL.String()+"\n")
}

func (this *Controller) Forbidden(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "403 Forbidden")
}

//------------------------------
// type NotFoundController
//------------------------------
type NotFoundController struct {
	Controller
}
