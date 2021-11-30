package calendar

import (
	"calendarWorkshop/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, credentials models.Auth) (string, error)
	Logout(ctx context.Context, token string) (int, error)
	ChangeTimezone(ctx context.Context, token string, tz string) (int, error)
	IndexEvent(ctx context.Context, filters ...models.Filter) ([]models.Event, error)
	StoreEvent(ctx context.Context, e *models.Event) (int, error)
	ShowEvent(ctx context.Context, eventID string) (*models.Event, error)
	UpdateEvent(ctx context.Context, eventID string, e *models.Event) (int, error)
	DeleteEvent(ctx context.Context, eventID string) (int, error)
	ServiceStatus(ctx context.Context) (int, error)
}
