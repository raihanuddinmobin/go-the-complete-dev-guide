package main

import (
	"database/sql"

	"mobin.dev/internal/app"
	"mobin.dev/pkg/config"
)

func main() {

	// Loading environment variable
	config.Load()

	appInstance := app.NewApp(&sql.DB{})
	appInstance.StartServer()
}
