package services

import (
	"github.com/baydogan/clonello/api-gateway/internal/grpc"
	"github.com/baydogan/clonello/api-gateway/internal/models"
	"github.com/baydogan/clonello/api-gateway/internal/proto/listpb"
	"log"
)

type ListService struct {
	listClient *grpc.ListClient
}

func NewListService() *ListService {
	return &ListService{
		listClient: grpc.NewListClient(),
	}
}

func (s *ListService) CreateList(req models.CreateListRequest) (*models.CreateListResponse, error) {
	grpcReq := &listpb.CreateListRequest{
		BoardId: req.BoardID,
		Title:   req.Title,
	}

	resp, err := s.listClient.CreateList(grpcReq)
	if err != nil {
		log.Printf("Error calling CreateList: %v", err)
		return nil, err
	}

	return &models.CreateListResponse{ListID: resp.ListId}, nil
}

func (s *ListService) GetLists(boardID string) (*models.GetListsResponse, error) {
	grpcReq := &listpb.GetListsRequest{
		BoardId: boardID,
	}

	resp, err := s.listClient.GetLists(grpcReq)
	if err != nil {
		log.Printf("Error calling GetLists: %v", err)
		return nil, err
	}

	var lists []models.List
	for _, l := range resp.Lists {
		lists = append(lists, models.List{
			ID:      l.Id,
			Title:   l.Title,
			BoardID: l.BoardId,
		})
	}

	return &models.GetListsResponse{Lists: lists}, nil
}
