package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/precinct", app.showPrecinct)
	mux.HandleFunc("/health", app.health)
	return mux
}
