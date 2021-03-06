package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name" binding:"required" gorm:"type:varchar(50);not null"`
	Price       int    `json:"price" form:"price" binding:"required"`
	Stock       int    `json:"stock" form:"stock" binding:"required" gorm:"not null; default:0"`
	Description string `json:"description" form:"description" `
}
