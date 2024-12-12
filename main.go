package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()


	app.get("/hello", func(c *fiber.Ctx)error{
		return c.SendString("hello gofiber again ")
	})

	app.Listen(":8080")
}