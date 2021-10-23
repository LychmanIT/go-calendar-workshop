package db

import (
	"calendarWorkshop/models"
	"strings"
)

func JoinQueryFilter(query string, filters []models.Filter) string {
	query += " WHERE "
	for key, filter := range filters {
		query += filter.Key
		query += " = "
		query += "'"
		query += filter.Value
		query += "' "
		if key != (len(filters) - 1) {
			query += "AND "
		}
	}

	strings.TrimSuffix(query, "AND ")

	return query
}
