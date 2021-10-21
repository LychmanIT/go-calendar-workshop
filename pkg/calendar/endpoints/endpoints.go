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
	IndexEventEndpoint    endpoint.Endpoint
	StoreEventEndpoint    endpoint.Endpoint
	ShowEventEndpoint     endpoint.Endpoint
	UpdateEventEndpoint   endpoint.Endpoint
	DeleteEventEndpoint   endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc calendar.Service) Set {
	return Set{
		IndexEventEndpoint:    MakeIndexEventEndpoint(svc),
		StoreEventEndpoint:    MakeStoreEventEndpoint(svc),
		ShowEventEndpoint:     MakeShowEventEndpoint(svc),
		UpdateEventEndpoint:   MakeUpdateEventEndpoint(svc),
		DeleteEventEndpoint:   MakeDeleteEventEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeIndexEventEndpoint(svc calendar.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(IndexEventRequest)
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
