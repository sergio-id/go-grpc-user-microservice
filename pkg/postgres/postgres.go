package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
)

const (
	maxOpenConns    = 50
	connMaxLifetime = 2 * time.Minute
	maxIdleConns    = 30
	connMaxIdleTime = 20 * time.Second

	defaultConnAttempts = 3
	defaultConnTimeout  = time.Second
)

type postgres struct {
	connAttempts int
	connTimeout  time.Duration

	db *sqlx.DB
}

var _ DBEngine = (*postgres)(nil)

func NewPostgresDB(cfgPostgres Config) (DBEngine, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		cfgPostgres.Host,
		cfgPostgres.Port,
		cfgPostgres.User,
		cfgPostgres.DBName,
		cfgPostgres.Password,
	)

	pg := &postgres{
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.db, err = sqlx.Connect(cfgPostgres.PgDriverName, dataSourceName)
		if err == nil {
			break
		}

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	if pg.db == nil && err != nil {
		return nil, err
	}

	pg.db.SetMaxOpenConns(maxOpenConns)
	pg.db.SetConnMaxLifetime(connMaxLifetime)
	pg.db.SetMaxIdleConns(maxIdleConns)
	pg.db.SetConnMaxIdleTime(connMaxIdleTime)

	if err = pg.db.Ping(); err != nil {
		return nil, err
	}

	return pg, nil
}

func (p *postgres) GetDB() *sqlx.DB {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
