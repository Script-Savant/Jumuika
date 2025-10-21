package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	Meetings    []Meeting `gorm:"foreignKey:CategoryID"`
}

type Location struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Address   string
	City      string `gorm:"index"`
	Country   string
	Latitude  float64
	Longitude float64
	Meetings  []Meeting `gorm:"foreignKey:LocationID"`
}

type Meeting struct {
	gorm.Model
	Title        string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	StartTime    time.Time `gorm:"not null"`
	EndTime      time.Time
	CreatorID    uint `gorm:"index;not null"`
	LocationID   uint
	CategoryID   uint
	IsOnline     bool
	MeetingLink  string
	MaxAttendees uint
	ImageURl     string
	Paid         bool `gorm:"default:false"`
	Price        float64

	Creator   User      `gorm:"foreignKey:CreatorID;constraint:OnDelete:CASCADE"`
	Location  Location  `gorm:"constraint:OnDelete:SET NULL"`
	Category  Category  `gorm:"constraint:OnDelete:SET NULL"`
	Attendees []RSVP    `gorm:"foreignKey:MeetingID"`
	Comments  []Comment `gorm:"foreignKey:MeetingID"`
}

type RSVP struct {
	gorm.Model
	UserID    uint   `gorm:"index;not null"`
	MeetingID uint   `gorm:"index;not null"`
	Status    string // attending, interested

	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Meeting Meeting `gorm:"foreignKey:MeetingID;constraint:OnDelete:CASCADE"`
}

type Comment struct {
	gorm.Model
	UserID    uint   `gorm:"index;not null"`
	MeetingID uint   `gorm:"index;not null"`
	Content   string `gorm:"type:text;not null"`

	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Meeting Meeting `gorm:"foreignKey:MeetingID;constraint:OnDelete:CASCADE"`
}
