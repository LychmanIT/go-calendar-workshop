package calendar

import (
	"calendarWorkshop/internal/domain/calendar"
	"context"
)

type Service interface {
	IndexEvent(ctx context.Context, filters ...calendar.Filter) ([]calendar.Event, error)
	StoreEvent(ctx context.Context, e *calendar.Event) (int, error)
	ShowEvent(ctx context.Context, eventID string) (*calendar.Event, error)
	UpdateEvent(ctx context.Context, eventID string, e *calendar.Event) (int, error)
	DeleteEvent(ctx context.Context, eventID string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
