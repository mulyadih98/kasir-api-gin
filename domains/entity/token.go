package entity

import "gorm.io/gorm"

type AuthToken struct {
	gorm.Model
	RefreshToken string `json:"refresh_token" gorm:"type:text;not null"`
	UserID       uint
	User         User
}
