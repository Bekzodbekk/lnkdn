package redis

import (
	"auth-service/internal/pkg/config"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func InitRedis(cfg *config.Config) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
		DB:       0,
		Password: "",
	})

	return rdb
}
