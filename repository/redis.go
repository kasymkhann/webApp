package repository

import "github.com/go-redis/redis"

var Client *redis.Client

func RedisDB() {
	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:5323",
	})
}
