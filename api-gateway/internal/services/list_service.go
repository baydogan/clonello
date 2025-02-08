package services

import (
	"context"
	"github.com/baydogan/clonello/api-gateway/internal/proto/listpb"
	"google.golang.org/grpc"
	"log"
)

type ListService struct {
	listClient listpb.ListServiceClient
}

func NewListService(conn *grpc.ClientConn) *ListService {
	return &ListService{
		listClient: listpb.NewListServiceClient(conn),
	}
}

func (s *ListService) CreateList(boardID, title string) (*listpb.CreateListResponse, error) {
	resp, err := s.listClient.CreateList(context.Background(), &listpb.CreateListRequest{
		BoardId: boardID,
		Title:   title,
	})
	if err != nil {
		log.Printf("Error calling CreateList: %v", err)
		return nil, err
	}
	return resp, nil
}

func (s *ListService) GetLists(boardID string) (*listpb.GetListsResponse, error) {
	resp, err := s.listClient.GetLists(context.Background(), &listpb.GetListsRequest{
		BoardId: boardID,
	})
	if err != nil {
		log.Printf("Error calling GetLists: %v", err)
		return nil, err
	}
	return resp, nil
}
