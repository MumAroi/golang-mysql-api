package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	userController "github.com/MumAroi/golang-mysql-api/api/controllers/user-controller"
)

func InitializeRoutes(db *gorm.DB, route *gin.Engine) {

	userC := userController.NewRepository(db)

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello word",
		})
	})

	route.POST("/user", userC.CreateUser)

}
