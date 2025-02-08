package grpc

import (
	"context"
	boardpb "github.com/baydogan/clonello/api-gateway/internal/proto/boardpb"
	"google.golang.org/grpc"
	"log"
)

type BoardClient struct {
	client boardpb.BoardServiceClient
}

func NewBoardClient() *BoardClient {
	conn, err := grpc.Dial("board-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to Board Service: %v", err)
	}

	return &BoardClient{
		client: boardpb.NewBoardServiceClient(conn),
	}
}

func (b *BoardClient) CreateBoard(req *boardpb.CreateBoardRequest) (*boardpb.CreateBoardResponse, error) {
	resp, err := b.client.CreateBoard(context.Background(), req)
	if err != nil {
		log.Printf("Error calling CreateBoard: %v", err)
		return nil, err
	}
	log.Printf("logged")
	return resp, nil
}

func (b *BoardClient) GetBoards(req *boardpb.GetBoardsRequest) (*boardpb.GetBoardsResponse, error) {
	resp, err := b.client.GetBoards(context.Background(), req)
	if err != nil {
		log.Printf("Error calling GetBoards: %v", err)
		return nil, err
	}
	return resp, nil
}
