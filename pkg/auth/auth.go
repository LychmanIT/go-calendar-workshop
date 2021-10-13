package auth

import (
	"calendarWorkshop/internal/domain/auth"
	"context"
)

type authService struct{}

func NewService() Service { return &authService{} }

func (a *authService) Login(ctx context.Context, credentials auth.Auth) (*auth.User, error) {
	return &auth.User{}, nil
}
