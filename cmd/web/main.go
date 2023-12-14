package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com.br/Leodf/go-web-app/pkg/config"
	"github.com.br/Leodf/go-web-app/pkg/handler"
	"github.com.br/Leodf/go-web-app/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app)
	handler.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/about", handler.Repo.About)

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
