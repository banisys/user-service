package main

import (
	"flag"
	"log"
	"net"

	"github.com/banisys/user-service/internal/handlers"
	"google.golang.org/grpc"

	pb "github.com/banisys/user-service/user_service_grpc"
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, &handlers.Server{})

	log.Printf("server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
