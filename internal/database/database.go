package database

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"mongosteen/config/queries"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sql.DB
var DBCtx = context.Background()

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
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	DB = db
	err = DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
}

func NewQuery() *queries.Queries {
	return queries.New(DB)
}

func CreateMigration(filename string) {
	cmd := exec.Command("migrate", "create", "-ext", "sql", "-dir", "config/migrations", "-seq", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
}

func Migrate() {
	dir, err := os.Getwd()
	name := filepath.Base(dir)
	for !strings.Contains(name, "mongosteen") {
		dir = filepath.Dir(dir)
		name = filepath.Base(dir)
	}
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			log.Fatalln(err)
		}
	}
}

func MigrateDown() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	m, err := migrate.New(
		fmt.Sprintf("file://%s/config/migrations", dir),
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalln(err)
	}
	err = m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	}
}

func Crud() {
	q := queries.New(DB)
	id := rand.Int()
	u, err := q.CreateUser(DBCtx, fmt.Sprintf("%d@qq.com", id))
	if err != nil {
		log.Fatalln(err)
	}
	err = q.UpdateUser(DBCtx, queries.UpdateUserParams{
		ID:      u.ID,
		Email:   u.Email,
		Phone:   u.Phone,
		Address: "Tokyo",
	})
	if err != nil {
		log.Fatalln(err)
	}

	u, err = q.FindUserByEmail(DBCtx, fmt.Sprintf("%d@qq.com", id))
	if err != nil {
		log.Fatalln(err)
	}
}

func Close() {

}
