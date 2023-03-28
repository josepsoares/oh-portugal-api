package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexUnescoWorldHeritageSites is a function to get all Unesco World Heritage sites of Portugal data from the database
// @Summary Get all Unesco World Heritage sites of Portugal
// @Description Get all Unesco World Heritage sites of Portugal
// @Tags unesco-world-heritage-sites
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.UnescoWorldHeritageSite}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/unesco-world-heritage-sites [get]
func IndexUnescoWorldHeritageSites(c *fiber.Ctx) error {
	// TODO => add query params support (region, sorting by id, date)

	unescoWorldHeritageSites := []models.UnescoWorldHeritageSite{}
	database.DB.Db.Find(&unescoWorldHeritageSites)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting unesco world heritage sites",
		Data:    unescoWorldHeritageSites,
	})
}

// GetUnescoWorldHeritageSiteByID is a function to get a Unesco World Heritage site by ID
// @Summary Get Unesco World Heritage site by ID
// @Description Get Unesco World Heritage site by ID
// @Tags regions
// @Accept json
// @Produce json
// @Param id path int true "Unesco World Heritage site ID"
// @Success 200 {object} ResponseHTTP{data=[]models.UnescoWorldHeritageSite}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/unesco-world-heritage-sites/{id} [get]
func GetUnescoWorldHeritageSiteByID(c *fiber.Ctx) error {
	id := c.Params("id")

	unescoWorldHeritageSite := new(models.UnescoWorldHeritageSite)
	if err := database.DB.Db.First(&unescoWorldHeritageSite, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Unesco World Heritage site with ID %v not found", id),
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
		Message: "Success getting Unesco World Heritage site by ID",
		Data:    *unescoWorldHeritageSite,
	})
}
