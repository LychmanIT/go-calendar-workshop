package transport

import (
	authsvc "calendarWorkshop/api/v1/pb/auth"
	"calendarWorkshop/models"
	"calendarWorkshop/pkg/auth/endpoints"
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	login          grpctransport.Handler
	logout         grpctransport.Handler
	verifyToken    grpctransport.Handler
	changeTimezone grpctransport.Handler
	getTimezone    grpctransport.Handler
	serviceStatus  grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) authsvc.AuthServer {
	return &grpcServer{
		login: grpctransport.NewServer(
			ep.LoginEndpoint,
			decodeGRPCLoginRequest,
			encodeGRPCLoginResponse,
		),
		logout: grpctransport.NewServer(
			ep.LogoutEndpoint,
			decodeGRPCLogoutRequest,
			encodeGRPCLogoutResponse,
		),
		verifyToken: grpctransport.NewServer(
			ep.VerifyTokenEndpoint,
			decodeGRPCVerifyTokenRequest,
			encodeGRPCVerifyTokenResponse,
		),
		changeTimezone: grpctransport.NewServer(
			ep.ChangeTimezoneEndpoint,
			decodeGRPCChangeTimezoneRequest,
			encodeGRPCChangeTimezoneResponse,
		),
		getTimezone: grpctransport.NewServer(
			ep.GetTimezoneEndpoint,
			decodeGRPCGetTimezoneRequest,
			encodeGRPCGetTimezoneResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			encodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) AuthLogin(ctx context.Context, r *authsvc.AuthLoginRequest) (*authsvc.AuthLoginReply, error) {
	_, rep, err := g.login.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthLoginReply), nil
}

func (g *grpcServer) AuthLogout(ctx context.Context, r *authsvc.AuthLogoutRequest) (*authsvc.AuthLogoutReply, error) {
	_, rep, err := g.logout.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthLogoutReply), nil
}

func (g *grpcServer) AuthVerifyToken(ctx context.Context, r *authsvc.AuthVerifyTokenRequest) (*authsvc.AuthVerifyTokenReply, error) {
	_, rep, err := g.verifyToken.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthVerifyTokenReply), nil
}

func (g *grpcServer) AuthChangeTimezone(ctx context.Context, r *authsvc.AuthChangeTimezoneRequest) (*authsvc.AuthChangeTimezoneReply, error) {
	_, rep, err := g.changeTimezone.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthChangeTimezoneReply), nil
}

func (g *grpcServer) AuthGetTimezone(ctx context.Context, r *authsvc.AuthGetTimezoneRequest) (*authsvc.AuthGetTimezoneReply, error) {
	_, rep, err := g.getTimezone.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthGetTimezoneReply), nil
}

func (g *grpcServer) AuthServiceStatus(ctx context.Context, r *authsvc.AuthServiceStatusRequest) (*authsvc.AuthServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*authsvc.AuthServiceStatusReply), nil
}

func decodeGRPCLoginRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*authsvc.AuthLoginRequest)
	auth := models.Auth{
		Username: req.Credentials.Username,
		Password: req.Credentials.Password,
	}
	return endpoints.LoginRequest{Credentials: auth}, nil
}

func decodeGRPCLogoutRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*authsvc.AuthLogoutRequest)
	return endpoints.LogoutRequest{Token: req.Token}, nil
}

func decodeGRPCVerifyTokenRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*authsvc.AuthVerifyTokenRequest)
	return endpoints.VerifyTokenRequest{Token: req.Token}, nil
}

func decodeGRPCChangeTimezoneRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*authsvc.AuthChangeTimezoneRequest)
	return endpoints.ChangeTimezoneRequest{Token: req.Token, Timezone: req.Timezone}, nil
}

func decodeGRPCGetTimezoneRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*authsvc.AuthGetTimezoneRequest)
	return endpoints.GetTimezoneRequest{Token: req.Token}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}

func encodeGRPCLoginResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.LoginResponse)
	return &authsvc.AuthLoginReply{Token: reply.Token, Err: reply.Err}, nil
}

func encodeGRPCLogoutResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.LogoutResponse)
	return &authsvc.AuthLogoutReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCVerifyTokenResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.VerifyTokenResponse)
	return &authsvc.AuthVerifyTokenReply{Err: reply.Err}, nil
}

func encodeGRPCChangeTimezoneResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.ChangeTimezoneResponse)
	return &authsvc.AuthChangeTimezoneReply{Status: int32(reply.Status), Err: reply.Err}, nil
}

func encodeGRPCGetTimezoneResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.GetTimezoneResponse)
	return &authsvc.AuthGetTimezoneReply{Timezone: reply.Timezone, Err: reply.Err}, nil
}

func encodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.ServiceStatusResponse)
	return &authsvc.AuthServiceStatusReply{Status: int32(reply.Status), Err: reply.Err}, nil
}
