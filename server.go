package main

import (
	"BlogProject/database"

	"github.com/gofiber/fiber/v2"
)

func init() {
	database.ConnectDatabase()
}

func main() {

	sqlDB, err := database.DBconn.DB()
	if err != nil {
		panic("Error in sql connection")
	}
	defer sqlDB.Close()
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	app.Listen("localhost:8000")

}
