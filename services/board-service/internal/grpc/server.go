package grpcserver

import (
	"context"
	p "github.com/baydogan/clonello/board-service/internal/proto"
	"github.com/baydogan/clonello/board-service/internal/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type BoardGRPCServer struct {
	p.UnimplementedBoardServiceServer
	Service *services.BoardService
}

func NewBoardGRPCServer(service *services.BoardService) *BoardGRPCServer {
	return &BoardGRPCServer{Service: service}
}

func (s *BoardGRPCServer) GetBoard(ctx context.Context, req *p.GetBoardRequest) (*p.GetBoardResponse, error) {
	board, err := s.Service.GetBoard(ctx, req.BoardId)
	if err != nil {
		return nil, err
	}

	return &p.GetBoardResponse{
		BoardId: board.ID.Hex(),
		Title:   board.Title,
		OwnerId: board.OwnerID,
	}, nil
}

func StartGRPCServer(service *services.BoardService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	p.RegisterBoardServiceServer(grpcServer, NewBoardGRPCServer(service))

	reflection.Register(grpcServer)

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
