package models

import "time"

type Subject struct {
	Id              uint `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time
	Subjectname     string `json:"subject_name"`
	Description     string `json:"description"`
	Messagequantity int    `json:"quantity_of_messages"`
}
