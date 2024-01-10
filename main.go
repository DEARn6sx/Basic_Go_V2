package main

import (
	lesson "Go_lang/lesson/ep4"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	lesson.Connect_DB()

	app := fiber.New()

	app.Get("/product/:id", lesson.GetProductHandler)
	app.Get("/product", lesson.GetallProductHandler)
	app.Post("/product", lesson.CreateProductHandler)
	app.Patch("/product/:id", lesson.UpdateProductHandler)
	app.Delete("/product/:id", lesson.DeleteProductHandler)
	app.Listen(":8080")

}
