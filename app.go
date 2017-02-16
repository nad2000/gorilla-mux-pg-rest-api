package main

import (
	"database/sql"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (app *App) Initialize(user, password, dbname string) {
	var connectionString string

	if user != "" {
		connectionString = "user=" + user
	}
	if password != "" {
		connectionString += " password=" + password
	}
	if dbname != "" {
		connectionString += "dbname=" + dbname
	}

	var err error
	app.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {}
