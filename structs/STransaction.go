package structs

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ProductID         uint
	ProductExchangeID uint
	Amount            uint
	Status            uint
}
