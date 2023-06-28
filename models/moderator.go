package models

import "time"

type Moderator struct {
	Id               uint   `json:"id" gorm:"primaryKey"`
	Uuid             string `json:"uuid"`
	CreatedAt        time.Time
	Username         string      `json:"username"`
	Password         string      `json:"password"`
	Email            string      `json:"email"`
	UserProfileRefer int         `json:"user_profile_id"`
	UserProfile      UserProfile `gorm:"foreignKey:UserProfileRefer"`
	MessageRefer     int         `json:"message_id"`
	Message          Message     `gorm:"foreignkey:MessageRefer"`
}
