package databases

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/lowcode")

	if err != nil {
		log.Panicln("err:", err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
