package middlewares

import (
	"github.com/gofiber/fiber"
	"shop_server_go/models"
)

func Serialize() func(c *fiber.Ctx) {
	return func (c *fiber.Ctx) {
		c.Next()
		data := c.Locals("data")
		res := models.Result{
			Success: true,
			Data:    data,
			Message: "操作成功",
		}
		c.Locals("data", nil)
		c.Locals("res", res)
		c.JSON(res)
	}
}