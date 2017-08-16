package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(dbname, user, password, host string, port int) {}

func (a *App) Run(addr string) {}
