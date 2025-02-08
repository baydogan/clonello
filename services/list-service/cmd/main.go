package main

import (
	"github.com/baydogan/clonello/list-service/internal/proto/pb"
	"log"
	"net"

	"github.com/baydogan/clonello/list-service/internal/database"
	"github.com/baydogan/clonello/list-service/internal/grpcserver"
	"google.golang.org/grpc"
)

func main() {
	database.ConnectMongoDB("mongodb://root:secret@mongo-list:27017")

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Port dinlenemedi: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterListServiceServer(grpcServer, &grpcserver.ListServer{})

	log.Println("List Service gRPC server 50052 portunda çalışıyor...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("gRPC Server başlatılamadı: %v", err)
	}
}
