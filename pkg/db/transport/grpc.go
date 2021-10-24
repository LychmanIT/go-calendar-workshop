package transport

import (
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/models"
	"calendarWorkshop/pkg/db/endpoints"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	getUser       grpctransport.Handler
	allEvents     grpctransport.Handler
	addEvent      grpctransport.Handler
	showEvent     grpctransport.Handler
	updateEvent   grpctransport.Handler
	deleteEvent   grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) dbsvc.DBServer {
	return &grpcServer{
		getUser: grpctransport.NewServer(
			ep.GetUserEndpoint,
			decodeGRPCGetUserRequest,
			encodeGRPCGetUserResponse,
		),
		allEvents: grpctransport.NewServer(
			ep.AllEventsEndpoint,
			decodeGRPCAllEventsRequest,
			encodeGRPCAllEventsResponse,
		),
		addEvent: grpctransport.NewServer(
			ep.AddEventEndpoint,
			decodeGRPCAddEventRequest,
			encodeGRPCAddEventResponse,
		),
		showEvent: grpctransport.NewServer(
			ep.ShowEventEndpoint,
			decodeGRPCShowEventRequest,
			encodeGRPCShowEventResponse,
		),
		updateEvent: grpctransport.NewServer(
			ep.UpdateEventEndpoint,
			decodeGRPCUpdateEventRequest,
			encodeGRPCUpdateEventResponse,
		),
		deleteEvent: grpctransport.NewServer(
			ep.DeleteEventEndpoint,
			decodeGRPCDeleteEventRequest,
			encodeGRPCDeleteEventResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			encodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) GetUser(ctx context.Context, r *dbsvc.GetUserRequest) (*dbsvc.GetUserReply, error) {
	_, rep, err := g.getUser.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.GetUserReply), nil
}

func (g *grpcServer) AllEvents(ctx context.Context, r *dbsvc.AllEventsRequest) (*dbsvc.AllEventsReply, error) {
	_, rep, err := g.allEvents.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.AllEventsReply), nil
}

func (g *grpcServer) AddEvent(ctx context.Context, r *dbsvc.AddEventRequest) (*dbsvc.AddEventReply, error) {
	_, rep, err := g.addEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.AddEventReply), nil
}

func (g *grpcServer) ShowEvent(ctx context.Context, r *dbsvc.ShowEventRequest) (*dbsvc.ShowEventReply, error) {
	_, rep, err := g.showEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.ShowEventReply), nil
}

func (g *grpcServer) UpdateEvent(ctx context.Context, r *dbsvc.UpdateEventRequest) (*dbsvc.UpdateEventReply, error) {
	_, rep, err := g.updateEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.UpdateEventReply), nil
}

func (g *grpcServer) DeleteEvent(ctx context.Context, r *dbsvc.DeleteEventRequest) (*dbsvc.DeleteEventReply, error) {
	_, rep, err := g.deleteEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.DeleteEventReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *dbsvc.ServiceStatusRequest) (*dbsvc.ServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*dbsvc.ServiceStatusReply), nil
}

func decodeGRPCGetUserRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.GetUserRequest)
	auth := &models.Auth{
		Username: req.Auth.Username,
		Password: req.Auth.Password,
	}
	return endpoints.GetUserRequest{Auth: auth}, nil
}

func decodeGRPCAllEventsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.AllEventsRequest)
	var filters []models.Filter
	for _, f := range req.Filters {
		filters = append(filters, models.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.AllEventsRequest{Filters: filters}, nil
}

func decodeGRPCAddEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.AddEventRequest)
	var notes []string
	for _, note := range req.Event.Notes {
		notes = append(notes, note)
	}
	event := &models.Event{
		ID:          req.Event.Id,
		Title:       req.Event.Title,
		Description: req.Event.Description,
		Time:        req.Event.Time,
		Timezone:    req.Event.Timezone,
		Duration:    req.Event.Duration,
		Notes:       notes,
	}
	return endpoints.AddEventRequest{Event: event}, nil
}

func decodeGRPCShowEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.ShowEventRequest)
	return endpoints.ShowEventRequest{EventID: req.ID}, nil
}

func decodeGRPCUpdateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.UpdateEventRequest)
	var notes []string
	for _, note := range req.Event.Notes {
		notes = append(notes, note)
	}
	event := &models.Event{
		ID:          req.Event.Id,
		Title:       req.Event.Title,
		Description: req.Event.Description,
		Time:        req.Event.Time,
		Timezone:    req.Event.Timezone,
		Duration:    req.Event.Duration,
		Notes:       notes,
	}
	return endpoints.UpdateEventRequest{EventID: req.Id, Event: event}, nil
}

func decodeGRPCDeleteEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*dbsvc.DeleteEventRequest)
	return endpoints.DeleteEventRequest{EventID: req.Id}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func encodeGRPCGetUserResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.GetUserResponse)
	return &dbsvc.GetUserReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCAllEventsResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.AllEventsResponse)
	var events []*dbsvc.Event
	for _, e := range reply.Events {
		var notes []string
		for _, note := range e.Notes {
			notes = append(notes, note)
		}
		event := dbsvc.Event{
			Id:          e.ID,
			Title:       e.Title,
			Description: e.Description,
			Time:        e.Time,
			Timezone:    e.Timezone,
			Duration:    e.Duration,
			Notes:       notes,
		}
		events = append(events, &event)
	}
	return &dbsvc.AllEventsReply{Events: events, Err: reply.Err}, nil
}

func encodeGRPCAddEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.AddEventResponse)
	return &dbsvc.AddEventReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCShowEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.ShowEventResponse)
	var notes []string
	for _, note := range reply.Event.Notes {
		notes = append(notes, note)
	}
	event := &dbsvc.Event{
		Id:          reply.Event.ID,
		Title:       reply.Event.Title,
		Description: reply.Event.Description,
		Time:        reply.Event.Time,
		Timezone:    reply.Event.Timezone,
		Duration:    reply.Event.Duration,
		Notes:       notes,
	}
	return &dbsvc.ShowEventReply{Event: event, Err: reply.Err}, nil
}

func encodeGRPCUpdateEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.UpdateEventResponse)
	return &dbsvc.UpdateEventReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCDeleteEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.DeleteEventResponse)
	return &dbsvc.DeleteEventReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.ServiceStatusResponse)
	return &dbsvc.ServiceStatusReply{Status: int32(reply.Status), Err: reply.Err}, nil
}
