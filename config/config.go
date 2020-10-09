package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"log"
)

// ConnectMysql is a function that used to connect to mysql
func ConnectMysql() *sql.DB {
	user := "root"
	password := "Buk@dong11"
	host := "127.0.0.1"
	port := "3306"
	database := "api_openfeeder"

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)

	db, err := sql.Open("mysql", connection)

	if err != nil {
		log.Fatal(err)
	}

	return db

}
