package models

import (
	"errors"
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
	ID        int        `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	ImageURL  *string    `json:"image_url"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Content == "" {
		return errors.New("content is required")
	}

	return nil
}

type PostList struct {
	PostID   int    `json:"post_id"`
	Title    string `json:"title"`
	Comments int    `json:"comments"`
	Likes    int    `json:"likes"`
}

type Comment struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	PostID    int        `json:"post_id"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (c *Comment) Validate() error {
	if c.PostID == 0 {
		return errors.New("post_id is required")
	}
	if c.Comment == "" {
		return errors.New("comment is required")
	}

	return nil
}

type Like struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	PostID    int        `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (l *Like) Validate() error {
	if l.PostID == 0 {
		return errors.New("post_id is required")
	}

	return nil
}

type Repost struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	PostID    int        `json:"post_id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	ImageURL  *string    `json:"image_url"`
	Comments  int        `json:"comments"`
	Likes     int        `json:"likes"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (r *Repost) Validate() error {
	if r.PostID == 0 {
		return errors.New("post_id is required")
	}

	return nil
}