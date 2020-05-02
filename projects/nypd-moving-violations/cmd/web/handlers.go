package main

import (
	"fmt"
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
	err := app.handleInvalidRequest(w, r)
	if err != nil {
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

	js, err := app.fmtJSON(p)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

func (app *application) getPrecincts(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	pcts, err := app.precincts.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	js, err := app.fmtJSON(pcts)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

func (app *application) showTally(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	p, err := app.tallies.Get(id)
	if err != nil {
		app.serverError(w, err)
	}

	js, err := app.fmtJSON(p)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

// TODO: Implement with filter logic
func (app *application) getTallies(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	tallies, err := app.tallies.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	js, err := app.fmtJSON(tallies)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

func (app *application) showMovingViolation(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	p, err := app.movingViolations.Get(id)
	if err != nil {
		app.serverError(w, err)
	}

	js, err := app.fmtJSON(p)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

func (app *application) getMovingViolations(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	movingViolations, err := app.movingViolations.List()
	if err != nil {
		app.serverError(w, err)
		return
	}

	js, err := app.fmtJSON(movingViolations)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	js, err := app.fmtJSON(config.HealthStatus{Status: "OK"})
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))
}
