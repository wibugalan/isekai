package main

import (
	"isekai/Isekai/db"
	"isekai/Isekai/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.Connect()

	app := fiber.New()

	routes.Setup(app)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}
