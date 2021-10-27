package db

import (
	"calendarWorkshop/models"
	"github.com/jackc/pgx/v4"
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

func IsRowsAffected(row pgx.Rows) bool {
	counter := 0
	for row.Next() {
		counter++
	}
	if counter == 0 {
		return false
	}
	return true
}
