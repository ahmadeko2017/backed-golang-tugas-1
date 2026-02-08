package entity

import "time"

type Transaction struct {
	ID        uint                `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time           `json:"created_at"`
	Total     float64             `json:"total"`
	Details   []TransactionDetail `gorm:"foreignKey:TransactionID" json:"details"`
}
