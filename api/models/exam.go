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
	Students    []Student     `json:"students" gorm:"many2many:exam_presentations"`
}

type ExamPresentations struct {
	ExamID    uint          `json:"exam" gorm:"primaryKey"`
	StudentID uint          `json:"student" gorm:"primaryKey"`
	Grade     float32       `json:"grade"`
	Time      time.Duration `json:"time"`
}
