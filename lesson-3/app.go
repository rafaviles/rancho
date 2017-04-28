// app.go

package main

import (
    "database/sql"

    "github.com/gorilla/mux"
    _ "github.com/lib/pq"
    "fmt"
    "log"
)

type App struct {
    Router *mux.Router
    DB     *sql.DB
}

func (a *App) Run(addr string) { }

func (a *App) Initialize(user, dbname string) {
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", user, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}