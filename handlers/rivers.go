package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexRivers is a function to get all rivers data from the database
// @Summary Get all rivers
// @Description Get all rivers
// @Tags rivers
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.River}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/rivers [get]
func IndexRivers(c *fiber.Ctx) error {
	// TODO => add query params support (region, source country, sorting by id, length, avg. flow, watershed area)

	rivers := []models.River{}
	database.DB.Db.Find(&rivers)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting rivers",
		Data:    rivers,
	})
}

// GetRiverByID is a function to get a river by ID
// @Summary Get river by ID
// @Description Get river by ID
// @Tags rivers
// @Accept json
// @Produce json
// @Param id path int true "River ID"
// @Success 200 {object} ResponseHTTP{data=[]models.River}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/rivers/{id} [get]
func GetRiverByID(c *fiber.Ctx) error {
	id := c.Params("id")

	river := new(models.River)
	if err := database.DB.Db.First(&river, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("River with ID %v not found", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(utils.ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		}
	}

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting river by ID",
		Data:    *river,
	})
}
