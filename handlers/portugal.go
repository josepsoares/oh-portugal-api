package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

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
	portugal := []models.Portugal{}
	database.DB.Db.Find(&portugal)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting regions",
		Data:    portugal,
	})
}
