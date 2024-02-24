package some_domain

import (
	"github.com/gofiber/fiber/v2"
)

func HandleRoot(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": false,
		"msg": "Hello, from some domain root!",
	})
}

func HandleEvent(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"error": false,
		"msg": c.Params("name"),
	})
}
