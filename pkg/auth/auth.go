package auth

import (
	"calendarWorkshop/models"
	"context"
)

type authService struct{}

func NewService() Service { return &authService{} }

func (a *authService) Login(ctx context.Context, credentials models.Auth) (*models.User, error) {
	return &models.User{}, nil
}
