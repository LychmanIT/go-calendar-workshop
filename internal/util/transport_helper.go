package util

import (
	"calendarWorkshop/models"
	"net/url"
)

func MapToFilters(req url.Values) []models.Filter {
	var result []models.Filter

	for key, value := range req {
		f := models.Filter{
			Key:   key,
			Value: value[0],
		}
		result = append(result, f)
	}

	return result
}
