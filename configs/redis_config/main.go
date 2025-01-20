package redis_config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

var Config *RedisConfig

func LoadConfig() *RedisConfig {
	cfg := new(RedisConfig)
	if err := env.Parse(cfg); err != nil {
		panic(err)
	}
	return cfg
}
