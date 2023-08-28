package handlers

import (
	"josepsoares/iberiapi/db"
	"josepsoares/iberiapi/models"
	"josepsoares/iberiapi/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zeimedee/go-postgres/database"
	"gorm.io/gorm/clause"
)

// IndexCountries is a function to get all countries data from the database
// @Summary Get all countries
// @Description Get all countries
// @Tags countries
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Island}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/countries [get]
func IndexCountries(c *fiber.Ctx) error {
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")

	populationQueryParam := c.Query("population")
	populationDensityQueryParam := c.Query("population_density")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	countries := []models.Island{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterIntClause(queryClauses, "population", populationQueryParam)
	utils.FilterIntClause(queryClauses, "population_density", populationDensityQueryParam)

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "population" || sortByQueryParam == "population_density") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	db.DBConn.Joins("regions").Clauses(queryClauses...).Order(sort + " " + order).Find(&countries)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting countries",
		Data:    countries,
	})
}

// GetCountryByID is a function to get all general data from a specific country of the database
// @Summary Get general data of Country
// @Description Get general data of Country
// @Tags country
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.Country}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/ [get]
func GetCountryByID(c *fiber.Ctx) error {
	id := c.Params("id")

	country := new(models.Country)
	if err := database.DB.Db.First(&country, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: "Country info not found",
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
		Message: "Success getting Country info",
		Data:    country,
	})
}
