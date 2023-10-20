package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
    UserID    uuid.UUID `gorm:"not null" json:"user,omitempty"`
    PostID    uuid.UUID `gorm:"not null" json:"post,omitempty"`
    Text      string    `gorm:"not null" json:"text,omitempty"`
    CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostComment struct {
	Text      string    `json:"text"  binding:"required"`
	UserID    string    `json:"user,omitempty"`
	PostID    string    `json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateComment struct {
	Text      string    `json:"text,omitempty"`
	UserID    string    `json:"user,omitempty"`
	PostID    string    `json:"post,omitempty"`
	CreateAt  time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
