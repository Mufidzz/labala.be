package structs

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserDefinedID string `gorm:"unique"`
	Email         string `gorm:"unique"`
	Password      string

	FullName string
	Bio      string
	Whatsapp string

	Access          uint `gorm:"-"`
	PasswordTokenID uint
	RegisterTokenID uint
	PrivilegesLevel uint `gorm:"default:0"`
}
