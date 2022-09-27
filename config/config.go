package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
