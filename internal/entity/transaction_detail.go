package entity

type TransactionDetail struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Product       Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"` // price per unit at time of sale
}
