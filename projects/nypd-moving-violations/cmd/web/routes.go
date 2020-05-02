package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/precinct", http.HandlerFunc(app.showPrecinct))
	mux.HandleFunc("/precincts", app.getPrecincts)

	mux.HandleFunc("/tally", app.showTally)
	mux.HandleFunc("/tallies", app.getTallies)

	mux.HandleFunc("/moving_violation", app.showMovingViolation)
	mux.HandleFunc("/moving_violations", app.getMovingViolations)

	mux.HandleFunc("/health", app.health)
	return baseHeaders(mux)
}
