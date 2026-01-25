package handler

import (
	"myApp/pkg/config"
	"myApp/pkg/http/render"
	"myApp/pkg/models"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repositry type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// list, view, create
func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// Performing some logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, from Home Handler"

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
