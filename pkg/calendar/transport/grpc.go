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

func (g *grpcServer) CalendarIndexEvent(ctx context.Context, r *calendarsvc.CalendarIndexEventRequest) (*calendarsvc.CalendarIndexEventReply, error) {
	_, rep, err := g.indexEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarIndexEventReply), nil
}

func (g *grpcServer) CalendarStoreEvent(ctx context.Context, r *calendarsvc.CalendarStoreEventRequest) (*calendarsvc.CalendarStoreEventReply, error) {
	_, rep, err := g.storeEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarStoreEventReply), nil
}

func (g *grpcServer) CalendarShowEvent(ctx context.Context, r *calendarsvc.CalendarShowEventRequest) (*calendarsvc.CalendarShowEventReply, error) {
	_, rep, err := g.showEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarShowEventReply), nil
}

func (g *grpcServer) CalendarUpdateEvent(ctx context.Context, r *calendarsvc.CalendarUpdateEventRequest) (*calendarsvc.CalendarUpdateEventReply, error) {
	_, rep, err := g.updateEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarUpdateEventReply), nil
}

func (g *grpcServer) CalendarDeleteEvent(ctx context.Context, r *calendarsvc.CalendarDeleteEventRequest) (*calendarsvc.CalendarDeleteEventReply, error) {
	_, rep, err := g.deleteEvent.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarDeleteEventReply), nil
}

func (g *grpcServer) CalendarServiceStatus(ctx context.Context, r *calendarsvc.CalendarServiceStatusRequest) (*calendarsvc.CalendarServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*calendarsvc.CalendarServiceStatusReply), nil
}

func decodeGRPCIndexEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.CalendarIndexEventRequest)
	var filters []models.Filter
	for _, f := range req.Filters {
		filters = append(filters, models.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoints.IndexEventRequest{Filters: filters}, nil
}

func decodeGRPCStoreEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.CalendarStoreEventRequest)
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
	req := grpcReq.(*calendarsvc.CalendarShowEventRequest)
	return endpoints.ShowEventRequest{EventID: req.ID}, nil
}

func decodeGRPCUpdateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*calendarsvc.CalendarUpdateEventRequest)
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
	req := grpcReq.(*calendarsvc.CalendarDeleteEventRequest)
	return endpoints.DeleteEventRequest{EventID: req.Id}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func decodeGRPCIndexEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.CalendarIndexEventReply)
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
	reply := grpcReply.(*calendarsvc.CalendarStoreEventReply)
	return endpoints.StoreEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

// TODO reply, ok
func decodeGRPCShowEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.CalendarShowEventReply)
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
	reply := grpcReply.(*calendarsvc.CalendarUpdateEventReply)
	return endpoints.UpdateEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCDeleteEventResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.CalendarDeleteEventReply)
	return endpoints.DeleteEventResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*calendarsvc.CalendarServiceStatusReply)
	return endpoints.ServiceStatusResponse{Status: int(reply.Status), Err: reply.Err}, nil
}

// TODO: some DRY prettify
// TODO 1.18 jenerics
//func grpcToEvent(req interface{}) (*db.Event, error) {
//	var notes []string
//	for _, note := range req.Event.Notes {
//		notes = append(notes, note)
//	}
//	event := &db.Event{
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
