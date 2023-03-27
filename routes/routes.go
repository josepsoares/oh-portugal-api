package routes

import (
	"josepsoares/oh-portugal-api/handlers"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// New create an instance of Book app routes
func InitRoutes(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"hello": "world",
		})
	})

	app.Get("/docs/*", swagger.HandlerDefault)

	api := app.Group("/api")
	v1 := api.Group("/v1", func(c *fiber.Ctx) error {
		c.JSON(fiber.Map{
			"message": "🐣 v1",
		})
		return c.Next()
	})
	v1.Get("/", handlers.GetPortugal)

	v1.Get("/regions", handlers.IndexRegions)
	v1.Get("/regions/:id", handlers.GetRegionByID)

	v1.Get("/islands", handlers.IndexIslands)
	v1.Get("/rivers", handlers.IndexRivers)
	v1.Get("/lagoons", handlers.IndexLagoons)
	v1.Get("/mountains", handlers.IndexMountains)
	v1.Get("/unesco-world-heritage-sites", handlers.IndexUnescoWorldHeritageSites)

	return app
}
