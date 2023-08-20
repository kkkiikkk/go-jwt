package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkkiikkk/go-jwt/handler"
	"github.com/kkkiikkk/go-jwt/middleware"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// Image
	api.Post("/upload_image", middleware.Protected(), handler.UploadImage)
	api.Get("/images", middleware.Protected(), handler.GetImages)
}