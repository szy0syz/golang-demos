package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"jerryshi.com/fornodejs/internal/handlers"
	"log"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {
	app := fiber.New()

	app.Get("/healthcheck", healthcheck)

	app.Post("/api/products", handlers.CreateProduct)

	fmt.Println("Hello World")

	log.Fatal(app.Listen(":3000"))
}
