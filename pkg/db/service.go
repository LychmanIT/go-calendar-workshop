package db

import (
	"calendarWorkshop/internal/domain/calendar"
	"context"
)

type Service interface {
	GetUser(ctx context.Context, ID string) (int, error)
	AddEvent(ctx context.Context, event *calendar.Event) (string, error)
	AllEvents(ctx context.Context, filters ...calendar.Filter) ([]calendar.Event, error)
	ShowEvent(ctx context.Context, ID string) (*calendar.Event, error)
	UpdateEvent(ctx context.Context, ID string, e *calendar.Event) (int, error)
	DeleteEvent(ctx context.Context, ID string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
