package cmd

import (
	"log"
	"mongosteen/internal/database"
	"mongosteen/internal/router"
)

func RunServer() {
	database.Connect()
	database.CreateTables()
	defer database.Close()
	r := router.New()
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("r.Run next line")
}
