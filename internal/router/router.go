package router

import (
	"mongosteen/internal/controller"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	// Pint test
	r.GET("/ping", controller.Ping)
	return r
}
