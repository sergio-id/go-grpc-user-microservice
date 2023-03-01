package migrations

// Config is a migrations config
type Config struct {
	Enable    bool   `yaml:"enable" env:"MIGRATIONS_ENABLE" env-default:"false"`
	Recreate  bool   `yaml:"recreate" env:"MIGRATIONS_RECREATE" env-default:"false"`
	SourceURL string `yaml:"sourceURL" env:"MIGRATIONS_SOURCE_URL" env-default:"file://migrations"`
	DbURL     string `yaml:"dbURL" env:"MIGRATIONS_DB_URL" env-default:"postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"`
}
