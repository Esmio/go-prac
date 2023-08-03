package router

import (
	"mongosteen/internal"
	"mongosteen/internal/controller"

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
	r := gin.Default()
	internal.InitRouter(r)
	docs.SwaggerInfo.Version = "1.0"

	api := r.Group("/api")

	for _, ctrl := range loadControllers() {
		ctrl.RegisterRoutes(api)
	}

	r.GET("/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
