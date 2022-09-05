package models

import "time"

type Customer struct {
	ID uint64 `gorm:"primary_key" json:"-"`
	//CustomerID       uint64 `gorm:"foreignkey:" json:"customer_id"`
	CustomerName     string `json:"customer_name"`
	CustomerLocation string `json:"customer_location"`
	InvoiceDate      time.Time
	ProductID        uint64    `json:"-"`
	Product          []Product `gorm:"foreignKey:CustomerName" json:"items"`
}
