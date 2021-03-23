package structs

import "gorm.io/gorm"

type Follow struct {
	ID        uint `gorm:"primarykey"`
	DeletedAt gorm.DeletedAt

	UserID         uint
	FollowedUserID uint
}
