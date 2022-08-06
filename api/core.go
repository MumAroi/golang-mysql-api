package api

import (
	"log"

	"github.com/MumAroi/golang-mysql-api/api/configs"
	"github.com/MumAroi/golang-mysql-api/api/routes"
	"github.com/MumAroi/golang-mysql-api/api/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func Run() {
	//  set up route
	router := SetupRouter()

	// start server
	log.Fatal(router.Run(":" + utils.GodotEnv("GO_PORT")))

}

func SetupRouter() *gin.Engine {

	//  load db connection
	db := configs.Connection()

	router := gin.Default()

	gin.SetMode(gin.DebugMode)

	// load routes
	routes.InitializeRoutes(db, router)

	return router
}
