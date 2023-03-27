package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexLagoons is a function to get all lagoons data from database
// @Summary Get all lagoons
// @Description Get all lagoons
// @Tags lagoons
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Lagoon}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/lagoons [get]
func IndexLagoons(c *fiber.Ctx) error {
	lagoons := []models.Lagoon{}
	database.DB.Db.Find(&lagoons)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting lagoons",
		Data:    lagoons,
	})
}
