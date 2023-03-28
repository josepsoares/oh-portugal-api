package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// GetPortugal is a function to get all general data of Portugal from the database
// @Summary Get general data of Portugal
// @Description Get general data of Portugal
// @Tags portugal
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Portugal}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/ [get]
func GetPortugal(c *fiber.Ctx) error {
	portugal := new(models.Country)
	if err := database.DB.Db.First(&portugal, 1).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: "Portugal info not found",
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
		Message: "Success getting Portugal info",
		Data:    portugal,
	})
}
