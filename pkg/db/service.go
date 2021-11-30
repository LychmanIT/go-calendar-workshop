package db

import (
	"calendarWorkshop/models"
	"context"
)

type Service interface {
	GetUser(ctx context.Context, auth *models.Auth) (string, error)
	AddEvent(ctx context.Context, event *models.Event) (int, error)
	AllEvents(ctx context.Context, filters ...models.Filter) ([]models.Event, error)
	ShowEvent(ctx context.Context, ID string) (*models.Event, error)
	UpdateEvent(ctx context.Context, ID string, e *models.Event) (int, error)
	DeleteEvent(ctx context.Context, ID string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
