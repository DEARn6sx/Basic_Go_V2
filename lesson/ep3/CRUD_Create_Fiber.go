package lesson

import (
	"github.com/gofiber/fiber/v2"
)


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

