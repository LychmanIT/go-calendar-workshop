package main

import (
	pb "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/pkg/db"
	"calendarWorkshop/pkg/db/endpoints"
	"calendarWorkshop/pkg/db/transport"
	"context"
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultGRPCPort = "8084"
)

func main() {
	var (
		dbConn   *pgxpool.Pool
		grpcAddr = net.JoinHostPort("0.0.0.0", envString("GRPC_PORT", defaultGRPCPort))
	)

	dbConn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect the database ", err.Error())
		os.Exit(1)
	}
	defer dbConn.Close()

	var (
		service    = db.NewService(dbConn)
		eps        = endpoints.NewEndpointSet(service)
		grpcServer = transport.NewGRPCServer(eps)
	)

	var g run.Group
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
			pb.RegisterDBServer(baseServer, grpcServer)
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
