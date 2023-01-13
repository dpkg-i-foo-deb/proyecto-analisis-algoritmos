package database

import (
	"exams-api/util"

	"github.com/glebarez/sqlite"

	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	var err error

	DB, err = gorm.Open(sqlite.Open("exams.db"), &gorm.Config{})

	util.HandleErrorStop(err)
}
