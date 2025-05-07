package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Rdb *redis.Client
var Ctx = context.Background()

func InitRedis() {
	Rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
