package structs

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      uint
	Name        string
	Size        string
	Description string
}

func (ProductWProductExchange) TableName() string {
	return "products"
}

type ProductWProductExchange struct {
	Product
	ProductExchanges []ProductExchange `gorm:"foreignkey:ProductID;references:ID"`
}
