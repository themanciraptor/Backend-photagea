package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	userapi "github.com/themanciraptor/Backend-photagea/API/user"
	"google.golang.org/grpc"
)

const (
	port = ":5555"
)

func main() {
	// Sign in to DB
	db, err := sql.Open("mysql", "ezdev:ForkmeMuthafukka@/photagea?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Listen on the correct port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Setup gRPC server
	s := grpc.NewServer()

	us := userapi.UnimplementedUserServiceServer{}
	userapi.RegisterUserServiceServer(s, &us)

	// Serve RPC
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve RPC")
	}
}
