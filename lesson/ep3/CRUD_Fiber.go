package lesson

import "github.com/gofiber/fiber/v2"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func Ep3_ListBook() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "DEAR Fiber API", Author: "DEAR"})
	books = append(books, Book{ID: 2, Title: "DEAR Fiber API V 2", Author: "DEAR V2"})

	app.Get("/books", getBook)

	app.Listen(":8080")
}

func getBook(c *fiber.Ctx) error {
	return c.JSON(books)
}
