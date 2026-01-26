package main

import (
	"log"
	"myApp/pkg"
	"myApp/pkg/config"
	"myApp/pkg/http/handler"
	"myApp/pkg/http/render"
	"myApp/pkg/templates"
	"net/http"
)

func main() {

	var app config.AppConfig

	ourTemplateCache, err := render.CreateTemplateCache(templates.TemplatesFS)
	if err != nil {
		log.Fatal("cannot create a template cache")
	}
	app.TemplateCache = ourTemplateCache
	app.UseCache = false
	render.NewTemplate(&app)

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	log.Println("Server is running on Port", pkg.PORT)

	serve := &http.Server{
		Addr:    pkg.PORT,
		Handler: Router(&app),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
