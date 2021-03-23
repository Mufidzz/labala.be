package structs

import "gorm.io/gorm"

type Like struct {
	ID        uint `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt

	UserID    uint
	ProductID uint
}
