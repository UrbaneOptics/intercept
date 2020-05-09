package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/config"
)

type TalliesRequest struct {
	PrecinctIDs []int `json:"precinct_ids"`
}

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

func (app *application) getTallies(w http.ResponseWriter, r *http.Request) {
	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	var t TalliesRequest
	err = dec.Decode(&t)
	// TODO: Add better error handling. See https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	if err != nil {
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	tallies, err := app.tallies.List(t.PrecinctIDs)
	if err != nil {
		app.serverError(w, err)
		return
	}

	js, err := app.fmtJSON(tallies)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sendGzipResponse(w, []byte(js))
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
