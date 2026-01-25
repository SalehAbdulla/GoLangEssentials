package render

import (
	"bytes"
	"html/template"
	"io/fs"
	"log"
	"myApp/pkg/config"
	"myApp/pkg/models"
	"myApp/templates"
	"net/http"
	"path"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var templateCache map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache(templates.TemplatesFS)
	}

	// get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Println("Could not get template from cache")
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	buffer := new(bytes.Buffer)
	err := template.Execute(buffer, templateData)
	if err != nil {
		log.Println(err)
		http.Error(w, "Render error", http.StatusInternalServerError)
		return
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		buffer := new(bytes.Buffer)
		err := template.Execute(buffer, templateData)
		if err != nil {
			log.Println(err)
			http.Error(w, "Render error", http.StatusInternalServerError)
			return
		}

		_, err = buffer.WriteTo(w)
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(err)
	}
}

func CreateTemplateCache(templateFS fs.FS) (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	// embedded paths always use forward slashes
	pages, err := fs.Glob(templateFS, "*.page.tmpl")
	if err != nil {
		return cache, err
	}

	layouts, err := fs.Glob(templateFS, "*.layout.tmpl")
	if err != nil {
		return cache, err
	}

	for _, pageFile := range pages {
		name := path.Base(pageFile)

		// parse page + all layouts
		files := append([]string{pageFile}, layouts...)
		templateParsed, err := template.ParseFS(templateFS, files...)
		if err != nil {
			return cache, err
		}

		cache[name] = templateParsed
	}

	return cache, nil
}
