package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	authController "github.com/MumAroi/golang-mysql-api/api/controllers/auth-controller"
	imageController "github.com/MumAroi/golang-mysql-api/api/controllers/image-controller"
	userController "github.com/MumAroi/golang-mysql-api/api/controllers/user-controller"
)

func InitializeRoutes(db *gorm.DB, route *gin.Engine) {

	authC := authController.NewService(db)
	userC := userController.NewService(db)
	imageC := imageController.NewService(db)

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello word",
		})
	})

	route.POST("/login", authC.Login)

	route.POST("/users", userC.CreateUser)

	route.GET("/images", imageC.GetImages)

	route.POST("/images", imageC.CreateImage)

	route.DELETE("/images/:id", imageC.DeleteImage)

	route.PUT("/images/:id", imageC.UpdateImage)

	route.GET("/images/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.File("upload/" + name)
	})
}
