package main

import (
	docs "github.com/Thrashy190/info-center-api/docs"
	"github.com/Thrashy190/info-center-api/pkg/connection"
	"github.com/Thrashy190/info-center-api/pkg/utils"

	"github.com/Thrashy190/info-center-api/router"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializers() {
	utils.LoadEnvironment()
	connection.Dbconnection()
	connection.ModelAutoMigrations()
}

//	@Summary		Information about L2_SF APIs
//	@title			Api Rest for information center
//	@version		1.0
//	@description	This is The api for information center
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Thrashy190
//	@contact.url	https://thrashy190.github.io/
//	@contact.email	diego.lopez.mtz@protonmail.com

// @license.name				MIT
// @license.url				https://opensource.org/licenses/MIT
//
// @host						localhost:8080
// @BasePath					/
// @query.collection.format	multi
func main() {

	utils.Process("Connecting to 8080...")

	r := gin.Default()

	docs.SwaggerInfo.Title = "Information center API REST"
	docs.SwaggerInfo.Description = "This is a API REST Documentation for the ITS Information Center"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"

	utils.Succes("Connection to localhost:8080 successfull")
	initializers()

	router.SetupRouter(r, connection.DB)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
