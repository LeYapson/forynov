package models

import "time"

type Message struct {
	Id                 uint `json:"id" gorm:"primaryKey"`
	CreatedAt          time.Time
	MessageContent     string  `json:"message_content"`
	Authorofthemessage string  `json:"author_of_the_message"`
	SubjectRefer       uint     `json:"subject_id"`
	Subject            Subject `gorm:"foreignKey:SubjectRefer"`
}
