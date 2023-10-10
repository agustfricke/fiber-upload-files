package main

import (
	"github.com/agustfricke/fiber-upload-files/database"
	"github.com/agustfricke/fiber-upload-files/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

  database.ConnectDB()

  engine := html.New("./templates", ".html")

  app := fiber.New(fiber.Config{
    Views: engine, 
  })

  app.Static("/", "./public")

  app.Get("/", handlers.Home)
  app.Post("/", handlers.Upload)

  app.Listen(":8080")
}
