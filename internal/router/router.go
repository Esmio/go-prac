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

	api := r.Group("/api")

	sc := controller.SessionController{}
	sc.RegisterRoutes(api)

	vcc := controller.ValidationCodeController{}
	vcc.RegisterRoutes(api)

	r.GET("/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
