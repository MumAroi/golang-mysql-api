package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	authController "github.com/MumAroi/golang-mysql-api/api/controllers/auth-controller"
	imageController "github.com/MumAroi/golang-mysql-api/api/controllers/image-controller"
	userController "github.com/MumAroi/golang-mysql-api/api/controllers/user-controller"
	middleware "github.com/MumAroi/golang-mysql-api/api/middlewares"
)

func InitializeRoutes(db *gorm.DB, route *gin.Engine) {

	// set route controller
	authC := authController.NewService(db)
	userC := userController.NewService(db)
	imageC := imageController.NewService(db)

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello word",
		})
	})

	route.POST("/login", authC.Login)

	// sest middleware
	protected := route.Group("/", middleware.AuthorizationMiddleware)

	protected.GET("/users", userC.GetUsers)

	protected.POST("/users", userC.CreateUser)

	protected.GET("/images", imageC.GetImages)

	protected.POST("/images", imageC.CreateImage)

	protected.DELETE("/images/:id", imageC.DeleteImage)

	protected.PUT("/images/:id", imageC.UpdateImage)

	protected.GET("/images/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.File("upload/" + name)
	})
}
