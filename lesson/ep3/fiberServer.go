package lesson

import (
	"github.com/gofiber/fiber/v2"
)

func Ep3_FiberRunHttpServer() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	app.Listen(":8080")
}


// func main() {
// 	lesson.Ep3_FiberRunHttpServer()	
// }

