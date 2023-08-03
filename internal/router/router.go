package router

import (
	"mongosteen/config"
	"mongosteen/internal/controller"
	"mongosteen/internal/database"
	"mongosteen/internal/middleware"

	"mongosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func loadControllers() []controller.Controller {
	return []controller.Controller{
		&controller.SessionController{},
		&controller.ValidationCodeController{},
	}
}

func New() *gin.Engine {
	config.LoadAppConfig()
	r := gin.Default()
	r.Use(middleware.Me())
	docs.SwaggerInfo.Version = "1.0"

	database.Connect()

	api := r.Group("/api")

	for _, ctrl := range loadControllers() {
		ctrl.RegisterRoutes(api)
	}

	r.GET("/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
