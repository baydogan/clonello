package services

import (
	"github.com/baydogan/clonello/api-gateway/internal/grpc"
	"github.com/baydogan/clonello/api-gateway/internal/models"
	"log"

	pb "github.com/baydogan/clonello/proto/pb"
)

type BoardService struct {
	boardClient *grpc.BoardClient
}

func NewBoardService() *BoardService {
	return &BoardService{
		boardClient: grpc.NewBoardClient(),
	}
}

func (s *BoardService) CreateBoard(req models.CreateBoardRequest) (*models.CreateBoardResponse, error) {
	grpcReq := &pb.CreateBoardRequest{
		OwnerId: req.UserID,
		Title:   req.Name,
	}

	resp, err := s.boardClient.CreateBoard(grpcReq)
	if err != nil {
		log.Printf("Error calling CreateBoard: %v", err)
		return nil, err
	}

	return &models.CreateBoardResponse{BoardID: resp.BoardId}, nil
}

func (s *BoardService) GetBoards(userID string) (*models.GetBoardsResponse, error) {
	resp, err := s.boardClient.GetBoards(&pb.GetBoardsRequest{OwnerId: userID})
	if err != nil {
		log.Printf("Error calling GetBoards: %v", err)
		return nil, err
	}

	var boards []models.Board
	for _, b := range resp.Boards {
		boards = append(boards, models.Board{
			ID:      b.Id,
			Name:    b.Title,
			OwnerID: b.OwnerId,
		})
	}

	return &models.GetBoardsResponse{Boards: boards}, nil
}
