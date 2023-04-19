package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"gorm.io/gorm/clause"
)

type ResponseHTTP struct {
	Success bool                   `json:"success"`
	Data    interface{}            `json:"data"`
	Message string                 `json:"message"`
	Query   map[string]interface{} `json:"query"`
}

func CheckError(msg string, err error) {
	if err != nil {
		log.Printf("❗ ERROR DETECTED ❗ %s => %s", msg, err)
	}
}

func FailOnError(msg string, err error) {
	if err != nil {
		log.Fatalf(fmt.Sprintf("❗❗ PANIC ❗❗ %s => %s", msg, err))
	}
}

func FilterIntClause(clauses []clause.Expression, key string, param string) ([]clause.Expression, error) {
	if param == "" {
		return clauses, fmt.Errorf("empty param: %s", param)
	}

	splitParam := strings.Split(param, ":")

	if len(splitParam) != 2 {
		return clauses, fmt.Errorf("invalid param format: %s", param)
	}

	filterType := splitParam[0]

	valInt, err := strconv.Atoi(splitParam[1])
	if err != nil {
		return clauses, fmt.Errorf("invalid param format: %s", param)
	}

	switch filterType {
	case "lt":
		clauses = append(clauses, clause.Lt{Column: key, Value: valInt})
	case "lte":
		clauses = append(clauses, clause.Lte{Column: key, Value: valInt})
	case "gt":
		clauses = append(clauses, clause.Gt{Column: key, Value: valInt})
	case "gte":
		clauses = append(clauses, clause.Gte{Column: key, Value: valInt})
	default:
		return clauses, fmt.Errorf("invalid filter type: %s", filterType)
	}

	return clauses, nil
}

func FilterStrClause(clauses []clause.Expression, key string, param string) ([]clause.Expression, error) {
	if param == "" {
		return clauses, fmt.Errorf("empty param: %s", param)
	}

	splitParam := strings.Split(param, ":")

	if len(splitParam) != 2 {
		return clauses, fmt.Errorf("invalid param format: %s", param)
	}

	filterType := splitParam[0]

	valInt, err := strconv.Atoi(splitParam[1])
	if err != nil {
		return clauses, fmt.Errorf("invalid param format: %s", param)
	}

	switch filterType {
	case "lt":
		clauses = append(clauses, clause.Lt{Column: key, Value: valInt})
	case "lte":
		clauses = append(clauses, clause.Lte{Column: key, Value: valInt})
	case "gt":
		clauses = append(clauses, clause.Gt{Column: key, Value: valInt})
	case "gte":
		clauses = append(clauses, clause.Gte{Column: key, Value: valInt})
	default:
		return clauses, fmt.Errorf("invalid filter type: %s", filterType)
	}

	return clauses, nil
}
