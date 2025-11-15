package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"mobin.dev/pkg/config"
)

func Connect() (*sql.DB, error) {
	cnf := config.Get()

	mysqlConnectionStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cnf.MYSQL_User, cnf.MYSQL_Password, cnf.MYSQL_Host, cnf.MYSQL_Port, cnf.MYSQL_DbName)

	db, err := sql.Open("mysql", mysqlConnectionStr)
	db.SetConnMaxIdleTime(time.Second * 10)
	db.SetMaxIdleConns(50)
	db.SetMaxIdleConns(5)

	if err != nil {
		return nil, fmt.Errorf("failed to connect mysql, reason : %w ", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping mysql, reason : %w ", err)
	}

	fmt.Println("Successfully Connected mysql ✅")
	return db, nil
}
