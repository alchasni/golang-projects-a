package mysql

import "time"

type User struct {
	ID        uint64    `gorm:"primaryKey"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null"`
	AvatarUrl string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
	//DeletedAt time.Time
}
