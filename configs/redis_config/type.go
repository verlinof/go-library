package redis_config

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
	Db       int    `env:"REDIS_DB"`
}
