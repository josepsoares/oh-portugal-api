package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexRivers is a function to get all rivers data from database
// @Summary Get all rivers
// @Description Get all rivers
// @Tags rivers
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.River}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/rivers [get]
func IndexRivers(c *fiber.Ctx) error {
	rivers := []models.River{}
	database.DB.Db.Find(&rivers)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting islands",
		Data:    rivers,
	})
}
