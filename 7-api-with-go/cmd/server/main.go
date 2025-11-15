package main

import (
	"fmt"

	"mobin.dev/internal/app"
	"mobin.dev/internal/db/mongo"
	"mobin.dev/internal/db/pgsql"
	"mobin.dev/pkg/config"
)

func main() {
	// Loading environment variable
	config.Load()

	dbPgsql, errPgsql := pgsql.Connect()
	dbMongo, errMongo := mongo.Connect()

	if errPgsql != nil {
		fmt.Printf("❌ Pgsql Connection Failed : %v\n", errPgsql)
	}

	if errMongo != nil {
		fmt.Printf("❌ Mongo Connection Failed : %v\n", errMongo)
	}

	fmt.Println(dbMongo)
	defer dbPgsql.Close()

	appInstance := app.NewApp(dbPgsql)
	appInstance.StartServer()
}
