package auth

import (
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/internal/domain/auth"
	"calendarWorkshop/internal/util"
	"calendarWorkshop/models"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

type authService struct {
	redisClient *redis.Client
	dbClient    dbsvc.DBClient
}

func NewService(client *redis.Client, dbGRPCClient dbsvc.DBClient) Service {
	return &authService{redisClient: client, dbClient: dbGRPCClient}
}

//Login receives Credentials to log in with and returns bearer token if everything is good. Otherwise, returns error.
func (a *authService) Login(ctx context.Context, credentials models.Auth) (string, error) {
	cred := util.CredentialsToGRPC(credentials)
	res, err := a.dbClient.GetUser(ctx, &dbsvc.GetUserRequest{Auth: cred})

	if err != nil || res.Err != "" {
		return "", errors.New(res.Err)
	}

	ts, err := auth.CreateToken(res.Id)
	if err != nil {
		return "", err
	}

	at := time.Unix(ts.AtExpires, 0)
	now := time.Now()

	errAccess := a.redisClient.Set(ctx, ts.AccessUuid, res.Id, at.Sub(now)).Err()
	if errAccess != nil {
		return "", errAccess
	}

	return ts.AccessToken, nil
}

//Logout receives bearer token to log out with and returns status
func (a *authService) Logout(ctx context.Context, token string) (int32, error) {
	metadata, err := auth.ExtractTokenMetadata(token)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	deletedAt, err := a.redisClient.Del(ctx, metadata.AccessUUID).Result()
	if err != nil || deletedAt != 1 {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//VerifyToken receives bearer token to check. Returns error if there is problems.
func (a *authService) VerifyToken(ctx context.Context, token string) error {
	metadata, err := auth.ExtractTokenMetadata(token)
	if err != nil {
		return err
	}

	userid, err := a.redisClient.Get(ctx, metadata.AccessUUID).Result()
	if err != nil {
		return err
	}

	if metadata.UserID != userid {
		return util.ErrUnauthorized
	}
	return nil
}

//ChangeTimezone receives user`s bearer token to update his timezone and returns status
func (a *authService) ChangeTimezone(ctx context.Context, token string, tz string) (int32, error) {
	metadata, err := auth.ExtractTokenMetadata(token)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	timezone := metadata.AccessUUID + "_timezone"

	_, err = a.redisClient.Set(ctx, timezone, tz, 0).Result()
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

//GetTimezone receives user`s bearer token and returns his timezone
func (a *authService) GetTimezone(ctx context.Context, token string) (string, error) {
	metadata, err := auth.ExtractTokenMetadata(token)
	if err != nil {
		return "", err
	}

	timezone := metadata.AccessUUID + "_timezone"

	value, err := a.redisClient.Get(ctx, timezone).Result()
	if err == redis.Nil {
		_, err = a.redisClient.Set(ctx, timezone, "America/Chicago", 0).Result()
		if err != nil {
			return "", err
		}
	}

	return value, nil
}

func (a *authService) ServiceStatus(ctx context.Context) (int, error) {
	return http.StatusOK, nil
}
