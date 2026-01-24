package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache


	// get requested template from ache


	// render the template



	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err := parsedTemplate.Execute(w, nil); err != nil {
		fmt.Println("error parsing templates:", err.Error())
		return
	}
}


func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./template/*.page.tmpl")
	if err != nil {return myCache, err}

	for _, page := range pages {
		println(page)
	}

	
	return myCache, nil
}

