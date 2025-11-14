package main

import (
	"fmt"

	"mobin.dev/internal/app"
	"mobin.dev/internal/db/pgsql"
	"mobin.dev/pkg/config"
)

func main() {
	// Loading environment variable
	config.Load()

	dbPgsql, err := pgsql.Connect()

	if err != nil {
		fmt.Printf("❌ Pgsql Connection Failed : %v\n", err)
	}

	defer dbPgsql.Close()

	appInstance := app.NewApp(dbPgsql)
	appInstance.StartServer()
}
