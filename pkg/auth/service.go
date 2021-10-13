package auth

import (
	"calendarWorkshop/internal/domain/auth"
	"context"
)

type Service interface {
	Login(ctx context.Context, credentials auth.Auth) (*auth.User, error)
}
