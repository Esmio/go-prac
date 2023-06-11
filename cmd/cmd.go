package cmd

import "mongosteen/internal/router"

func RunServer() {
	r := router.New()
	r.Run(":8080")
}
