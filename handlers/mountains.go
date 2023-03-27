package handlers

import (
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
)

// IndexMountains is a function to get all mountains data from database
// @Summary Get all mountains
// @Description Get all mountains
// @Tags mountains
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Mountain}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/mountains [get]
func IndexMountains(c *fiber.Ctx) error {
	mountains := []models.Mountain{}
	database.DB.Db.Find(&mountains)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting regions",
		Data:    mountains,
	})
}
