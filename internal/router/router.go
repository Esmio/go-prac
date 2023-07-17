package router

import (
	"mongosteen/config"
	"mongosteen/internal/controller"
	"mongosteen/internal/database"

	"mongosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func New() *gin.Engine {
	config.LoadAppConfig()
	r := gin.Default()
	docs.SwaggerInfo.Version = "1.0"

	database.Connect()

	// Pint test
	r.GET("/api/v1/ping", controller.Ping)
	r.POST("/api/v1/validation_codes", controller.CreateValidationCode)
	r.POST("/api/v1/session", controller.CreateSession)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
