package configs

import (
	"fmt"
	"os"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"github.com/MumAroi/golang-mysql-api/api/seeds"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {
	databaseURI := make(chan string, 1)

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	databaseURI <- DBURL

	db, err := gorm.Open(mysql.Open(<-databaseURI), &gorm.Config{})

	if err != nil {
		defer logrus.Info("Connection to Database Failed")
		logrus.Fatal(err.Error())
	}

	logrus.Info("Connection to Database Successfully")

	// migrate table
	db.AutoMigrate(&models.User{}, &models.Image{})

	// seed data
	seeds.Load(db)

	if err != nil {
		logrus.Fatal(err.Error())
	}

	return db
}
