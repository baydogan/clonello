package grpcserver

import (
	"context"
	p "github.com/baydogan/clonello/board-service/internal/proto"
	"google.golang.org/grpc"
	"log"
)

type BoardClient struct {
	Client p.BoardServiceClient
}

func NewBoardClient() *BoardClient {
	conn, err := grpc.Dial("board-service:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Board Service: %v", err)
	}

	client := p.NewBoardServiceClient(conn)
	return &BoardClient{Client: client}
}

func (b *BoardClient) GetBoard(boardID string) (*p.GetBoardResponse, error) {
	req := &p.GetBoardRequest{BoardId: boardID}
	return b.Client.GetBoard(context.Background(), req)
}
