package endpoints

import (
	"calendarWorkshop/models"
	"calendarWorkshop/pkg/auth"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type Set struct {
	LoginEndpoint          endpoint.Endpoint
	LogoutEndpoint         endpoint.Endpoint
	VerifyTokenEndpoint    endpoint.Endpoint
	ChangeTimezoneEndpoint endpoint.Endpoint
	GetTimezoneEndpoint    endpoint.Endpoint
	ServiceStatusEndpoint  endpoint.Endpoint
}

func NewEndpointSet(svc auth.Service) Set {
	return Set{
		LoginEndpoint:          MakeLoginEndpoint(svc),
		LogoutEndpoint:         MakeLogoutEndpoint(svc),
		VerifyTokenEndpoint:    MakeVerifyTokenEndpoint(svc),
		ChangeTimezoneEndpoint: MakeChangeTimezoneEndpoint(svc),
		GetTimezoneEndpoint:    MakeGetTimezoneEndpoint(svc),
		ServiceStatusEndpoint:  MakeServiceStatusEndpoint(svc),
	}
}

func MakeLoginEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := svc.Login(ctx, req.Credentials)
		if err != nil {
			return LoginResponse{token, err.Error()}, nil
		}
		return LoginResponse{token, ""}, nil
	}
}

func MakeLogoutEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogoutRequest)
		status, err := svc.Logout(ctx, req.Token)
		if err != nil {
			return LogoutResponse{status, err.Error()}, nil
		}
		return LogoutResponse{status, ""}, nil
	}
}

func MakeVerifyTokenEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(VerifyTokenRequest)
		err := svc.VerifyToken(ctx, req.Token)
		if err != nil {
			return VerifyTokenResponse{err.Error()}, nil
		}
		return VerifyTokenResponse{Err: ""}, nil
	}
}

func MakeChangeTimezoneEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeTimezoneRequest)
		status, err := svc.ChangeTimezone(ctx, req.Token, req.Timezone)
		if err != nil {
			return ChangeTimezoneResponse{status, err.Error()}, nil
		}
		return ChangeTimezoneResponse{status, ""}, nil
	}
}

func MakeGetTimezoneEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTimezoneRequest)
		timezone, err := svc.GetTimezone(ctx, req.Token)
		if err != nil {
			return GetTimezoneResponse{timezone, err.Error()}, nil
		}
		return GetTimezoneResponse{timezone, ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{code, err.Error()}, nil
		}
		return ServiceStatusResponse{code, ""}, nil
	}
}

func (s *Set) Login(ctx context.Context, credentials models.Auth) (string, error) {
	resp, err := s.LoginEndpoint(ctx, LoginRequest{Credentials: credentials})
	if err != nil {
		return "", err
	}
	getResp := resp.(LoginResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Token, nil
}

func (s *Set) Logout(ctx context.Context, token string) (int32, error) {
	resp, err := s.LogoutEndpoint(ctx, LogoutRequest{Token: token})
	if err != nil {
		return http.StatusBadRequest, err
	}
	getResp := resp.(LogoutResponse)
	if getResp.Err != "" {
		return http.StatusInternalServerError, errors.New(getResp.Err)
	}
	return getResp.Status, nil
}

func (s *Set) VerifyToken(ctx context.Context, token string) error {
	resp, err := s.VerifyTokenEndpoint(ctx, VerifyTokenRequest{Token: token})
	if err != nil {
		return err
	}
	getResp := resp.(VerifyTokenResponse)
	if getResp.Err != "" {
		return errors.New(getResp.Err)
	}
	return nil
}

func (s *Set) ChangeTimezone(ctx context.Context, token string, tz string) (int32, error) {
	resp, err := s.ChangeTimezoneEndpoint(ctx, ChangeTimezoneRequest{Token: token, Timezone: tz})
	if err != nil {
		return http.StatusBadRequest, err
	}
	getResp := resp.(ChangeTimezoneResponse)
	if getResp.Err != "" {
		return http.StatusInternalServerError, errors.New(getResp.Err)
	}
	return http.StatusOK, nil
}

func (s *Set) GetTimezone(ctx context.Context, token string) (string, error) {
	resp, err := s.GetTimezoneEndpoint(ctx, GetTimezoneRequest{Token: token})
	if err != nil {
		return "", err
	}
	getResp := resp.(GetTimezoneResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Timezone, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Status, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Status, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Status, nil
}
