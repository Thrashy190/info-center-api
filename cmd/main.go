package main

import (
	"github.com/Thrashy190/info-center-api/pkg/connection"
	"github.com/Thrashy190/info-center-api/pkg/utils"
	"github.com/Thrashy190/info-center-api/router"
	"github.com/gin-gonic/gin"
)

func initializers() {
	utils.LoadEnvironment()
	connection.Dbconnection()
	connection.ModelAutoMigrations()
}

func main() {

	utils.Process("Connecting to 8080...")
	r := gin.Default()
	utils.Succes("Connection to localhost:8080 successfull")
	initializers()

	router.SetupRouter(r, connection.DB)

	r.Run(":8080")
}
