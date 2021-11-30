package auth

import (
	"calendarWorkshop/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, credentials models.Auth) (string, error)
	Logout(ctx context.Context, token string) (int32, error)
	VerifyToken(ctx context.Context, token string) error
	ChangeTimezone(ctx context.Context, token string, tz string) (int32, error)
	GetTimezone(ctx context.Context, token string) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
