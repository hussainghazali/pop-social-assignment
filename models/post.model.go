package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
    Title     string    `gorm:"uniqueIndex;not null" json:"title,omitempty"`
    Content   string    `gorm:"not null" json:"content,omitempty"`
    ImagePath string    `gorm:"not null" json:"imagePath,omitempty"`
    VideoPath string    `gorm:"not null" json:"videoPath,omitempty"`
    UserID    uuid.UUID `gorm:"not null" json:"user,omitempty"`
    Likes     []Like    `gorm:"foreignKey:PostID"` // One-to-Many relationship
	LikeCount     int    `gorm:"-"`
    Comments  []Comment `gorm:"foreignKey:PostID"` // One-to-Many relationship
    CommentCount  int    `gorm:"-"`
    CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostRequest struct {
	Title     string    `json:"title"  binding:"required"`
	Content   string    `json:"content" binding:"required"`
	ImagePath string    `json:"imagePath" binding:"required"`
    VideoPath string    `json:"videoPath" binding:"required"`
	UserID    string    `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdatePost struct {
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	ImagePath string    `json:"imagePath,omitempty"`
    VideoPath string    `json:"videoPath,omitempty"`
	UserID    string    `json:"user,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}