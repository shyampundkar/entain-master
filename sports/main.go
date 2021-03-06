package main

import (
	"database/sql"
	"flag"
	"log"
	"net"

	"github.com/shyampundkar/entain-master/sports/db"
	"github.com/shyampundkar/entain-master/sports/proto/sports"
	"github.com/shyampundkar/entain-master/sports/service"
	"google.golang.org/grpc"
)

var (
	grpcEndpoint = flag.String("sportsgrpcEndpoint", "localhost:10000", "gRPC server endpoint")
)

func main() {
	flag.Parse()

	if err := run(); err != nil {
		log.Fatalf("failed running grpc server: %s\n", err)
	}
}

func run() error {
	conn, err := net.Listen("tcp", ":10000")
	if err != nil {
		return err
	}

	sportsDB, err := sql.Open("sqlite3", "./db/sports.db")
	if err != nil {
		return err
	}

	eventsRepo := db.NewEventsRepo(sportsDB)
	if err := eventsRepo.Init(); err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	sports.RegisterSportServer(
		grpcServer,
		service.NewSportsService(
			eventsRepo,
		),
	)

	log.Printf("gRPC server listening on: %s\n", *grpcEndpoint)

	if err := grpcServer.Serve(conn); err != nil {
		return err
	}

	return nil
}
