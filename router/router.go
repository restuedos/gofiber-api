package router

import (
	"gofiber-example/handler"
	"gofiber-example/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Get("/", middleware.Protected(), handler.GetAllUsers)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Book
	book := api.Group("/book")
	book.Get("/:id", middleware.Protected(), handler.GetBook)
	book.Get("/", middleware.Protected(), handler.GetAllBooks)
	book.Post("/", middleware.Protected(), handler.CreateBook)
	book.Patch("/:id", middleware.Protected(), handler.UpdateBook)
	book.Delete("/:id", middleware.Protected(), handler.DeleteBook)
}
