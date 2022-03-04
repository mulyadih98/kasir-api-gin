package entity

import "gorm.io/gorm"

type TransactionDetail struct {
	gorm.Model
	TransactionID uint
	ProductID     uint `json:"product_id" binding:"required"`
	Product       Product
	Qty           int `json:"qty" binding:"required"`
	TotalPrice    int `json:"total_price" binding:"required"`
}
