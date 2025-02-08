package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"

	pb "github.com/baydogan/clonello/proto/pb"
)

type ListClient struct {
	client pb.ListServiceClient
}

func NewListClient() *ListClient {
	conn, err := grpc.Dial("list-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to List Service: %v", err)
	}

	return &ListClient{
		client: pb.NewListServiceClient(conn),
	}
}

func (l *ListClient) CreateList(req *pb.CreateListRequest) (*pb.CreateListResponse, error) {
	resp, err := l.client.CreateList(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateList: %v", err)
		return nil, err
	}
	log.Printf("List Created: %s", resp.ListId)
	return resp, nil
}

func (l *ListClient) GetLists(req *pb.GetListsRequest) (*pb.GetListsResponse, error) {
	resp, err := l.client.GetLists(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetLists: %v", err)
		return nil, err
	}

	log.Printf("Lists Retrieved: %d lists found", len(resp.Lists))
	return resp, nil
}
