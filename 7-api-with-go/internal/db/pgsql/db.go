package pgsql

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"mobin.dev/pkg/config"
)

func Connect() (*sql.DB, error) {

	cnf := config.Get()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cnf.PGHost, cnf.PGPort, cnf.PGUser, cnf.PGPassword, cnf.PGDbName)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("failed to connect db, reason :%w ", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db, reason :%w ", err)
	}

	fmt.Println("Successfully Connected PgSql ✅")
	return db, nil
}
