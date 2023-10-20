package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
    Name      string    `gorm:"type:varchar(255);not null" json:"name,omitempty"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email,omitempty"`
    Password  string    `gorm:"not null" json:"password,omitempty"`
    Verified  bool      `gorm:"not null" json:"verified,omitempty"`
    Posts     []Post    `gorm:"foreignKey:UserID"` // One-to-Many relationship
	Likes     []Like    `gorm:"foreignKey:UserID"` // One-to-Many relationship
    Comments  []Comment `gorm:"foreignKey:UserID"` // One-to-Many relationship
    CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
    UpdatedAt time.Time `gorm:"not null" json:"updated_at,omitempty"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}