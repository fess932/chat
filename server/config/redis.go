package config

import (
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func CreateRedisClient() {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		log.Fatal(err)
	}

	Redis = redis.NewClient(opt)
}
