package handlers

import (
	"github.com/husanibragimov/golang-tutorial/pkg/render"
	"net/http"
)

/* Home is the home page handler */
func HomePage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

/* About is the about page handler */
func AboutPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
