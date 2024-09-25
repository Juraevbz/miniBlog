package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           uint      `json:"id" gorm:"primarykey"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at"`
}

func (u *User) Validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}
	if u.PasswordHash == "" {
		return errors.New("password is required")
	}

	return nil
}

type Post struct {
	ID        int        `json:"id" gorm:"primaryKey"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Comments  []*Comment `json:"comments"`
	Likes     []*Like    `json:"likes"`
}

func (u *User) ComparePassword(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(rawPassword))
	return err == nil
}

func GeneratePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
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
