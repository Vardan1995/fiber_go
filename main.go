package main

import (
	"os"

	"github.com/Vardan1995/fiber-crud/database"
	"github.com/Vardan1995/fiber-crud/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		
	})
	database.ConnectDB()
	router.SetupRouter(app)

	port := os.Getenv("DB_HOST")
	if port == "" {
		port="3000"
	}
	app.Listen(":"+port)

	// defer database.DB.Close()
}
