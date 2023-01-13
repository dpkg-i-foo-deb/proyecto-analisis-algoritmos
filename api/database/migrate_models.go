package database

import (
	"exams-api/models"
	"exams-api/util"
)

func migrateModels() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Administrator{},
		&models.Student{},
		&models.Question{},
		&models.Option{},
	)

	util.HandleErrorStop(err)
}
