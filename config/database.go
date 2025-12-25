package config

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var PoolDB *sql.DB = nil

func GetConnection(username string, password string, dbName string) *sql.DB {
	err := error(nil)
	PoolDB, err = sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+dbName+"?parseTime=true")
	if err != nil {
		return nil
	}

	// Check if the connection to database is alive
	// If user inserted wrong password, username, or database name, err != nil
	err = PoolDB.Ping()
	if err != nil {
		fmt.Println("Failed to connect")
		return nil
	}

	PoolDB.SetMaxIdleConns(5)
	PoolDB.SetMaxOpenConns(10)
	PoolDB.SetConnMaxIdleTime(3 * time.Minute)
	PoolDB.SetConnMaxLifetime(60 * time.Minute)
	fmt.Println("Success make connection")
	return PoolDB
}
