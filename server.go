package main

import (
	"shop_server_go/config"
	"shop_server_go/middlewares"
	"shop_server_go/models"

	"github.com/gofiber/fiber"
	"github.com/ilibs/gosql/v2"
)

func main() {
	app := fiber.New()

	app.Use(middlewares.Cache())

	app.Use(middlewares.Serialize())

	app.Get("/", func(c *fiber.Ctx) {
		var goods []models.Goods
		sql := "select * from goods"
		gosql.Select(&goods, sql)
		c.Locals("data", goods)
		return
	})

	port := config.GetValue("server", "http_port")
	app.Listen(port)
}
