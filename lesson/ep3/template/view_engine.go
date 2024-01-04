package lesson

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func Ep3_Views_Template() {
	// Initialize standard Go html template engine
	engine := html.New("./lesson/ep3/template/views", ".html")

	// Pass the engine to Fiber
	app := fiber.New(fiber.Config{
	  Views: engine,
	})
  
	// Setup route
	app.Get("/template", renderTemplate)
  
	app.Listen(":8080")
  }
  
  func renderTemplate(c *fiber.Ctx) error {
	// Render the template with variable data
	return c.Render("template", fiber.Map{
	  "Name": "World",
	})
  }