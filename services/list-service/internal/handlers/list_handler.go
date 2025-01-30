package handlers

import (
	"context"
	"github.com/baydogan/clonello/list-service/internal/models"
	p "github.com/baydogan/clonello/list-service/internal/proto"
	"github.com/baydogan/clonello/list-service/internal/services"
)

type ListHandler struct {
	Service *services.ListService
	p.UnimplementedListServiceServer
}

func NewListHandler(service *services.ListService) *ListHandler {
	return &ListHandler{Service: service}
}

func (h *ListHandler) CreateList(ctx context.Context, req *p.CreateListRequest) (*p.CreateListResponse, error) {
	list := &models.List{
		BoardID: req.BoardId,
		Title:   req.Title,
	}

	err := h.Service.CreateList(ctx, list)
	if err != nil {
		return nil, err
	}

	return &p.CreateListResponse{ListId: list.ID.Hex()}, nil
}

func (h *ListHandler) GetListsByBoard(ctx context.Context, req *p.GetListsByBoardRequest) (*p.GetListsByBoardResponse, error) {
	lists, err := h.Service.GetListsByBoard(ctx, req.BoardId)
	if err != nil {
		return nil, err
	}

	var pbLists []*p.List
	for _, list := range lists {
		pbLists = append(pbLists, &p.List{
			ListId:  list.ID.Hex(),
			Title:   list.Title,
			BoardId: list.BoardID,
		})
	}

	return &p.GetListsByBoardResponse{Lists: pbLists}, nil
}
