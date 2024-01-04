package lesson

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Ep3_UpdateListBook() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "DEAR Fiber API", Author: "DEAR"})
	books = append(books, Book{ID: 2, Title: "DEAR Fiber API V 2", Author: "DEAR V2"})

	app.Get("/books", getBook)
	app.Get("/books/:id", getBookID)
	app.Post("/books", createBook)
	app.Put("/books/:id", updateBook)

	app.Listen(":8080")
}

func updateBook(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	for i, book := range books {
		if book.ID == bookId {
			books[i].Title = bookUpdate.Title
			books[i].Author = bookUpdate.Author
			return c.JSON(books[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}


// func main() {
	
// 	lesson.Ep3_UpdateListBook()
// }