package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexIslands is a function to get all islands data from database
// @Summary Get all islands
// @Description Get all islands
// @Tags islands
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Island}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/islands [get]
func IndexIslands(c *fiber.Ctx) error {
	islands := []models.Island{}
	database.DB.Db.Find(&islands)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting islands",
		Data:    islands,
	})
}
