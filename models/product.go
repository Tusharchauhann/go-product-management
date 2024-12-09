package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	UserID             uint           `json:"user_id"`
	ProductName        string         `json:"product_name"`
	ProductDescription string         `json:"product_description"`
	ProductPrice       float64        `json:"product_price"`
	ProductImages      pq.StringArray `gorm:"type:text[]" json:"product_images"`
}

// TableName function to specify the correct table name
func (Product) TableName() string {
	return "productm" // Correct table name for products
}
