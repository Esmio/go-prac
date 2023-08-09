package main

import (
	"mongosteen/cmd"
	"mongosteen/config"
)

//	@title			Mongosteen API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Simon
//	@contact.url	https://dongxiaoming.com
//	@contact.email	support@swagger.io

//	@license.name	MIT
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apiKey	Bearer
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	config.LoadAppConfig()
	cmd.Run()
}
