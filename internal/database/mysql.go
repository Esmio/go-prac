package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlHost     = "localhost"
	mysqlPort     = 3307
	mysqlUser     = "mongosteen"
	mysqlPassword = "123456"
	mysqlDbname   = "mongosteen_dev"
)

func MysqlConnect() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:3307)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlDbname)
	fmt.Println(connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully connected to db")
}

func MysqlClose() {
	DB.Close()
	log.Println("Successfully closed db")
}
