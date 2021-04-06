package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"rafaignacio.com/auth/src/internal/utils"
	"rafaignacio.com/auth/src/pkg/apis"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	config, err := utils.LoadConfig("../configs")

	if err != nil {
		panic(err)
	}

	app.Use(requestid.New())
	app.Use(logger.New())

	api := app.Group("/api")
	users := api.Group("/users")

	users.Post("/", apis.AddNewUser)
	users.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": []string{},
		})
	})

	app.Listen(config.ServerAddress)
}
