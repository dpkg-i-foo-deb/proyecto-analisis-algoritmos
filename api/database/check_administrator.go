package database

import (
	"exams-api/models"
	"exams-api/util"
	"log"
	"os"

	"gorm.io/gorm"
)

func checkAdministrator() {
	var admin models.User

	res := DB.First(&admin)

	if res.Error == gorm.ErrRecordNotFound {
		log.Println("The default administrator account does not exist")

		pass, err := util.HashPassword(os.Getenv("ADMINISTRATOR_PASSWORD"))

		util.HandleErrorStop(err)

		admin = models.User{
			Model:    gorm.Model{},
			Email:    os.Getenv("ADMINISTRATOR_EMAIL"),
			Password: pass,
			Name:     "Administrator",
			Type:     string(models.ADMINISTRATOR),
		}

		res = DB.Create(&admin)

		util.HandleErrorStop(res.Error)

		log.Println("The default administrator has been created")
		log.Println("Email: " + admin.Email)
		log.Println("Password: " + os.Getenv("ADMINISTRATOR_PASSWORD"))
	} else {
		util.HandleErrorStop(res.Error)
	}

}
