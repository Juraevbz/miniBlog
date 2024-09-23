package models

import (
	"time"
)

// type User struct {
// 	ID           uint      `json:"id" gorm:"primarykey"`
// 	Username     string    `json:"username"`
// 	Email        string    `json:"email"`
// 	PasswordHash string    `json:"password_hash"`
// 	AvatarURL    *string   `json:"avatar_url"`
// 	CreatedAt    time.Time `json:"created_at"`
// 	UpdatedAt    time.Time `json:"updated_at"`
// 	DeletedAt    time.Time `json:"deleted_at"`
// 	Posts        []Post    `gorm:"foreignKey:UserID"`
// 	Comments     []Comment `gorm:"foreignKey:UserID"`
// 	Repost       []Repost  `gorm:"foreignKey:UserID"`
// }

type Post struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	ImageURL  *string    `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}

type PostList struct {
	PostID   uint   `json:"post_id"`
	Title    string `json:"title"`
	Comments int    `json:"comments"`
	Likes    int    `json:"likes"`
}

type Comment struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	PostID    uint       `json:"post_id"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Like struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	PostID    uint       `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Repost struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	PostID    uint       `json:"post_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	ImageURL  *string    `json:"image_url"`
	Comments  int        `json:"comments"`
	Likes     int        `json:"likes"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
