package postgres

import "github.com/jmoiron/sqlx"

type DBEngine interface {
	GetDB() *sqlx.DB
	Close()
}
