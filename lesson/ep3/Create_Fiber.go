package lesson

import (
	"github.com/gofiber/fiber/v2"
)

func Ep3_CreateListBook() {
	app := fiber.New()

	books = append(books, Book{ID: 1, Title: "DEAR Fiber API", Author: "DEAR"})
	books = append(books, Book{ID: 2, Title: "DEAR Fiber API V 2", Author: "DEAR V2"})

	app.Get("/books", getBook)
	app.Get("/books/:id", getBookID)
	app.Post("/books", createBook)

	app.Listen(":8080")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book) //ตัวแทนในการรับ request
	//c.BodyParser(book) //แปลงข้อมูลจาก body request map ไปยัง struck book
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

// func main() {
	
// 	lesson.Ep3_CreateListBook()
// }

