package database

import (
	"exams-api/models"
	"exams-api/util"
)

func migrateModels() {

	err := DB.SetupJoinTable(&models.Exam{}, "Students", &models.ExamPresentations{})

	util.HandleErrorStop(err)

	err = DB.AutoMigrate(
		&models.User{},
		&models.Administrator{},
		&models.Student{},
		&models.Question{},
		&models.Option{},
		&models.Exam{},
		&models.ExamPresentations{},
	)

	util.HandleErrorStop(err)
}
