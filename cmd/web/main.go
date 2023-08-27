package main

import (
	"fmt"
	"github.com/husanibragimov/golang-tutorial/pkg/config"
	"github.com/husanibragimov/golang-tutorial/pkg/handlers"
	"github.com/husanibragimov/golang-tutorial/pkg/render"
	"log"
	"net/http"
	"time"
)

var port = ":8080"

// Main is the main application func
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/about", handlers.AboutPage)

	fmt.Println(fmt.Sprintf("Starting aplication on port %s", port))
	fmt.Println("RFC1123:", time.Now().Format(time.RFC1123))
	fmt.Println(fmt.Sprintf("Starting development server at http://localhost%s/", port))
	fmt.Println("Quit the server with CONTROL-C. \n")
	_ = http.ListenAndServe(port, nil)
}
