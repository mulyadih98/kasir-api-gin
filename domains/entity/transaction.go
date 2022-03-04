package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Amount            int `json:"amount" form:"amount" binding:"required"`
	Change            int `json:"change" form:"change" binding:"required"`
	Pay               int `json:"pay" form:"pay" binding:"required"`
	UserID            uint
	TransactionDetail []TransactionDetail `json:"details"`
}

func (transaction Transaction) SavedTransaction() struct {
	TransactionID uint
	Amount        int
	Pay           int
	Change        int
} {
	return struct {
		TransactionID uint
		Amount        int
		Pay           int
		Change        int
	}{
		TransactionID: transaction.ID,
		Amount:        transaction.Amount,
		Pay:           transaction.Pay,
		Change:        transaction.Change,
	}
}
