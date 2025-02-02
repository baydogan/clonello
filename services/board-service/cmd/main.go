package main

import (
	"github.com/baydogan/clonello/board-service/internal/database"
	"github.com/baydogan/clonello/board-service/internal/grpcserver"
	pb "github.com/baydogan/clonello/board-service/internal/grpcserver/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	database.ConnectMongoDB("mongodb://root:secret@mongo-board:27017")

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error listening port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterBoardServiceServer(grpcServer, &grpcserver.BoardServer{})

	log.Println("✅ Board Service working on gRPC server 50051...")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("cannot start gRPC Server: %v", err)
	}
}
