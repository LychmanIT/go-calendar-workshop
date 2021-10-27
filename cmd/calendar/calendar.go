package main

import (
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/pkg/calendar"
	"calendarWorkshop/pkg/calendar/endpoints"
	"calendarWorkshop/pkg/calendar/transport"
	"fmt"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "calendarWorkshop/api/v1/pb/calendar"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
)

const (
	defaultHTTPPort   = "8083"
	defaultGRPCPort   = "8084"
	dbContainerName   = "dbsvc"
	defaultDBGRPCPort = "8085"
)

func main() {
	var (
		httpAddr   = net.JoinHostPort("0.0.0.0", envString("HTTP_PORT", defaultHTTPPort))
		grpcAddr   = net.JoinHostPort("0.0.0.0", envString("GRPC_PORT", defaultGRPCPort))
		grpcDBAddr = net.JoinHostPort(envString("DB_CONTAINER_NAME", dbContainerName), envString("DB_GRPC_PORT", defaultDBGRPCPort))
	)

	dbConn, err := grpc.Dial(grpcDBAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var (
		dbGRPCClient = dbsvc.NewDBClient(dbConn)
		service      = calendar.NewService(dbGRPCClient)
		eps          = endpoints.NewEndpointSet(service)
		httpHandler  = transport.NewHTTPHandler(eps)
		grpcServer   = transport.NewGRPCServer(eps)
	)

	var g run.Group
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			log.Println("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			log.Println("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}
	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Println("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			log.Println("transport", "gRPC", "addr", grpcAddr)
			// we add the Go Kit gRPC Interceptor to our gRPC service as it is used by
			// the here demonstrated zipkin tracing middleware.
			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterCalendarServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	log.Println("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
