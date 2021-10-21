package transport

import (
	calendarsvc "calendarWorkshop/api/v1/pb/calendar"
	"calendarWorkshop/models"
	"calendarWorkshop/pkg/calendar/endpoints"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	indexEvent    grpctransport.Handler
	storeEvent    grpctransport.Handler
	showEvent     grpctransport.Handler
	updateEvent   grpctransport.Handler
	deleteEvent   grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) calendarsvc.CalendarServer {
	return &grpcServer{
		indexEvent: grpctransport.NewServer(
			ep.IndexEventEndpoint,
			decodeGRPCIndexEventRequest,
			decodeGRPCIndexEventResponse,
		),
		storeEvent: grpctransport.NewServer(
			ep.StoreEventEndpoint,
			decodeGRPCStoreEventRequest,
			decodeGRPCStoreEventResponse,
		),
		showEvent: grpctransport.NewServer(
			ep.ShowEventEndpoint,
			decodeGRPCShowEventRequest,
			decodeGRPCShowEventResponse,
		),
		updateEvent: grpctransport.NewServer(
			ep.UpdateEventEndpoint,
			decodeGRPCUpdateEventRequest,
			decodeGRPCUpdateEventResponse,
		),
		deleteEvent: grpctransport.NewServer(
			ep.DeleteEventEndpoint,
			decodeGRPCDeleteEventRequest,
			decodeGRPCDeleteEventResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) IndexEvent(ctx context.Context, r *calendarsvc.IndexEventRequest) (*calendarsvc.IndexEventReply, error) {
	_, rep, err := g.indexEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.IndexEventReply), nil
}

func (g *grpcServer) StoreEvent(ctx context.Context, r *calendarsvc.StoreEventRequest) (*calendarsvc.StoreEventReply, error) {
	_, rep, err := g.storeEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.StoreEventReply), nil
}

func (g *grpcServer) ShowEvent(ctx context.Context, r *calendarsvc.ShowEventRequest) (*calendarsvc.ShowEventReply, error) {
	_, rep, err := g.showEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.ShowEventReply), nil
}

func (g *grpcServer) UpdateEvent(ctx context.Context, r *calendarsvc.UpdateEventRequest) (*calendarsvc.UpdateEventReply, error) {
	_, rep, err := g.updateEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.UpdateEventReply), nil
}

func (g *grpcServer) DeleteEvent(ctx context.Context, r *calendarsvc.DeleteEventRequest) (*calendarsvc.DeleteEventReply, error) {
	_, rep, err := g.deleteEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.DeleteEventReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *calendarsvc.ServiceStatusRequest) (*calendarsvc.ServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.ServiceStatusReply), nil
}

func decodeGRPCIndexEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.IndexEventRequest)
	var filters []models.Filter
	for _, f := range req.Filters {
		filters = append(filters, models.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.IndexEventRequest{Filters: filters}, nil
}

func decodeGRPCStoreEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.StoreEventRequest)
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
	return endpoints.StoreEventRequest{Event: event}, nil
}

func decodeGRPCShowEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.ShowEventRequest)
	return endpoints.ShowEventRequest{EventID: req.ID}, nil
}

func decodeGRPCUpdateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.UpdateEventRequest)
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
	req := grpcReq.(*calendarsvc.DeleteEventRequest)
	return endpoints.DeleteEventRequest{EventID: req.Id}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCIndexEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.IndexEventReply)
	var events []models.Event
	for _, e := range reply.Events {
		var notes []string
		for _, note := range e.Notes {
			notes = append(notes, note)
		}
		event := models.Event{
			ID:          e.Id,
			Title:       e.Title,
			Description: e.Description,
			Time:        e.Time,
			Timezone:    e.Timezone,
			Duration:    e.Duration,
			Notes:       notes,
		}
		events = append(events, event)
	}
	return endpoints.IndexEventResponse{Events: events, Err: reply.Err}, nil
}

func decodeGRPCStoreEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.StoreEventReply)
	return endpoints.StoreEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

// TODO reply, ok
func decodeGRPCShowEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.ShowEventReply)
	var notes []string
	for _, note := range reply.Event.Notes {
		notes = append(notes, note)
	}
	event := &models.Event{
		ID:          reply.Event.Id,
		Title:       reply.Event.Title,
		Description: reply.Event.Description,
		Time:        reply.Event.Time,
		Timezone:    reply.Event.Timezone,
		Duration:    reply.Event.Duration,
		Notes:       notes,
	}
	return endpoints.ShowEventResponse{Event: event, Err: reply.Err}, nil
}

func decodeGRPCUpdateEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.UpdateEventReply)
	return endpoints.UpdateEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCDeleteEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.DeleteEventReply)
	return endpoints.DeleteEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

// TODO: some DRY prettify
// TODO 1.18 jenerics
//func grpcToEvent(req interface{}) (*calendar.Event, error) {
//	var notes []string
//	for _, note := range req.Event.Notes {
//		notes = append(notes, note)
//	}
//	event := &calendar.Event{
//		ID:          req.Event.Id,
//		Title:       req.Event.Title,
//		Description: req.Event.Description,
//		Time:        req.Event.Time,
//		Timezone:    req.Event.Timezone,
//		Duration:    req.Event.Duration,
//		Notes:       notes,
//	}
//	return event, nil
//}
