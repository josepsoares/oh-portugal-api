package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexMountains is a function to get all mountains data from the database
// @Summary Get all mountains
// @Description Get all mountains
// @Tags mountains
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Mountain}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mountains [get]
func IndexMountains(c *fiber.Ctx) error {
	// TODO => add query params support (region, sorting by id, altitude)

	mountains := []models.Mountain{}
	database.DB.Db.Find(&mountains)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting mountains",
		Data:    mountains,
	})
}

// GetMountainByID is a function to get a mountain by ID
// @Summary Get mountain by ID
// @Description Get mountain by ID
// @Tags mountains
// @Accept json
// @Produce json
// @Param id path int true "Mountain ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Mountain}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mountains/{id} [get]
func GetMountainByID(c *fiber.Ctx) error {
	id := c.Params("id")

	mountain := new(models.Mountain)
	if err := database.DB.Db.First(&mountain, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Mountain with ID %v not found.", id),
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
		Message: "Success getting mountain by ID",
		Data:    *mountain,
	})
}
