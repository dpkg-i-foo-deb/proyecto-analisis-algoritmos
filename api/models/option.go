package models

import "gorm.io/gorm"

type Option struct {
	gorm.Model
	Order       int16  `json:"order"`
	Description string `json:"description"`
	Picture     []byte `json:"picture"`
	QuestionID  uint   `json:"question"`
}
