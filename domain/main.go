package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	inicializeServer()
}

func inicializeServer() {

	engine := html.New("../templates", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/", "../static")

	app.Use(cors.New(cors.ConfigDefault))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	app.Listen("localhost:3000")

}
