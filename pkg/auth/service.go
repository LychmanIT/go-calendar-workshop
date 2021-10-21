package auth

import (
	"calendarWorkshop/models"
	"context"
)

type Service interface {
	Login(ctx context.Context, credentials models.Auth) (*models.User, error)
}
