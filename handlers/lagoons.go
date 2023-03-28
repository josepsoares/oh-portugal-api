package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexLagoons is a function to get all lagoons data from the database
// @Summary Get all lagoons
// @Description Get all lagoons
// @Tags lagoons
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Lagoon}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/lagoons [get]
func IndexLagoons(c *fiber.Ctx) error {
	// TODO => add query params support (region, sorting by id, area, depth)

	lagoons := []models.Lagoon{}
	database.DB.Db.Find(&lagoons)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting lagoons",
		Data:    lagoons,
	})
}

// GetLagoonByID is a function to get a lagoon by ID
// @Summary Get lagoon by ID
// @Description Get lagoon by ID
// @Tags lagoons
// @Accept json
// @Produce json
// @Param id path int true "Lagoon ID"
// @Success 200 {object} ResponseHTTP{data=[]models.Lagoon}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/lagoons/{id} [get]
func GetLagoonByID(c *fiber.Ctx) error {
	id := c.Params("id")

	lagoon := new(models.Lagoon)
	if err := database.DB.Db.First(&lagoon, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Lagoon with ID %v not found.", id),
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
		Message: "Success getting lagoon by ID",
		Data:    *lagoon,
	})
}
