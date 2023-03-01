package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/domain"
	"github.com/sergio-id/go-grpc-user-microservice/internal/user/types"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_userPostgresqlRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repo := NewUserPostgresqlRepository(sqlxDB)

	columns := []string{
		"id",
		"email",
		"password",
		"first_name",
		"last_name",
		"about",
		"phone_number",
		"gender",
		"status",
		"last_ip",
		"last_device",
		"avatar_url",
		"updated_at",
		"created_at",
	}
	mockUser := &domain.User{
		ID:          1,
		Email:       "example@gmail.com",
		Password:    "111111",
		FirstName:   "firstName",
		LastName:    "lastName",
		About:       "about",
		PhoneNumber: "+100000000000",
		Gender:      string(types.Male),
		Status:      string(types.Active),
		LastIP:      "192.168.255.255",
		LastDevice:  "lastDevice",
		AvatarURL:   "https://avatar_url.com",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.ID,
		mockUser.Email,
		mockUser.Password,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(createUserCommand).WithArgs(
		mockUser.Email,
		mockUser.Password,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
	).WillReturnRows(rows)

	createdUser, err := repo.Create(context.Background(), mockUser)
	require.NoError(t, err)
	require.NotNil(t, createdUser)
	require.Equal(t, mockUser.Email, createdUser.Email)
	require.Equal(t, mockUser.Password, createdUser.Password)
	require.Equal(t, mockUser.FirstName, createdUser.FirstName)
	require.Equal(t, mockUser.LastName, createdUser.LastName)

	// ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func Test_userPostgresqlRepository_GetByEmail(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repo := NewUserPostgresqlRepository(sqlxDB)

	columns := []string{
		"id",
		"email",
		"password",
		"first_name",
		"last_name",
		"about",
		"phone_number",
		"gender",
		"status",
		"last_ip",
		"last_device",
		"avatar_url",
		"updated_at",
		"created_at",
	}
	mockUser := &domain.User{
		ID:          1,
		Email:       "example@gmail.com",
		Password:    "111111",
		FirstName:   "firstName",
		LastName:    "lastName",
		About:       "about",
		PhoneNumber: "+100000000000",
		Gender:      string(types.Male),
		Status:      string(types.Active),
		LastIP:      "192.168.255.255",
		LastDevice:  "lastDevice",
		AvatarURL:   "https://avatar_url.com",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.ID,
		mockUser.Email,
		mockUser.Password,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(getByEmailQuery).WithArgs(mockUser.Email).WillReturnRows(rows)

	foundUser, err := repo.GetByEmail(context.Background(), mockUser.Email)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.Email, mockUser.Email)
}

func Test_userPostgresqlRepository_GetById(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repo := NewUserPostgresqlRepository(sqlxDB)

	columns := []string{
		"id",
		"email",
		"password",
		"first_name",
		"last_name",
		"about",
		"phone_number",
		"gender",
		"status",
		"last_ip",
		"last_device",
		"avatar_url",
		"updated_at",
		"created_at",
	}
	mockUser := &domain.User{
		ID:          1,
		Email:       "example@gmail.com",
		Password:    "111111",
		FirstName:   "firstName",
		LastName:    "lastName",
		About:       "about",
		PhoneNumber: "+100000000000",
		Gender:      string(types.Male),
		Status:      string(types.Active),
		LastIP:      "192.168.255.255",
		LastDevice:  "lastDevice",
		AvatarURL:   "https://avatar_url.com",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.ID,
		mockUser.Email,
		mockUser.Password,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
		time.Now(),
		time.Now(),
	)

	mock.ExpectQuery(getByIDQuery).WithArgs(mockUser.ID).WillReturnRows(rows)

	foundUser, err := repo.GetById(context.Background(), mockUser.ID)
	require.NoError(t, err)
	require.NotNil(t, foundUser)
	require.Equal(t, foundUser.ID, mockUser.ID)
}

func Test_userPostgresqlRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	repo := NewUserPostgresqlRepository(sqlxDB)

	columns := []string{
		"id",
		"email",
		"password",
		"first_name",
		"last_name",
		"about",
		"phone_number",
		"gender",
		"status",
		"last_ip",
		"last_device",
		"avatar_url",
		"updated_at",
		"created_at",
	}
	mockUser := &domain.User{
		ID:          1,
		Email:       "example@gmail.com",
		Password:    "111111",
		FirstName:   "firstName",
		LastName:    "lastName",
		About:       "about",
		PhoneNumber: "+100000000000",
		Gender:      string(types.Male),
		Status:      string(types.Active),
		LastIP:      "192.168.255.255",
		LastDevice:  "lastDevice",
		AvatarURL:   "https://avatar_url.com",
	}

	rows := sqlmock.NewRows(columns).AddRow(
		mockUser.ID,
		mockUser.Email,
		mockUser.Password,
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
		time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	)

	mock.ExpectQuery(updateUserCommand).WithArgs(
		mockUser.FirstName,
		mockUser.LastName,
		mockUser.About,
		mockUser.PhoneNumber,
		mockUser.Gender,
		mockUser.Status,
		mockUser.LastIP,
		mockUser.LastDevice,
		mockUser.AvatarURL,
		mockUser.ID,
	).WillReturnRows(rows)

	updatedUser, err := repo.Update(context.Background(), mockUser)
	require.NoError(t, err)
	require.NotNil(t, updatedUser)
	require.Equal(t, mockUser, updatedUser)
}
