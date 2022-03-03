package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name  string `json:"name" form:"name" binding:"required"`
	Price int    `json:"price" form:"price" binding:"required"`
	Stock int    `json:"stock" form:"stock" binding:"required"`
}
