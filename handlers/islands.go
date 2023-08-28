package handlers

import (
	"fmt"
	"josepsoares/iberiapi/db"
	"josepsoares/iberiapi/models"
	"josepsoares/iberiapi/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
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
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	desertedQueryParam := c.Query("deserted")
	populationQueryParam := c.Query("population")
	populationDensityQueryParam := c.Query("population_density")
	areaQueryParam := c.Query("area")
	// regionIdQueryParam := c.Query("region_id")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	islands := []models.Island{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "population", populationQueryParam)
	utils.FilterIntClause(queryClauses, "population_density", populationDensityQueryParam)
	utils.FilterIntClause(queryClauses, "area", areaQueryParam)
	// utils.FilterIntClause(queryClauses, "area", regionIdQueryParam)

	if desertedQueryParam != "" && (desertedQueryParam == "true" || desertedQueryParam == "false") {
		if desertedQueryParam == "true" {
			queryClauses = append(queryClauses, clause.Like{Column: "population", Value: 0})
		} else {
			queryClauses = append(queryClauses, clause.Gte{Column: "population", Value: 1})
		}
	}

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "population" || sortByQueryParam == "population_density" || sortByQueryParam == "area") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	db.DBConn.Joins("regions").Clauses(queryClauses...).Order(sort + " " + order).Find(&islands)

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
	if err := db.DBConn.First(&island, id).Error; err != nil {
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
