//go:generate mockgen -source usecase.go -destination ../mock/usecase.go -package mock
package domain

import (
	"context"
)

// UseCase is an interface for the usecase.
type UseCase interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*User, error)
}
