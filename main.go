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

	engine := html.New("./templates", ".html")

	app := fiber.New(fiber.Config{Views: engine})

	app.Static("/static", "./static")

	app.Use(cors.New(cors.ConfigDefault))

	data := fiber.Map{
		"Title": "Hello, World!"}

	app.Get("/", func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "refresh",
			Value: "refresh",
			Path:  "/",
		})

		return c.Render("Base", data)
	})

	app.Listen("localhost:3000")

}
