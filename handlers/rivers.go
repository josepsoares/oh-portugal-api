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

// IndexRivers is a function to get all rivers data from the database
// @Summary Get all rivers
// @Description Get all rivers
// @Tags rivers
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.River}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/rivers [get]
func IndexRivers(c *fiber.Ctx) error {
	// TODO => define a select object for queries

	// default vars
	defaultSortByVal := "name"
	defaultOrderByVal := "asc"

	// query params vars
	idQueryParam := c.Query("id")
	nameQueryParam := c.Query("name")
	nationalQueryParam := c.Query("national")
	lengthQueryParam := c.Query("length")
	avgFlowQueryParam := c.Query("average_flow")
	watershedAreaQueryParam := c.Query("watershed_area")
	// regionIdQueryParam := c.QueryInt("region_id")

	sortByQueryParam := c.Query("sort_by")
	orderByQueryParam := c.Query("order_by")

	// mutable vars
	queryClauses := make([]clause.Expression, 0)
	sort := defaultSortByVal
	order := defaultOrderByVal
	rivers := []models.River{}

	utils.FilterIntClause(queryClauses, "id", idQueryParam)
	utils.FilterStrClause(queryClauses, "name", nameQueryParam)
	utils.FilterIntClause(queryClauses, "length", lengthQueryParam)
	utils.FilterIntClause(queryClauses, "average_flow", avgFlowQueryParam)
	utils.FilterIntClause(queryClauses, "watershed_area", watershedAreaQueryParam)

	if nationalQueryParam != "" && (nationalQueryParam == "true" || nationalQueryParam == "false") {
		val, err := strconv.ParseBool(nationalQueryParam)

		if err != nil {
			return c.Status(500).JSON(utils.ResponseHTTP{
				Success: false,
				Message: "Error filtering by national parameter",
			})
		}

		if val {
			queryClauses = append(queryClauses, clause.Like{Column: "national", Value: 0})
		} else {
			queryClauses = append(queryClauses, clause.Gte{Column: "national", Value: 1})
		}
	}

	if sortByQueryParam != "" && (sortByQueryParam == "name" || sortByQueryParam == "id" || sortByQueryParam == "length" || sortByQueryParam == "average_flow" || sortByQueryParam == "watershed_area") {
		sort = sortByQueryParam
	}

	if orderByQueryParam != "" && (orderByQueryParam == "desc" || orderByQueryParam == "asc") {
		order = orderByQueryParam
	}

	/*
		  if regionIdQueryParam != 0 {
				db.DBConn.Joins("regions").Where("RegionID = ?", regionIdQueryParam).Order(sort + " " + order).Find(&rivers)

				return c.Status(200).JSON(utils.ResponseHTTP{
					Success: true,
					Message: "Success getting rivers filtered by region",
					Data:    rivers,
				})
			}
	*/

	db.DBConn.Joins("regions").Clauses(queryClauses...).Order(sort + " " + order).Find(&rivers)

	return c.Status(200).JSON(utils.ResponseHTTP{
		Success: true,
		Message: "Success getting rivers sites",
		Data:    rivers,
	})
}

// GetRiverByID is a function to get a river by ID
// @Summary Get river by ID
// @Description Get river by ID
// @Tags rivers
// @Accept json
// @Produce json
// @Param id path int true "River ID"
// @Success 200 {object} ResponseHTTP{data=[]models.River}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/rivers/{id} [get]
func GetRiverByID(c *fiber.Ctx) error {
	id := c.Params("id")

	river := new(models.River)
	if err := db.DBConn.First(&river, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(utils.ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("River with ID %v not found", id),
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
		Message: "Success getting river by ID",
		Data:    *river,
	})
}
