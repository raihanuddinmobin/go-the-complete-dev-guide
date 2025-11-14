package main

import (
	"database/sql"

	"mobin.dev/internal/app"
)

func main() {

	appInstance := app.NewApp(&sql.DB{})
	appInstance.StartServer()
}
