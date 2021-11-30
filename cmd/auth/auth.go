package main

import (
	pb "calendarWorkshop/api/v1/pb/auth"
	dbsvc "calendarWorkshop/api/v1/pb/db"
	"calendarWorkshop/pkg/auth"
	"calendarWorkshop/pkg/auth/endpoints"
	"calendarWorkshop/pkg/auth/transport"
	"context"
	"fmt"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/go-redis/redis/v8"
	"github.com/oklog/run"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	defaultGRPCPort           = "8084"
	defaultRedisPort          = "6379"
	defaultRedisContainerName = "Redis"
	dbContainerName           = "dbsvc"
	defaultDBGRPCPort         = "8085"
)

func main() {
	var (
		grpcAddr   = net.JoinHostPort("0.0.0.0", envString("GRPC_PORT", defaultGRPCPort))
		grpcDBAddr = net.JoinHostPort(envString("DB_CONTAINER_NAME", dbContainerName), envString("DB_GRPC_PORT", defaultDBGRPCPort))
		redisAddr  = net.JoinHostPort(envString("REDIS_CONTAINER_NAME", defaultRedisContainerName), envString("REDIS_PORT", defaultRedisPort))
	)

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	defer func() {
		err := redisClient.Close()
		if err != nil {
			panic(err)
		}
	}()

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	dbConn, err := grpc.Dial(grpcDBAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var (
		dbGRPCClient = dbsvc.NewDBClient(dbConn)
		service      = auth.NewService(redisClient, dbGRPCClient)
		eps          = endpoints.NewEndpointSet(service)
		grpcServer   = transport.NewGRPCServer(eps)
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
			pb.RegisterAuthServer(baseServer, grpcServer)
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
