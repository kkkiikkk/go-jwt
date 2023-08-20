package main

import (
	"github.com/kkkiikkk/go-jwt/database"
	"github.com/kkkiikkk/go-jwt/router"
	"github.com/kkkiikkk/go-jwt/config"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}