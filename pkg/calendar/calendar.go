package calendar

import (
	"calendarWorkshop/internal/domain/calendar"
	"context"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"net/http"
	"os"
	"time"
)

type calendarService struct {
	//grpcclientDB
}

func NewService() Service { return &calendarService{} }

// IndexEvent ...
func (c *calendarService) IndexEvent(_ context.Context, filters ...calendar.Filter) ([]calendar.Event, error) {

	//grpclientDB.store(smth)

	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return nil, err
	}

	e := calendar.Event{
		ID:          uuid.New().String(),
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
	return []calendar.Event{e}, nil
}

// StoreEvent ...
func (c *calendarService) StoreEvent(ctx context.Context, e *calendar.Event) (int, error) {
	return http.StatusOK, nil
}

// ShowEvent ...
func (c *calendarService) ShowEvent(ctx context.Context, eventID string) (*calendar.Event, error) {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return nil, err
	}

	e := calendar.Event{
		ID:          uuid.New().String(),
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

// UpdateEvent ...
func (c *calendarService) UpdateEvent(ctx context.Context, eventID string, e *calendar.Event) (int, error) {
	return http.StatusAccepted, nil
}

// DeleteEvent ...
func (c *calendarService) DeleteEvent(ctx context.Context, eventID string) (int, error) {
	return http.StatusAccepted, nil
}

// ServiceStatus ...
func (c *calendarService) ServiceStatus(ctx context.Context) (int, error) {
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
