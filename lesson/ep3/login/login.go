package lesson

import "github.com/gofiber/fiber/v2"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var memberUser = User{
	Email:    "user@ex.com",
	Password: "pass123",
}

func Ep3_login() {
	app := fiber.New()
	app.Post("/login", login)

	app.Listen(":8080")
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if user.Email != memberUser.Email || user.Password != memberUser.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"massage": "login success",
	})
}
