package models

type Product struct {
	ProductID       uint64 `gorm:"primary_key" json:"-"`
	ProductName     string `json:"product_name"`
	ProductQuantity int    `json:"product_quantity"`
	ProductStatus   string `json:"product_status"`
	ProductPrice    string `json:"product_price"`
	Amount          string `json:"amount"`
	CustomerName    string `json:"-"`
}
