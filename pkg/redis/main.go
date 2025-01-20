package pkg_redis

import (
	"github.com/redis/go-redis/v9"
)

type RedisManager struct {
	client *redis.Client
}

func NewRedisManager(addr string, password string, db int) RedisManager {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return RedisManager{
		client: client,
	}
}

func (r *RedisManager) Client() *redis.Client {
	return r.client
}

func (r *RedisManager) Close() error {
	return r.client.Close()
}
