package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username          string `gorm:"unique; not null"`
	Email             string `gorm:"unique; not null"`
	PasswordHash      string `gorm:"not null"`
	IsBlocked         bool   `gorm:"default:false"`
	PrifilePictureURL string
	Bio               string
	Posts             []Post     `gorm:"foreignKey:AuthorID"`
	Comments          []Comment  `gorm:"foreignKey:AuthorID"`
	Reactions         []Reaction `gorm:"foreignKey:AuthorID"`
	Favorites         []Favorite `gorm:"foreignKey:AuthorID"`
}

type Post struct {
	gorm.Model
	Title          string     `gorm:"not null"`
	Content        string     `gorm:"not null"`
	AuthorID       uint       `gorm:"not null"`
	RepostedFromID *uint      `gorm:"index"`
	Author         User       `gorm:"foreignKey:AuthorID"`
	Comments       []Comment  `gorm:"foreignKey:PostID"`
	Reactions      []Reaction `gorm:"foreignKey:PostID"`
	Favorites      []Favorite `gorm:"foreignKey:PostID"`
}

type Comment struct {
	gorm.Model
	Content  string `gorm:"not null"`
	PostID   uint   `gorm:"not null"`
	AuthorID uint   `gorm:"not null"`
	Post     Post   `gorm:"foreignKey:PostID"`
	Author   User   `gorm:""foreignKey:AuthorID`
}

type Reaction struct {
	gorm.Model
	PostID uint   `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Type   string `gorm:"not null"`
	Post   Post   `gorm:"foreignKey:PostID"`
	User   User   `gorm:"foreignKey:UserID"`
}

type Favorite struct {
	ID     uint `gorm:"foreignKey"`
	PostID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Post   Post `gorm:"foreignKey:PostID"`
	User   User `gorm:"foreignKey:UserID"`
}
