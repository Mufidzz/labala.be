package structs

import "time"

type Token struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	ExpiredAt time.Time
	ClaimedAt *time.Time
	Token     string
}
