package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexRegions is a function to get all regions data from the database
// @Summary Get all regions
// @Description Get all regions
// @Tags regions
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Region}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/regions [get]
func IndexRegions(c *fiber.Ctx) error {
	// TODO => add query params support (autonomous, sorting by id, population total, population density, area, nr. districts, nr. municipalities, nr. freguesias)

	regions := []models.Region{}
	database.DB.Db.Find(&regions)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting regions",
		Data:    regions,
	})
}

// GetRegionByID is a function to get a region by ID
// @Summary Get region by ID
// @Description Get region by ID
// @Tags regions
// @Accept json
// @Produce json
// @Param id path int true "Region ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Region}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/regions/{id} [get]
func GetRegionByID(c *fiber.Ctx) error {
	id := c.Params("id")

	region := new(models.Region)
	if err := database.DB.Db.First(&region, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Region with ID %v not found", id),
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
		Message: "Success getting region by ID",
		Data:    *region,
	})
}
