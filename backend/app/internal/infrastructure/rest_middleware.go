package infrastructure

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func initRestMiddleware(app *fiber.App, appKey string) error {
	app.Use(limiter.New(limiter.ConfigDefault))
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		if c.GetRespHeader("x-app-key") != appKey {
			return c.SendStatus(fiber.StatusForbidden)
		}
		return nil
	})
	return nil
}
