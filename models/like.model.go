package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
    UserID    uuid.UUID `gorm:"not null" json:"user,omitempty"`
    PostID    uuid.UUID `gorm:"not null" json:"post,omitempty"`
    CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type CreatePostLike struct {
	UserID    string    `json:"user,omitempty"`
	PostID    string    `json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}