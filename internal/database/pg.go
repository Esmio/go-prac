package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "mongosteen"
	password = "123456"
	dbname   = "mongosteen_dev"
)

func PgConnect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to db")
}

func PgCreateTables() {
	// users
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully created users table")
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Migrate() {
	// 给 user 添加手机字段
	_, err := DB.Exec(`ALTER TABLE users ADD COLUMN phone VARCHAR(50)`)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Successfully added phone column to user table")
	}

	_, err = DB.Exec(`ALTER TABLE users ADD COLUMN address VARCHAR(150)`)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Successfully added address column to user table")
	}

	// 新增 Items 表，字段为 id, amount, happened_at, created_at, updated_at
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		amount INT NOT NULL,
		happened_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Successfully created items table")
	}

	_, err = DB.Exec(`ALTER TABLE items ALTER COLUMN happened_at TYPE TIMESTAMP
	`)
	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Successfully update happened_at to TIMESTAMP")
	}
}

func Crud() {
	// 创建一个 User
	_, err := DB.Exec(`INSERT INTO users (email) values ('2@qq.com')`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully created a user")
	}

	_, err = DB.Exec(`Update users set phone = '13812345678' where email = '2@qq.com'`)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully updated a user")
	}

	stmt, err := DB.Prepare(`SELECT phone FROM users where email = $1 offset $2 limit $3`)
	if err != nil {
		log.Fatalln(err)
	} else {

	}
	result, err := stmt.Query("2@qq.com", 0, 3)
	if err != nil {
		log.Println(err)
	} else {
		for result.Next() {
			var phone string
			result.Scan(&phone)
			log.Println("phone", phone)
		}
		log.Println("Successfully read users")
	}
}

func PgClose() {
	DB.Close()
	log.Println("Successfully closed db")
}
