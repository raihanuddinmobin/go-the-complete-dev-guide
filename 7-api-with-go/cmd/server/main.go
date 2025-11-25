package main

import (
	"fmt"

	"mobin.dev/internal/app"
	"mobin.dev/internal/db/mongo"
	"mobin.dev/internal/db/mysql"
	"mobin.dev/internal/db/pgsql"
	"mobin.dev/internal/db/redis"
	"mobin.dev/pkg/config"
)

func main() {
	// Loading environment variable
	config.Load()

	dbPgsql, errPgsql := pgsql.Connect()
	dbMysql, errMysql := mysql.Connect()
	_, errMongo := mongo.Connect()
	dbRedis, errRedis := redis.Connect()

	if errPgsql != nil {
		fmt.Printf("❌ Pgsql Connection Failed : %v\n", errPgsql)
	}

	if errMysql != nil {
		fmt.Printf("❌ Mysql Connection Failed : %v\n", errMysql)
	}

	if errMongo != nil {
		fmt.Printf("❌ Mongo Connection Failed : %v\n", errMongo)
	}

	if errRedis != nil {
		fmt.Printf("❌ Redis Connection Failed : %v\n", errMongo)
	}

	defer dbPgsql.Close()
	defer dbMysql.Close()
	defer mongo.Disconnect()
	defer dbRedis.Close()

	appInstance := app.NewApp(dbPgsql, dbRedis)
	appInstance.StartServer()
}
