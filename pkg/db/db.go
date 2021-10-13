package db

import (
	"calendarWorkshop/internal/domain/calendar"
	"context"
	"net/http"
	"time"
)

type dbService struct{}

func NewService() Service { return &dbService{} }

func (d *dbService) GetUser(ctx context.Context, ID string) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) AddEvent(_ context.Context, e *calendar.Event) (string, error) {
	return "", nil
}

func (d *dbService) AllEvents(_ context.Context, filters ...calendar.Filter) ([]calendar.Event, error) {
	return []calendar.Event{}, nil
}

func (d *dbService) ShowEvent(_ context.Context, ID string) (*calendar.Event, error) {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return nil, err
	}

	e := calendar.Event{
		ID:          ID,
		Title:       "New Event",
		Description: "It`s description",
		Time:        time.Now().In(loc).String(),
		Timezone:    loc.String(),
		Duration:    2,
		Notes: []string{
			"Note1",
			"Note2",
			"Note3",
		},
	}
	return &e, nil
}

func (d *dbService) UpdateEvent(_ context.Context, ID string, e *calendar.Event) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) DeleteEvent(_ context.Context, ID string) (int, error) {
	return http.StatusOK, nil
}

func (d *dbService) ServiceStatus(_ context.Context) (int, error) {
	return http.StatusOK, nil
}
