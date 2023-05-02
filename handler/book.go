package handler

import (
	"gofiber-example/database"
	"gofiber-example/model"

	"github.com/gofiber/fiber/v2"
)

// GetAllBooks query all books
func GetAllBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []model.Book
	db.Find(&books)
	return c.JSON(fiber.Map{"status": "success", "message": "All books", "data": books})
}

// GetBook query book
func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book model.Book
	db.Find(&book, id)
	if book.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Book found", "data": book})
}

// CreateBook new book
func CreateBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(model.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create book", "data": err})
	}
	db.Create(&book)
	return c.JSON(fiber.Map{"status": "success", "message": "Created book", "data": book})
}

// UpdateBook update book
func UpdateBook(c *fiber.Ctx) error {
	type UpdateBookInput struct {
		Description string `json:"description"`
	}
	var ubi UpdateBookInput
	if err := c.BodyParser(&ubi); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	id := c.Params("id")

	db := database.DBConn
	var book model.Book

	db.First(&book, id)
	if book.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book found with ID", "data": nil})
	}
	book.Description = ubi.Description
	db.Save(&book)

	return c.JSON(fiber.Map{"status": "success", "message": "Book successfully updated", "data": book})
}

// DeleteBook delete book
func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.First(&book, id)
	if book.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No book found with ID", "data": nil})
	}
	db.Delete(&book)
	return c.JSON(fiber.Map{"status": "success", "message": "Book successfully deleted", "data": nil})
}
