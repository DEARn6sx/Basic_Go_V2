package lesson

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)


func getBook(c *fiber.Ctx) error {
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