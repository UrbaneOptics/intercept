package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
)

type ErrorJSON struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func JSONError(w http.ResponseWriter, err string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorJSON{Message: err, Code: code})
}

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 internal Server Error repsonse to the user.
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
func (app *application) clientError(w http.ResponseWriter, status int) {
	JSONError(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) fmtJSON(d interface{}) ([]byte, error) {
	js, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// Write a response compressed using GZIP
func (app *application) sendGzipResponse(w http.ResponseWriter, res []byte) error {
	w.Header().Set("Content-Encoding", "gzip")
	zw := gzip.NewWriter(w)
	_, err := zw.Write(res)
	if err != nil {
		return err
	}
	if err := zw.Close(); err != nil {
		app.serverError(w, err)
		return err
	}
	return nil
}

func (app *application) handleInvalidRequest(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		app.clientError(w, http.StatusMethodNotAllowed)
		return fmt.Errorf("Invalid status method")
	}
	return nil
}
