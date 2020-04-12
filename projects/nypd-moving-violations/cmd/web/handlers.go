package main

import (
	"net/http"
)

// Define a home handler function which writes a byte slice
// containing "Hello from Snippetbox" as the response body.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/"
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Hello world"))
}
