package web

import (
	"html/template"
	"net/http"
)

var (
	ViewPath  string
	ViewParse *template.Template
)

func init() {

}

func SetViewPath(path string) {
	ViewPath = path
	ViewParse = template.Must(template.ParseGlob(_view_path + "*/*"))
}

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
	tplFile := ViewPath + tplName
	tpl, err := template.ParseFiles(tplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, context)
	return
}

// show view
// TODO deal template file not exist
func (this *Controller) ShowView(w http.ResponseWriter, view_name string, context map[string]interface{}) {
	err := ViewParse.ExecuteTemplate(w, view_name, context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


func (this *Controller) NotFound(writer http.ResponseWriter, request *http.Request) {
	http.Error(writer, "404 not found", http.StatusNotFound)
}

// TODO init controller
func (this *Controller) Init() {

}

// Before runs before request function execution.
func (this *Controller) Before(writer http.ResponseWriter, request *http.Request) {

}

// After runs after request function execution.
func (this *Controller) After(writer http.ResponseWriter, request *http.Request) {

}
//------------------------------
// type NotFoundController
//------------------------------
type NotFoundController struct {
	Controller
}
