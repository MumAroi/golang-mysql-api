package imageController

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{db: db}
}

func (r *service) CreateImage(c *gin.Context) {

	var image models.Image

	db := r.db.Model(&image)

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "The file cannot be received",
			"error":   err.Error(),
		})
		return
	}

	extension := filepath.Ext(file.Filename)

	newFileName := uuid.New().String() + extension

	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		if err := os.Mkdir("upload", os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	path, err := filepath.Abs("upload/" + newFileName)

	if err != nil {

		log.Fatal(err)
	}

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "The file is received, so let's save it",
			"error":   err.Error(),
		})
		return
	}

	imageURL := "http://localhost:8080/images/" + newFileName

	c.Bind(&image)

	image.Url = imageURL

	if result := db.Create(&image); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messs": "Can not create image",
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, image)

}
