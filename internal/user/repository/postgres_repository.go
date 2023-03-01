package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
)

// PostgresqlRepository is a repository for Postgresql.
type userPostgresqlRepository struct {
	db *sqlx.DB
}

// NewUserPostgresqlRepository creates a new Postgresql repository.
func NewUserPostgresqlRepository(db *sqlx.DB) *userPostgresqlRepository {
	return &userPostgresqlRepository{db: db}
}

// Create method creates a new user.
func (r *userPostgresqlRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserPostgresqlRepository.Create")
	defer span.Finish()

	var u domain.User
	if err := r.db.QueryRowxContext(
		ctx,
		createUserCommand,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.About,
		user.PhoneNumber,
		user.Gender,
		user.Status,
		user.LastIP,
		user.LastDevice,
		user.AvatarURL,
	).StructScan(&u); err != nil {
		return nil, errors.Wrap(err, "UserPostgresqlRepository.Create.QueryRowxContext")
	}

	return &u, nil
}

// Update method updates an existing user.
func (r *userPostgresqlRepository) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserPostgresqlRepository.Update")
	defer span.Finish()

	var u domain.User
	if err := r.db.QueryRowxContext(
		ctx,
		updateUserCommand,
		&user.FirstName,
		&user.LastName,
		&user.About,
		&user.PhoneNumber,
		&user.Gender,
		&user.Status,
		&user.LastIP,
		&user.LastDevice,
		&user.AvatarURL,
		&user.ID,
	).StructScan(&u); err != nil {
		return nil, errors.Wrap(err, "UserPostgresqlRepository.Update.QueryRowxContext")
	}

	return &u, nil
}

// Delete method deletes an existing user.
func (r *userPostgresqlRepository) Delete(ctx context.Context, id uint64) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserPostgresqlRepository.Delete")
	defer span.Finish()

	result, err := r.db.ExecContext(ctx, deleteByIDCommand, id)
	if err != nil {
		return errors.Wrap(err, "UserPostgresqlRepository.Delete.ExecContext")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "UserPostgresqlRepository.Delete.RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "UserPostgresqlRepository.Delete.rowsAffected")
	}

	return nil
}

// GetById method returns user by id.
func (r *userPostgresqlRepository) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserPostgresqlRepository.GetById")
	defer span.Finish()

	var u domain.User
	if err := r.db.GetContext(ctx, &u, getByIDQuery, id); err != nil {
		return nil, errors.Wrap(err, "UserPostgresqlRepository.GetById.GetContext")
	}

	return &u, nil
}

// GetByEmail method returns user by email.
func (r *userPostgresqlRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "UserPostgresqlRepository.GetByEmail")
	defer span.Finish()

	var u domain.User
	if err := r.db.GetContext(ctx, &u, getByEmailQuery, email); err != nil {
		return nil, errors.Wrap(err, "UserPostgresqlRepository.GetByEmail.GetContext")
	}

	return &u, nil
}
