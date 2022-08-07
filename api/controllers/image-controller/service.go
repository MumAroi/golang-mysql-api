package imageController

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"github.com/MumAroi/golang-mysql-api/api/utils"
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

func (s *service) CreateImage(c *gin.Context) {

	var image models.Image

	db := s.db.Model(&image)

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

	imageURL := "http://localhost:" + utils.GodotEnv("GO_PORT") + "/images/" + newFileName

	c.Bind(&image)

	image.Url = imageURL

	image.Prepare()

	if err := image.Validate(""); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := db.Create(&image); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messs": "Can not create image",
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, image)

}

func (s *service) GetImages(c *gin.Context) {

	var image models.Image

	images, err := image.FindAllImage(s.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"messs": "Can not get images",
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, images)

}

func (s *service) UpdateImage(c *gin.Context) {

	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"messs": "Not found id",
			"error": err.Error(),
		})
		return
	}

	var image models.Image

	c.Bind(&image)

	image.Prepare()

	if err := image.Validate("update"); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"meaage": "Data invalid",
			"error":  err.Error(),
		})
		return
	}

	imageUpdated, err := image.UpdateImage(s.db, uid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"meaage": "Update image fail",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, imageUpdated)

}

func (s *service) DeleteImage(c *gin.Context) {

	var image models.Image

	id := c.Param("id")
	uid, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"messs": "Not found id",
			"error": err.Error(),
		})
		return
	}

	_, err = image.DeleteImage(s.db, uid)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"meaage": "Update image fail",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusNoContent, "")
}
