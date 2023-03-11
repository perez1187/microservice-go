package main

import (
	"net/http"
)

func main() {
	// we create reference to handler package
	hh := handlers.NewHello() // hh - hello  handler
	http.ListenAndServe(":9090", nil)
}
