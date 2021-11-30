package calendar

import (
	authsvc "calendarWorkshop/api/v1/pb/auth"
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/internal/util"
	"calendarWorkshop/models"
	"context"
	"log"
	"net/http"
)

type calendarService struct {
	dbClient   dbsvc.DBClient
	authClient authsvc.AuthClient
}

func NewService(dbClient dbsvc.DBClient, authClient authsvc.AuthClient) Service {
	return &calendarService{dbClient: dbClient, authClient: authClient}
}

//Login receives credentials to log in with and returns bearer token if everything is good or error if there is problems
func (c *calendarService) Login(ctx context.Context, credentials models.Auth) (string, error) {
	cred := util.AuthCredentialsToGRPC(credentials)
	res, err := c.authClient.AuthLogin(ctx, &authsvc.AuthLoginRequest{Credentials: cred})
	if err != nil {
		return "", err
	}
	if res.Err != "" {
		return "", util.ErrUserNotFound
	}
	return res.Token, nil
}

//Logout receives bearer token to log out and returns status 200 if everything is good or status 500/error if there is
//problems
func (c *calendarService) Logout(ctx context.Context, token string) (int, error) {
	res, err := c.authClient.AuthLogout(ctx, &authsvc.AuthLogoutRequest{Token: token})
	if err != nil {
		return http.StatusUnauthorized, err
	}
	return int(res.Status), nil
}

//ChangeTimezone receives bearer token and id to change timezone. Returns status 200 if everything is good or status
//500/error if there is problems
func (c *calendarService) ChangeTimezone(ctx context.Context, token string, tz string) (int, error) {
	res, err := c.authClient.AuthChangeTimezone(ctx, &authsvc.AuthChangeTimezoneRequest{Token: token, Timezone: tz})
	if err != nil {
		return http.StatusUnauthorized, err
	}
	return int(res.Status), nil
}

// IndexEvent receives filters and returns list of the events.
func (c *calendarService) IndexEvent(ctx context.Context, filters ...models.Filter) ([]models.Event, error) {
	grpcFilters := util.FiltersToGRPC(filters)
	res, err := c.dbClient.AllEvents(ctx, &dbsvc.AllEventsRequest{Filters: grpcFilters})
	if err != nil {
		return []models.Event{}, nil
	}
	events := util.GRPCtoEvents(res.Events)
	return events, nil
}

// StoreEvent receives event and store it.
func (c *calendarService) StoreEvent(ctx context.Context, e *models.Event) (int, error) {
	event := util.EventToGRPC(e)
	res, err := c.dbClient.AddEvent(ctx, &dbsvc.AddEventRequest{Event: event})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// ShowEvent receives an ID of the needed event and returns it. If there is no event - returns error.
func (c *calendarService) ShowEvent(ctx context.Context, eventID string) (*models.Event, error) {
	res, err := c.dbClient.ShowEvent(ctx, &dbsvc.ShowEventRequest{ID: eventID})
	if err != nil {
		log.Print(err)
		return &models.Event{}, err
	}
	if res.Event.Id == "" {
		return nil, util.ErrEventNotFound
	}
	event := util.GRPCtoEvent(res.Event)
	return &event, nil
}

// UpdateEvent receives an ID of the needed event and new event instance to update it with. Returns status 200 if everything is good
// and status 404/500 if there is an error.
func (c *calendarService) UpdateEvent(ctx context.Context, eventID string, e *models.Event) (int, error) {
	event := util.EventToGRPC(e)
	res, err := c.dbClient.UpdateEvent(ctx, &dbsvc.UpdateEventRequest{Id: eventID, Event: event})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// DeleteEvent receives an ID of the needed event and deletes it. Returns status 200 if everything is good and status 404/500 if there is
// an error.
func (c *calendarService) DeleteEvent(ctx context.Context, eventID string) (int, error) {
	res, err := c.dbClient.DeleteEvent(ctx, &dbsvc.DeleteEventRequest{Id: eventID})
	if err != nil {
		log.Print(err)
		return http.StatusInternalServerError, err
	}
	return int(res.Status), nil
}

// ServiceStatus returns current service status.
func (c *calendarService) ServiceStatus(ctx context.Context) (int, error) {
	return http.StatusOK, nil
}
