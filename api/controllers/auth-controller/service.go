package authController

import (
	"net/http"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"github.com/MumAroi/golang-mysql-api/api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db: db}
}

func (s *service) Login(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Prepare()

	if err := user.Validate("login"); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := s.SignIn(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (s *service) SignIn(email, password string) (string, error) {

	var err error

	var user models.User

	err = s.db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return utils.CreateToken(user.ID)
}
