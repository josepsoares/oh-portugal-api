package main

import (
	"fmt"
	"josepsoares/iberiapi/db"
	seeder "josepsoares/iberiapi/db/seeding"
	"josepsoares/iberiapi/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

// @title Oh, Portugal API! :portugal:
// @version 1.0
// @description This is an API for getting some general info about Portugal

// @contact.name josepsoares
// @contact.email josepsoares.dev@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api
func main() {
	argLength := len(os.Args[1:])
	if argLength > 1 {

	}

	db.Connect()

	if argLength == 1 && os.Args[1] == "seed" {
		seeder.Seed()
		fmt.Println("Seeded the db successfully 🌱")
		fmt.Println("Exiting the program... 😌")
		os.Exit(0)
	} else {
		engine := html.New("./views", ".html")

		app := fiber.New(fiber.Config{
			Views: engine,
		})

		routes.InitRoutes(app)

		app.Use(cors.New())
		app.Use(limiter.New())
		app.Use(logger.New(logger.Config{
			Format:     "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Europe/Lisbon",
		}))

		app.Use(func(c *fiber.Ctx) error {
			return c.SendStatus(404) // => 404 "Not Found"
		})

		log.Fatal(app.Listen(":8080"))
	}
}
