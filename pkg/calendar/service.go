package calendar

import (
	"calendarWorkshop/models"
	"context"
)

type Service interface {
	IndexEvent(ctx context.Context, filters ...models.Filter) ([]models.Event, error)
	StoreEvent(ctx context.Context, e *models.Event) (int, error)
	ShowEvent(ctx context.Context, eventID string) (*models.Event, error)
	UpdateEvent(ctx context.Context, eventID string, e *models.Event) (int, error)
	DeleteEvent(ctx context.Context, eventID string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
