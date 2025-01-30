package grpcserver

import (
	p "github.com/baydogan/clonello/list-service/internal/proto"
	"github.com/baydogan/clonello/list-service/internal/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ListGRPCServer struct {
	Service *services.ListService
	p.UnimplementedListServiceServer
}

func StartGRPCServer(service *services.ListService) {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	server := grpc.NewServer()
	p.RegisterListServiceServer(server, &ListGRPCServer{Service: service})

	log.Println("gRPC Server listening on port 50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
