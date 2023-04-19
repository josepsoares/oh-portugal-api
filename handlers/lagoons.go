package handlers

import (
	"fmt"
	"josepsoares/oh-portugal-api/db"
	"josepsoares/oh-portugal-api/models"
	"josepsoares/oh-portugal-api/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
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
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	areaQueryParam := c.Query("area")
	depthQueryParam := c.Query("depth")
	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	lagoons := []models.Lagoon{}

	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterIntClause(queryClauses, "area", areaQueryParam)
	utils.FilterIntClause(queryClauses, "depth", depthQueryParam)

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "area" || sortByQueryParam == "depth") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	db.DBConn.Joins("lagoons").Clauses(queryClauses...).Order(sort + " " + order).Find(&lagoons)

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
	if err := db.DBConn.First(&lagoon, id).Error; err != nil {
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
