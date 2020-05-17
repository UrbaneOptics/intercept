package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/config"
	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/queries"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/"
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	err := app.handleInvalidRequest(w, r)
	if err != nil {
		return
	}

	// w.Write([]byte("Hello world"))
	var routes []string
	fmtRoutes := append(routes, "/precinct",
		"/precincts",
		"/tally",
		"/tallies",
		"/moving_violation",
		"/moving_violations",
		"/health")
	js, err := app.fmtJSON(fmtRoutes)
	if err != nil {
		app.serverError(w, err)
		return
	}

	fmt.Fprint(w, string(js))

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
		if err == sql.ErrNoRows {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
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
		if err == sql.ErrNoRows {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
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
	t := queries.NewTalliesRequest()
	err = dec.Decode(&t)
	// TODO: Add better error handling. See https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
	if err != nil {
		msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", err)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Check whether the required parameters are missing
	if t.PrecinctIDs == nil {
		msg := fmt.Sprintf("Request missing required query parameter 'precinct_ids'")
		app.customError(w, msg, http.StatusBadRequest)
		return
	}

	if t.MovingViolationIDs == nil {
		msg := fmt.Sprintf("Request missing required query parameter 'moving_violation_ids'")
		app.customError(w, msg, http.StatusBadRequest)
		return
	}

	// Generate list of found tallies
	tallies, err := app.tallies.List(&t)
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
		if err == sql.ErrNoRows {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
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
