package grpc

import (
	"context"
	pb "github.com/baydogan/clonello/api-gateway/internal/proto/pb"
	"google.golang.org/grpc"
	"log"
)

type BoardClient struct {
	client pb.BoardServiceClient
}

func NewBoardClient() *BoardClient {
	conn, err := grpc.Dial("board-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Board Service: %v", err)
	}

	return &BoardClient{
		client: pb.NewBoardServiceClient(conn),
	}
}

func (b *BoardClient) CreateBoard(req *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	resp, err := b.client.CreateBoard(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateBoard: %v", err)
		return nil, err
	}
	log.Printf("logged")
	return resp, nil
}

func (b *BoardClient) GetBoards(req *pb.GetBoardsRequest) (*pb.GetBoardsResponse, error) {
	resp, err := b.client.GetBoards(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetBoards: %v", err)
		return nil, err
	}
	return resp, nil
}
