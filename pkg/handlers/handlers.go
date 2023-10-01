package handlers

import (
	"go-web-app/pkg/config"
	"go-web-app/pkg/models"
	"go-web-app/pkg/render"
	"net/http"
)

// Repository used by the handlers
var Repo *Repository

// Repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	/* Create a new repository */
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	/* Set the repository for the handlers */
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about  page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
