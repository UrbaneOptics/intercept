package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
}

// Define an application struct to hold the application-wide dependencies for the web app.
// For now we'll only include fields for the two custom loggers,
// but we'll add more to it as the build progresses
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	cfg := new(Config)

	// Allow setting an address. Default to 4000
	// ie. go run ./cmd/web -addr=":9999"
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")

	// Parse command line flags. Must be called before using a flag var.
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of app containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
