package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string    `gorm:"uniqueIndex;not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"default:user"` // user or admin
	Profile      Profile   `gorm:"constraint:OnDelete:CASCADE"`
	Meetings     []Meeting `gorm:"foreignKey:CreatorID"`
}

type Profile struct {
	gorm.Model
	UserID    uint `gorm:"uniqueIndex"`
	FullName  string
	Bio       string
	AvatarURL string
	Location  string
	Website   string
}
