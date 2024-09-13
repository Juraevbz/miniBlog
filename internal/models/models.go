package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique; not null"`
	Email        string `gorm:"unique; not null"`
	PasswordHash string `gorm:"not null"`
	AvatarURL    *string
	Posts        []Post    `gorm:"foreignKey:UserID"`
	Comments     []Comment `gorm:"foreignKey:UserID"`
	Repost       []Repost  `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	UserID   uint   `gorm:"not null"`
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	ImageURL string
}

type Comment struct {
	gorm.Model
	UserID  uint
	PostID  uint
	Content string
}

type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	PostID    uint
	CreatedAt time.Time
}

type Repost struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint
	UserID    uint
	CreatedAt time.Time
}
