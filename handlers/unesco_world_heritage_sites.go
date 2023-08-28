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

// IndexUnescoWorldHeritageSites is a function to get all Unesco World Heritage sites of Portugal data from the database
// @Summary Get all Unesco World Heritage sites of Portugal
// @Description Get all Unesco World Heritage sites of Portugal
// @Tags unesco-world-heritage-sites
// @Accept json
// @Produce json
// @Param region query int false "Filter by region_id (optional)"
// @Param sort_by query string false "Sort by field (optional)"
// @Param order_by query string false "Order by field (optional)"
// @Success 200 {object} ResponseHTTP{data=[]models.UnescoWorldHeritageSite}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/unesco-world-heritage-sites [get]
func IndexUnescoWorldHeritageSites(c *fiber.Ctx) error {
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	inscriptionYearQueryParam := c.Query("inscription_year")
	approvedYearQueryParam := c.Query("approved_year")
	// regionIdQueryParam := c.Query("region_id")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	queryMap := make(map[string]interface{})
	sort := defaultSortByVal
	order := defaultOrderByVal
	unescoWorldHeritageSites := []models.UnescoWorldHeritageSite{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "inscription_year", inscriptionYearQueryParam)
	utils.FilterIntClause(queryClauses, "approved_year", approvedYearQueryParam)
	// utils.FilterIntClause(queryClauses, "area", regionIdQueryParam)

	//
	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "inscription_date" || sortByQueryParam == "approved_date") {
		sort = sortByQueryParam
	}

	//
	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	// finally, make the query to the DB
	db.DBConn.Joins("regions").Clauses(queryClauses...).Order(sort + " " + order).Find(&unescoWorldHeritageSites)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting unesco world heritage sites",
		Data:    unescoWorldHeritageSites,
		Query:   queryMap,
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
	if err := db.DBConn.First(&unescoWorldHeritageSite, id).Error; err != nil {
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
