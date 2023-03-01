package domain

import (
	"time"
)

// User is a struct for the user.
type User struct {
	ID          uint64    `json:"id" db:"id" validate:"omitempty"`
	Email       string    `json:"email" db:"email" validate:"omitempty,lte=64,email"`
	Password    string    `json:"password,omitempty" db:"password" validate:"omitempty,gte=8,lte=255"`
	FirstName   string    `json:"first_name" db:"first_name" validate:"omitempty,lte=64"`
	LastName    string    `json:"last_name" db:"last_name" validate:"omitempty,lte=64"`
	About       string    `json:"about" db:"about" validate:"omitempty,lte=4096"`
	PhoneNumber string    `json:"phone_number" db:"phone_number" validate:"omitempty,lte=64"`
	Gender      string    `json:"gender" db:"gender" validate:"required,gender"`
	Status      string    `json:"status" db:"status" validate:"required,status"`
	LastIP      string    `json:"last_ip" db:"last_ip" validate:"omitempty"`
	LastDevice  string    `json:"last_device" db:"last_device" validate:"omitempty,lte=64"`
	AvatarURL   string    `json:"avatar_url" db:"avatar_url" validate:"omitempty,lte=255"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
}
