package handlers

import (
	"github.com/husanibragimov/golang-tutorial/pkg/config"
	"github.com/husanibragimov/golang-tutorial/pkg/models"
	"github.com/husanibragimov/golang-tutorial/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
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

// HomePage /* Home is the home page handler */
func (m *Repository) HomePage(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.html", &models.TemplateData{})
}

// AboutPage /* About is the about page handler */
func (m *Repository) AboutPage(w http.ResponseWriter, _ *http.Request) {
	// perform some logic

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.RenderTemplate(w, "about.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
