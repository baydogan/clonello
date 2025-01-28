package grpc

import (
	"github.com/baydogan/clonello/shared/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type GRPCClients struct {
	ListClient pb.ListServiceClient
	CardClient pb.CardServiceClient
}

func NewGRPCClients(listServiceAddr, cardServiceAddr string) *GRPCClients {
	listConn, err := grpc.Dial(listServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to List Service: %v", err)
	}
	listClient := pb.NewListServiceClient(listConn)

	cardConn, err := grpc.Dial(cardServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Card Service: %v", err)
	}
	cardClient := pb.NewCardServiceClient(cardConn)

	return &GRPCClients{
		ListClient: listClient,
		CardClient: cardClient,
	}
}
