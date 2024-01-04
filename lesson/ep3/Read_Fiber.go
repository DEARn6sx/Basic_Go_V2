package lesson

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func Ep3_ReadListBook() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "DEAR Fiber API", Author: "DEAR"})
	books = append(books, Book{ID: 2, Title: "DEAR Fiber API V 2", Author: "DEAR V2"})

	app.Get("/books", getBook)
	app.Get("/books/:id", getBookID)

	app.Listen(":8080")
}

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBookID(c *fiber.Ctx) error {
	//แปลง string เป็น int
	bookId, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for _, book := range books {
		if book.ID == bookId {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}


// func main() {
	
// 	lesson.Ep3_ReadListBook()
// }