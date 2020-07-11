package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	ctx := context.Background()
	 RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	if _, err := RedisClient.Ping(ctx).Result(); err != nil {
		panic(err)
	} else {
		fmt.Println("redis server connected successfully")
	}

	
	
}