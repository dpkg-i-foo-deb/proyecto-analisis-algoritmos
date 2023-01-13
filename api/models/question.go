package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Statement     string   `json:"statement"`
	Description   string   `json:"description"`
	CorrectAnswer string   `json:"correctAnswer"`
	Options       []Option `json:"option" gorm:"many2many:question_options"`
}
