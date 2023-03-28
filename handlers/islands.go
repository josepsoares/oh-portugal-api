package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexIslands is a function to get all islands data from the database
// @Summary Get all islands
// @Description Get all islands
// @Tags islands
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Island}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/islands [get]
func IndexIslands(c *fiber.Ctx) error {
	// TODO => add query params support (region, deserted, sorting by id, population total, population density, area)

	islands := []models.Island{}
	database.DB.Db.Find(&islands)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting islands",
		Data:    islands,
	})
}

// GetIslandByID is a function to get a island by ID
// @Summary Get island by ID
// @Description Get island by ID
// @Tags islands
// @Accept json
// @Produce json
// @Param id path int true "Island ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Island}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/islands/{id} [get]
func GetIslandByID(c *fiber.Ctx) error {
	id := c.Params("id")

	island := new(models.Island)
	if err := database.DB.Db.First(&island, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Island with ID %v not found.", id),
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
		Message: "Success getting island by ID",
		Data:    *island,
	})
}
