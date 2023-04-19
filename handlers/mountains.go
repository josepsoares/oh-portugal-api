package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/db"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
	"gorm.io/gorm/clause"
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
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	altitudeQueryParam := c.Query("altitude")
	// regionIdQueryParam := c.Query("region_id")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	mountains := []models.Mountain{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "altitude", altitudeQueryParam)
	// utils.FilterIntClause(queryClauses, "area", regionIdQueryParam)

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "altitude") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	db.DBConn.Joins("regions").Clauses(queryClauses...).Order(sort + " " + order).Find(&mountains)

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
