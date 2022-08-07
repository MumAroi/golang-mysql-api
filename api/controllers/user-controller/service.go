package userController

import (
	"net/http"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db: db}
}

func (s *service) CreateUser(c *gin.Context) {

	var user models.User

	db := s.db.Model(&user)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Prepare()

	if err := user.Validate(""); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}
