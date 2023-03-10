package main

import (
	"fmt"
	"github.com/rach4you/booking/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

// main is the entry point of an application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
