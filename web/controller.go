package web

import (
	"html/template"
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

// render template
// TODO deal template file not exist
func (this *Controller) Render(w http.ResponseWriter, tplName string, context map[string]interface{}) {
	tpl, err := template.ParseFiles(tplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, context)
	return
}

func (this *Controller) Init(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "init Controller")
}

func (this *Controller) NotFound(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "404 not found", http.StatusNotFound)
}

//------------------------------
// type NotFoundController
//------------------------------
type NotFoundController struct {
	Controller
}
