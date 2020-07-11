package middlewares

import (
	"context"
	"shop_server_go/db"
	"shop_server_go/models"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Cache() func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		ctx := context.Background()
		server := c.Get("server")
		method := c.Method()
		key := method + "_" + c.OriginalURL()
		val, err := db.RedisClient.Get(ctx, key).Result()
		if err == redis.Nil || server != "" {
			c.Next()
		} else if err != nil {
			panic(err)
		} else {
			var result models.Result
			json.Unmarshal([]byte(val), &result)
			c.JSON(result)
			return
		}

		res := c.Locals("res")
		result, _ := res.(models.Result)
		success := result.Success

		if server == "" && method == "GET" && success {
			if err := db.RedisClient.Set(ctx, key, res, 10*time.Second).Err(); err != nil {
				panic(err)
			}
		}
	}
}
