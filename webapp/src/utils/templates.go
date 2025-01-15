package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

func RenderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	templates.ExecuteTemplate(w, templateName, data)
}
