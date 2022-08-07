package seeds

import (
	"log"

	"github.com/MumAroi/golang-mysql-api/api/models"
	"gorm.io/gorm"
)

var users = []models.User{
	models.User{
		Nickname: "Paramas Waewsuwan",
		Email:    "paramas.test@gmail.com",
		Password: "123456",
	},
}

func Load(db *gorm.DB) {

	var err error
	for i, _ := range users {
		err = db.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
