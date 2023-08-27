package handlers

import (
	"github.com/husanibragimov/golang-tutorial/pkg/render"
	"net/http"
)

// HomePage /* Home is the home page handler */
func HomePage(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "home.html")
}

// AboutPage /* About is the about page handler */
func AboutPage(w http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(w, "about.html")
}
