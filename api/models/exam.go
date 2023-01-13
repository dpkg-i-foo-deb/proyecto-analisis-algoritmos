package models

import (
	"time"

	"gorm.io/gorm"
)

type Exam struct {
	gorm.Model
	Points      int           `json:"points"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	LimitTime   time.Duration `json:"limitTime"`
	Questions   []Question    `json:"question" gorm:"many2many:exam_questions"`
}
