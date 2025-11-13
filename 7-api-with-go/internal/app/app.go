package app

import (
	"database/sql"
)

type App struct {
	mysql *sql.DB
}

func NewApp(mysql *sql.DB) *App {
	return &App{mysql}
}
