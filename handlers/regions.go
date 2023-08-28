package handlers

import (
	"fmt"
	"josepsoares/iberiapi/db"
	"josepsoares/iberiapi/models"
	"josepsoares/iberiapi/utils"
	"net/http"

	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
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
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	autonomousQueryParam := c.Query("autonomous")
	populationQueryParam := c.Query("population")
	populationDensityQueryParam := c.Query("population_density")
	areaQueryParam := c.Query("area")
	districtsQueryParam := c.Query("districts")
	municipalitiesQueryParam := c.Query("municipalities")
	freguesiasQueryParam := c.Query("freguesias")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	regions := []models.Region{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "population", populationQueryParam)
	utils.FilterIntClause(queryClauses, "population_density", populationDensityQueryParam)
	utils.FilterIntClause(queryClauses, "area", areaQueryParam)
	utils.FilterIntClause(queryClauses, "districts", districtsQueryParam)
	utils.FilterIntClause(queryClauses, "municipalities", municipalitiesQueryParam)
	utils.FilterIntClause(queryClauses, "freguesias", freguesiasQueryParam)

	if autonomousQueryParam != "" && (autonomousQueryParam == "true" || autonomousQueryParam == "false") {
		val, err := strconv.ParseBool(autonomousQueryParam)

		if err != nil {
			return c.Status(500).JSON(utils.ResponseHTTP{
				Success: false,
				Message: "Error filtering by autonomous parameter",
			})
		}

		if val {
			queryClauses = append(queryClauses, clause.Like{Column: "autonomous", Value: 0})
		} else {
			queryClauses = append(queryClauses, clause.Gte{Column: "autonomous", Value: 1})
		}
	}

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "population" || sortByQueryParam == "population_density" || sortByQueryParam == "area" || sortByQueryParam == "nr_districts" || sortByQueryParam == "nr_municipalities" || sortByQueryParam == "nr_freguesias") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	db.DBConn.Clauses(queryClauses...).Order(sort + " " + order).Find(&regions)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting regions",
		// Data:    regions,
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
	if err := db.DBConn.First(&region, id).Error; err != nil {
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
