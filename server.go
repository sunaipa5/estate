package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func startServer() {
	var app = fiber.New()

	//CORS allow origin ! Clear for stable build !
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//Call static files
	foAdmin := "./frontend/estate-admin/dist/index.html"

	app.Static("/assets", "./frontend/estate-admin/dist/assets")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile(foAdmin)
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.SendFile(foAdmin)
	})
	app.Listen(":3000")
}
