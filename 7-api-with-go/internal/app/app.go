package app

import "database/sql"

type App struct {
	pg *sql.DB
}

func NewApp(pg *sql.DB) *App {
	return &App{pg: pg}
}
