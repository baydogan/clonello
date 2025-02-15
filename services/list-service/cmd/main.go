package main

import (
	"log"
	"net"

	"github.com/baydogan/clonello/list-service/internal/database"
	"github.com/baydogan/clonello/list-service/internal/grpcserver"
	"google.golang.org/grpc"

	pb "github.com/baydogan/clonello/proto/pb"
)

func main() {
	database.ConnectMongoDB("mongodb://root:secret@mongo-list:27017")

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Error listening port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterListServiceServer(grpcServer, &grpcserver.ListServer{})

	log.Println("Board Service working on gRPC server 50052...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("cannot start gRPC server %v", err)
	}
}
