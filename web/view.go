package web

import (
	"html/template"
	"net/http"
)

var (
	viewPath  string
	viewCache *template.Template
)

func SetViewPath(path string) {
	viewPath = path
	viewCache = template.Must(template.ParseGlob(viewPath + "*/*"))
}

// render template file
func RenderFile(file string, context *Context) {
	file = viewPath + file
	tpl, err := template.ParseFiles(file)
	if err != nil {
		// TODO deal template file not exist
		http.Error(context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(context.ResponseWriter, context.ViewData)
	return
}

// show view
func ShowView(context *Context) {
	err := viewCache.ExecuteTemplate(context.ResponseWriter, context.ViewName, context.ViewData)
	if err != nil {
		// TODO deal template file not exist
		http.Error(context.ResponseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}
