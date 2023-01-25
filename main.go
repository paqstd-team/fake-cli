package main

import (
	"encoding/json"

	gf "github.com/brianvoe/gofakeit/v6"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		scheme := map[string]interface{}{
			"name":        gf.Name(),
			"description": gf.Phrase(),
		}
		resp, err := json.Marshal(scheme)
		if err != nil {
			panic(err)
		}

		return c.SendString((string(resp)))

	})

	app.Listen(":3000")
}
