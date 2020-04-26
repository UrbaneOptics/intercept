package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/config"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/"
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	w.Write([]byte("Hello world"))
}

func (app *application) showPrecinct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	p, err := app.precincts.Get(id)
	if err != nil {
		app.serverError(w, err)
	}

	js, err := json.Marshal(p)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(config.HealthStatus{Status: "OK"})
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
