package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5435
	user     = "mongosteen"
	password = "123456"
	dbname   = "mongosteen_dev"
)

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableAutomaticPing: true,
	})

	if err != nil {
		log.Fatal(err)
	}
	DB = db
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Successfully connected to db")
}

type User struct {
	ID        int
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func CreateTables() {
	u1 := User{Email: "1@gmail.com"}
	err := DB.Migrator().CreateTable(&u1)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Successfully created users table")
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Migrate() {

}

func Crud() {

}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Println("Successfully closed db")
}
