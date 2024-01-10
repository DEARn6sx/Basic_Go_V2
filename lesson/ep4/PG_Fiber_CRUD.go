package lesson

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	product, err := getProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)

}

func GetallProductHandler(c *fiber.Ctx) error {

	product, err := getAllProduct()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)

}

func CreateProductHandler(c *fiber.Ctx) error {
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := createProduct(p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(p)
}

func UpdateProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	p := new(Product)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	product, err := updateProduct(productId, p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(product)
}

func DeleteProductHandler(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	err = deleteProduct(productId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusNoContent)

}
