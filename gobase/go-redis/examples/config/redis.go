package config

import (
	"github.com/redis/go-redis/v9"
)

func GetRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})
}
