package grpc

import (
	"context"
	"github.com/baydogan/clonello/api-gateway/internal/proto/listpb"
	"google.golang.org/grpc"
	"log"
)

type ListClient struct {
	client listpb.ListServiceClient
}

func NewListClient() *ListClient {
	conn, err := grpc.Dial("board-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Board Service: %v", err)
	}

	return &ListClient{
		client: listpb.NewListServiceClient(conn),
	}
}

func (l *ListClient) CreateList(req *listpb.CreateListRequest) (*listpb.CreateListResponse, error) {
	resp, err := l.client.CreateList(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateList: %v", err)
		return nil, err
	}
	log.Printf("✅ List Created: %s", resp.ListId)
	return resp, nil
}

func (l *ListClient) GetLists(req *listpb.GetListsRequest) (*listpb.GetListsResponse, error) {
	resp, err := l.client.GetLists(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetLists: %v", err)
		return nil, err
	}
	log.Printf("✅ Lists Retrieved: %d lists found", len(resp.Lists))
	return resp, nil
}
