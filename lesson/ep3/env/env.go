package lesson

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func Ep3_Env() {

	godotenv.Load()
	app := fiber.New()
	app.Get("/config", getEnv)

	app.Listen(":8080")
}

func getEnv(c *fiber.Ctx) error {

	secret := os.Getenv("SECRET")

	if secret == "" {
		secret ="defaultsecret"
	}
	return c.JSON(fiber.Map{
		"SECRET": secret,
	})
}
