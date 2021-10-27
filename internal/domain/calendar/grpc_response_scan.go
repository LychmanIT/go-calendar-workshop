package calendar

import (
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/models"
)

func GRPCtoEvents(events []*dbsvc.Event) []models.Event {
	var result []models.Event

	for _, event := range events {
		e := models.Event{
			ID:          event.Id,
			Title:       event.Title,
			Description: event.Description,
			Time:        event.Time,
			Timezone:    event.Timezone,
			Duration:    event.Duration,
			Notes:       event.Notes,
		}
		result = append(result, e)
	}

	return result
}

func GRPCtoEvent(event *dbsvc.Event) models.Event {
	return models.Event{
		ID:          event.Id,
		Title:       event.Title,
		Description: event.Description,
		Time:        event.Time,
		Timezone:    event.Timezone,
		Duration:    event.Duration,
		Notes:       event.Notes,
	}
}

func EventToGRPC(event *models.Event) *dbsvc.Event {
	return &dbsvc.Event{
		Id:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Time:        event.Time,
		Timezone:    event.Timezone,
		Duration:    event.Duration,
		Notes:       event.Notes,
	}
}

func FiltersToGRPC(filters []models.Filter) []*dbsvc.AllEventsRequest_Filters {
	var result []*dbsvc.AllEventsRequest_Filters

	for _, filter := range filters {
		f := &dbsvc.AllEventsRequest_Filters{
			Key:   filter.Key,
			Value: filter.Value,
		}
		result = append(result, f)
	}

	return result
}
