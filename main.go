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

	sendReload := true

	app.Get("/reload", func(c *fiber.Ctx) error {
		// SEND TO RELOAD THE FIRST TIME (TRRIGER HOT RELOAD)
		if sendReload {
			sendReload = false
			return c.JSON(fiber.Map{"reload": "true"})

		} else {
			return c.JSON(fiber.Map{"reload": "false"})
		}
	})

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("Base", data)
	})

	app.Listen("localhost:3000")

}
