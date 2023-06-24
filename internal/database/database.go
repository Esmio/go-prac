package database

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Item struct {
	ID         int
	UserID     int
	Amount     int
	HappenedAt time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Tag struct {
	ID   int
	Name string
}

var models = []any{&User{}, &Item{}, &Tag{}}

func CreateTables() {
	for _, model := range models {
		err := DB.Migrator().CreateTable(model)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Successfully created users table")
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func Migrate() {
	DB.AutoMigrate(models...)
}

func Crud() {
	user := User{Email: "test2@qq.com"}
	tx := DB.Create(&user)
	if tx.Error != nil {
		log.Println(tx.Error)
	} else {
		log.Println(tx.RowsAffected)
	}
	user.Phone = "111111111"
	tx = DB.Save(&user)
	if tx.Error != nil {
		log.Println(tx.Error)
	}
	users := []User{}
	DB.Offset(0).Limit(3).Find(&users)

	for _, u := range users {
		log.Println(u.Phone)
	}
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.Close()
	log.Println("Successfully closed db")
}
