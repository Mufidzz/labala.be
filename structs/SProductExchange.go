package structs

import "gorm.io/gorm"

type ProductExchange struct {
	gorm.Model
	Name      string
	Amount    uint
	Size      string
	ProductID uint
}
