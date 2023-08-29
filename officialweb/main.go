package main

import (
	"log"

	i "officialweb/internal"

	"github.com/gofiber/fiber/v2"
)

const port = ":8081"

func main() {
	fiberapp := fiber.New()

	api := i.NewAPI(fiberapp)

	log.Fatal(api.App.Listen(port))
}
