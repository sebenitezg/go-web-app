package main

import (
	"fmt"
	"go-web-app/pkg/handlers"
	"net/http"
)

const portNumber = ":8000"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
