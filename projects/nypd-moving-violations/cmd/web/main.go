package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"urbaneoptics.com/intercept/nypd-moving-violations/pkg/config"
)

type Config struct {
	Addr string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func openDB(host string, port int, user string, pass string, name string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	cfg := new(Config)
	dbConfig := new(config.DBConfig)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")

	// DB
	flag.StringVar(&dbConfig.Host, "dbhost", "localhost", "DB Host")
	flag.IntVar(&dbConfig.Port, "dbport", 5432, "DB Port")
	flag.StringVar(&dbConfig.User, "dbuser", "nypdmv", "DB User")
	flag.StringVar(&dbConfig.Password, "dbpass", "pass", "DB Password")
	flag.StringVar(&dbConfig.DBname, "dbname", "intercept_nypd_mv", "DB name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of app containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	db, err := openDB(
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBname,
	)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", cfg.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
