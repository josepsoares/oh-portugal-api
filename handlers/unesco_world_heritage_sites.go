package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexUnescoWorldHeritageSites is a function to get all unesco world heritage sites of Portugal data from database
// @Summary Get all unesco world heritage sites of Portugal
// @Description Get all unesco world heritage sites of Portugal
// @Tags unesco-world-heritage-sites
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.UnescoWorldHeritageSite}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/unesco-world-heritage-sites [get]
func IndexUnescoWorldHeritageSites(c *fiber.Ctx) error {
	unescoWorldHeritageSites := []models.UnescoWorldHeritageSite{}
	database.DB.Db.Find(&unescoWorldHeritageSites)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting unesco world heritage sites",
		Data:    unescoWorldHeritageSites,
	})
}
