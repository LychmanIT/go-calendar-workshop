package calendar

import (
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/internal/domain/calendar"
	"calendarWorkshop/internal/util"
	"calendarWorkshop/models"
	"context"
	"log"
	"net/http"
)

type calendarService struct {
	dbClient dbsvc.DBClient
}

func NewService(client dbsvc.DBClient) Service { return &calendarService{dbClient: client} }

// IndexEvent receives filters and returns list of the events
func (c *calendarService) IndexEvent(ctx context.Context, filters ...models.Filter) ([]models.Event, error) {
	grpcFilters := calendar.FiltersToGRPC(filters)
	res, err := c.dbClient.AllEvents(ctx, &dbsvc.AllEventsRequest{Filters: grpcFilters})
	if err != nil {
		log.Print(err)
		return []models.Event{}, nil
	}
	events := calendar.GRPCtoEvents(res.Events)
	return events, nil
}

// StoreEvent receives event and store it
func (c *calendarService) StoreEvent(ctx context.Context, e *models.Event) (int, error) {
	event := calendar.EventToGRPC(e)
	res, err := c.dbClient.AddEvent(ctx, &dbsvc.AddEventRequest{Event: event})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// ShowEvent receives an ID of the needed event and returns it. If there is no event - returns error
func (c *calendarService) ShowEvent(ctx context.Context, eventID string) (*models.Event, error) {
	res, err := c.dbClient.ShowEvent(ctx, &dbsvc.ShowEventRequest{ID: eventID})
	if err != nil {
		log.Print(err)
		return &models.Event{}, err
	}
	if res.Event.Id == "" {
		return nil, util.ErrEventNotFound
	}
	event := calendar.GRPCtoEvent(res.Event)
	return &event, nil
}

// UpdateEvent receives an ID of the needed event and new event instance to update it with. Returns status 200 if everything is good
// and status 404/500 if there is an error
func (c *calendarService) UpdateEvent(ctx context.Context, eventID string, e *models.Event) (int, error) {
	event := calendar.EventToGRPC(e)
	res, err := c.dbClient.UpdateEvent(ctx, &dbsvc.UpdateEventRequest{Id: eventID, Event: event})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// DeleteEvent receives an ID of the needed event and deletes it. Returns status 200 if everything is good and status 404/500 if there is
// an error
func (c *calendarService) DeleteEvent(ctx context.Context, eventID string) (int, error) {
	res, err := c.dbClient.DeleteEvent(ctx, &dbsvc.DeleteEventRequest{Id: eventID})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// ServiceStatus returns current service status
func (c *calendarService) ServiceStatus(ctx context.Context) (int, error) {
	return http.StatusOK, nil
}
