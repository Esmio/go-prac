package router

import (
	"mongosteen/internal/controller"

	"mongosteen/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Mongosteen API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Simon
// @contact.url    https://dongxiaoming.com
// @contact.email  support@swagger.io

// @license.name  MIT
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth(JWT)

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func New() *gin.Engine {
	docs.SwaggerInfo.Version = "1.0"
	r := gin.Default()

	// Pint test
	r.GET("/api/v1/ping", controller.Ping)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
