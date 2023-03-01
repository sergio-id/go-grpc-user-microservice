//go:generate mockgen -source postgres_repository.go -destination ../mock/postgres_repository.go -package mock
package domain

import (
	"context"
)

// PostgresqlRepository is an interface for the postgresql repository.
type PostgresqlRepository interface {
	Create(ctx context.Context, user *User) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}
