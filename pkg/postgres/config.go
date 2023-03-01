package postgres

// Config is a postgres config
type Config struct {
	Host         string `yaml:"host" env:"POSTGRES_HOST" env-default:"localhost"`
	Port         string `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
	User         string `yaml:"user" env:"POSTGRES_USER" env-default:"postgres"`
	DBName       string `yaml:"dbName" env:"POSTGRES_DB" env-default:"postgres"`
	Password     string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
	PgDriverName string `yaml:"pgDriverName" env:"PG_DRIVER_NAME" env-default:"pgx"`
}
