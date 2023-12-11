package main

import (
	"fmt"
	"net/http"

	"github.com.br/Leodf/go-web-app/pkg/handler"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	http.HandleFunc("/", handler.Home)
	http.HandleFunc("/about", handler.About)

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
