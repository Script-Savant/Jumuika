package models

import "gorm.io/gorm"

type Friendship struct {
	gorm.Model
	FollowerID uint `gorm:"index;not null"`
	FollowedID uint `gorm:"index;not null"`

	Follower User `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE"`
	Followed User `gorm:"foreignKey:FollowedID;constraint:OnDelete:CASCADE"`
}
