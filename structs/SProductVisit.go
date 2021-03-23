package structs

type ProductVisit struct {
	ID        uint `gorm:"primarykey"`
	ProductID uint
	UserID    uint
}
