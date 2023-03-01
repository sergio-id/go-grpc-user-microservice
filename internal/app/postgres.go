package app

import (
	"github.com/jmoiron/sqlx"
	"github.com/sergio-id/go-grpc-user-microservice/pkg/postgres"
)

func getDB(cfgPostgres postgres.Config) (*sqlx.DB, func(), error) {
	db, err := postgres.NewPostgresDB(cfgPostgres)
	if err != nil {
		return nil, nil, err
	}
	return db.GetDB(), func() { db.Close() }, nil
}
