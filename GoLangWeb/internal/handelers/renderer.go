package handelers

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("error parsing templates:", err.Error())
		return
	}
}
