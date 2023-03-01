package logger

// Config is a logger config
type Config struct {
	LogLevel string `yaml:"logLevel" env:"LOG_LEVEL" env-default:"info"`
	Console  bool   `yaml:"console" env:"LOG_CONSOLE" env-default:"true"`
}
