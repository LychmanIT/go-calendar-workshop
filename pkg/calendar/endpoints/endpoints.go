package endpoints

import (
	domain "calendarWorkshop/models"
	"calendarWorkshop/pkg/calendar"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"net/http"
)

type Set struct {
	LoginEndpoint          endpoint.Endpoint
	LogoutEndpoint         endpoint.Endpoint
	ChangeTimezoneEndpoint endpoint.Endpoint
	IndexEventEndpoint     endpoint.Endpoint
	StoreEventEndpoint     endpoint.Endpoint
	ShowEventEndpoint      endpoint.Endpoint
	UpdateEventEndpoint    endpoint.Endpoint
	DeleteEventEndpoint    endpoint.Endpoint
	ServiceStatusEndpoint  endpoint.Endpoint
}

func NewEndpointSet(svc calendar.Service) Set {
	return Set{
		LoginEndpoint:          MakeLoginEndpoint(svc),
		LogoutEndpoint:         MakeLogoutEndpoint(svc),
		ChangeTimezoneEndpoint: MakeChangeTimezoneEndpoint(svc),
		IndexEventEndpoint:     MakeIndexEventEndpoint(svc),
		StoreEventEndpoint:     MakeStoreEventEndpoint(svc),
		ShowEventEndpoint:      MakeShowEventEndpoint(svc),
		UpdateEventEndpoint:    MakeUpdateEventEndpoint(svc),
		DeleteEventEndpoint:    MakeDeleteEventEndpoint(svc),
		ServiceStatusEndpoint:  MakeServiceStatusEndpoint(svc),
	}
}

func MakeLoginEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		token, err := svc.Login(ctx, req.Credentials)
		if err != nil {
			return LoginResponse{token, err.Error()}, nil
		}
		return LoginResponse{token, ""}, nil
	}
}

func MakeLogoutEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LogoutRequest)
		if req.Err != "" {
			return LogoutResponse{http.StatusInternalServerError, req.Err}, nil
		}
		status, err := svc.Logout(ctx, req.Token)
		if err != nil {
			return LogoutResponse{status, err.Error()}, nil
		}
		return LogoutResponse{status, ""}, nil
	}
}

func MakeChangeTimezoneEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ChangeTimezoneRequest)
		if req.Err != "" {
			return ChangeTimezoneResponse{http.StatusInternalServerError, req.Err}, nil
		}
		status, err := svc.ChangeTimezone(ctx, req.Token, req.Timezone)
		if err != nil {
			return ChangeTimezoneResponse{status, err.Error()}, nil
		}
		return ChangeTimezoneResponse{status, ""}, nil
	}
}

func MakeIndexEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IndexEventRequest)
		if req.Err != "" {
			return IndexEventResponse{nil, req.Err}, nil
		}
		events, err := svc.IndexEvent(ctx, req.Filters...)
		if err != nil {
			return IndexEventResponse{events, err.Error()}, nil
		}
		return IndexEventResponse{events, ""}, nil
	}
}

func MakeStoreEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StoreEventRequest)
		if req.Err != "" {
			return StoreEventResponse{http.StatusInternalServerError, req.Err}, nil
		}
		status, err := svc.StoreEvent(ctx, req.Event)
		if err != nil {
			return StoreEventResponse{status, err.Error()}, nil
		}
		return StoreEventResponse{status, ""}, nil
	}
}

func MakeShowEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ShowEventRequest)
		if req.Err != "" {
			return ShowEventResponse{nil, req.Err}, nil
		}
		event, err := svc.ShowEvent(ctx, req.EventID)
		if err != nil {
			return ShowEventResponse{event, err.Error()}, nil
		}
		return ShowEventResponse{event, ""}, nil
	}
}

func MakeUpdateEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEventRequest)
		if req.Err != "" {
			return UpdateEventResponse{http.StatusInternalServerError, req.Err}, nil
		}
		status, err := svc.UpdateEvent(ctx, req.EventID, req.Event)
		if err != nil {
			return UpdateEventResponse{status, err.Error()}, nil
		}
		return UpdateEventResponse{status, ""}, nil
	}
}

func MakeDeleteEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteEventRequest)
		if req.Err != "" {
			return DeleteEventResponse{http.StatusInternalServerError, req.Err}, nil
		}
		status, err := svc.DeleteEvent(ctx, req.EventID)
		if err != nil {
			return DeleteEventResponse{status, err.Error()}, nil
		}
		return DeleteEventResponse{status, ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{code, err.Error()}, nil
		}
		return ServiceStatusResponse{code, ""}, nil
	}
}

func (s *Set) Login(ctx context.Context, credentials domain.Auth) (string, error) {
	resp, err := s.LoginEndpoint(ctx, LoginRequest{Credentials: credentials})
	if err != nil {
		return "", errors.New("invalid credentials")
	}
	getResp := resp.(LoginResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Token, nil
}

func (s *Set) Logout(ctx context.Context, token string) (int, error) {
	resp, err := s.LogoutEndpoint(ctx, LogoutRequest{Token: token})
	if err != nil {
		return http.StatusUnauthorized, err
	}
	getResp := resp.(LogoutResponse)
	if getResp.Err != "" {
		return getResp.Status, errors.New(getResp.Err)
	}
	return getResp.Status, nil
}

func (s *Set) ChangeTimezone(ctx context.Context, token string, tz string) (int, error) {
	resp, err := s.ChangeTimezoneEndpoint(ctx, ChangeTimezoneRequest{Token: token, Timezone: tz})
	if err != nil {
		return http.StatusUnauthorized, err
	}
	getResp := resp.(ChangeTimezoneResponse)
	if getResp.Err != "" {
		return getResp.Status, errors.New(getResp.Err)
	}
	return getResp.Status, nil
}

func (s *Set) IndexEvent(ctx context.Context, filters ...domain.Filter) ([]domain.Event, error) {
	resp, err := s.IndexEventEndpoint(ctx, IndexEventRequest{Filters: filters})
	if err != nil {
		return []domain.Event{}, err
	}
	getResp := resp.(IndexEventResponse)
	if getResp.Err != "" {
		return []domain.Event{}, errors.New(getResp.Err)
	}
	return getResp.Events, nil
}

func (s *Set) StoreEvent(ctx context.Context, e *domain.Event) (int, error) {
	resp, err := s.StoreEventEndpoint(ctx, StoreEventRequest{Event: e})
	if err != nil {
		return http.StatusBadRequest, err
	}
	storeResp := resp.(StoreEventResponse)
	if storeResp.Err != "" {
		return http.StatusBadRequest, errors.New(storeResp.Err)
	}
	return storeResp.Status, nil
}

func (s *Set) ShowEvent(ctx context.Context, eventID string) (*domain.Event, error) {
	resp, err := s.ShowEventEndpoint(ctx, ShowEventRequest{EventID: eventID})
	if err != nil {
		return &domain.Event{}, err
	}
	showResp := resp.(ShowEventResponse)
	if showResp.Err != "" {
		return &domain.Event{}, errors.New(showResp.Err)
	}
	return showResp.Event, nil
}

func (s *Set) UpdateEvent(ctx context.Context, eventID string, e *domain.Event) (int, error) {
	resp, err := s.UpdateEventEndpoint(ctx, UpdateEventRequest{EventID: eventID, Event: e})
	if err != nil {
		return http.StatusBadRequest, err
	}
	updateResp := resp.(UpdateEventResponse)
	if updateResp.Err != "" {
		return http.StatusBadRequest, errors.New(updateResp.Err)
	}
	return http.StatusOK, nil
}

func (s *Set) DeleteEvent(ctx context.Context, eventID string) (int, error) {
	resp, err := s.DeleteEventEndpoint(ctx, DeleteEventRequest{EventID: eventID})
	if err != nil {
		return http.StatusBadRequest, err
	}
	deleteResp := resp.(DeleteEventResponse)
	if deleteResp.Err != "" {
		return http.StatusBadRequest, errors.New(deleteResp.Err)
	}
	return http.StatusOK, nil
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
