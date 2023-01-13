package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email          string          `json:"email" gorm:"unique; not null;default:null"`
	Password       string          `json:"password"`
	Name           string          `json:"name"`
	Administrators []Administrator `json:"administrators"`
	Students       []Student       `json:"student"`
}
