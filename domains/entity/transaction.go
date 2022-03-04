package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount int `json:"amount" form:"amount" binding:"required"`
	Change int `json:"change" form:"change" binding:"required"`
	UserID uint
}
