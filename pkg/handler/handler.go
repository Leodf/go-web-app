package handler

import (
	"net/http"

	"github.com.br/Leodf/go-web-app/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.tmpl")
}
