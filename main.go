package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/count", func(c *fiber.Ctx) error {
		connStr := "postgresql://postgres:a@localhost:5432/postgres?sslmode=disable"
		db, err := sql.Open("postgres", connStr)

		if err != nil {
			log.Fatal(err)
		}

		var result = "There will be a result."
		r, err := db.Query("SELECT count(*) FROM dates;")
		if err != nil {
			log.Fatalln(err)
			c.JSON("An error occurred")
		}

		for r.Next() {
			r.Scan(&result)
		}

		return c.SendString(result)
	})

	app.Listen(":3000")
}
