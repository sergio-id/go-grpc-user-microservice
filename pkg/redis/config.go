package redis

// Config is a redis config
type Config struct {
	Addr         string `yaml:"addr" env:"REDIS_ADDR" env-default:"localhost:6379"`
	MinIdleConns int    `yaml:"minIdleConns" env:"REDIS_MIN_IDLE_CONNS" env-default:"200"`
	PoolSize     int    `yaml:"poolSize" env:"REDIS_POOL_SIZE" env-default:"200"`
	PoolTimeout  int    `yaml:"poolTimeout" env:"REDIS_POOL_TIMEOUT" env-default:"5"`
	Password     string `yaml:"password" env:"REDIS_PASSWORD" env-default:""`
	DB           int    `yaml:"db" env:"REDIS_DB" env-default:"0"`
}
