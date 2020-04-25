package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db, _ = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.Ping()
}
