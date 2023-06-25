package router

import (
	"mongosteen/internal/controller"

	"mongosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	docs.SwaggerInfo.Version = "1.0"
	r := gin.Default()

	// Pint test
	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
