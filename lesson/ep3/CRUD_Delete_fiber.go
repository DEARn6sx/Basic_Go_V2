package lesson

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)



func deleteBook(c *fiber.Ctx) error {
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
			//[1,2,3,4,5] ลบ 3
			//[1,2] + [4,5] = [1,2,4,5]
			books = append(books[:i], books[i+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

// func main() {
	
// 	lesson.Ep3_DeleteListBook()
// }
