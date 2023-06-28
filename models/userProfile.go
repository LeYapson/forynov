package models

import "time"

type UserProfile struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UserProfile string `json:"user_profile"`
	Biography   string `json:"biography"`
	Profilepic  string `json:"profile_pic"`
}
