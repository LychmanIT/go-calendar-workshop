package main

import (
	"calendarWorkshop/models"
	"calendarWorkshop/pkg/db"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

const (
	defaultHTTPPort = "8083"
	defaultGRPCPort = "8084"
)

var logger log.Logger

func main() {
	var dbConn *pgxpool.Pool

	dbConn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("Unable to connect the database ", err.Error())
		os.Exit(1)
	}
	defer dbConn.Close()

	service := db.NewService(dbConn)

	auth := models.Auth{
		Username: "goadmin",
		Password: "gopassword",
	}

	code, err := service.GetUser(context.Background(), &auth)
	if err != nil {
		log.Fatalln("Error ", err)
	}
	log.Println("Success ", code)
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
