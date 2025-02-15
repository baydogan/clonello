package grpcserver

import (
	"context"
	"github.com/baydogan/clonello/list-service/internal/database"
	"github.com/baydogan/clonello/list-service/internal/models"
	pb "github.com/baydogan/clonello/proto/pb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type ListServer struct {
	pb.UnimplementedListServiceServer
}

func (s *ListServer) CreateList(ctx context.Context, req *pb.CreateListRequest) (*pb.CreateListResponse, error) {
	collection := database.GetCollection("lists")

	list := models.List{
		ID:      primitive.NewObjectID(),
		Title:   req.Title,
		BoardID: req.BoardId,
	}

	_, err := collection.InsertOne(ctx, list)
	if err != nil {
		log.Printf("List eklenirken hata: %v", err)
		return nil, err
	}

	return &pb.CreateListResponse{ListId: list.ID.Hex()}, nil
}

func (s *ListServer) GetLists(ctx context.Context, req *pb.GetListsRequest) (*pb.GetListsResponse, error) {
	collection := database.GetCollection("lists")

	log.Printf("Fetching lists for board: %s", req.BoardId)

	cursor, err := collection.Find(ctx, bson.M{"board_id": req.BoardId})
	if err != nil {
		log.Printf("Error getting lists:	 %v", err)
		return nil, err
	}

	var lists []*pb.List
	for cursor.Next(ctx) {
		var list models.List
		if err := cursor.Decode(&list); err != nil {
			log.Printf("List decode error: %v", err)
			continue
		}

		log.Printf("Found list: %s - %s", list.ID.Hex(), list.Title)

		lists = append(lists, &pb.List{
			Id:      list.ID.Hex(),
			Title:   list.Title,
			BoardId: list.BoardID,
		})
	}

	log.Printf("Returning %d lists for board %s", len(lists), req.BoardId)

	return &pb.GetListsResponse{Lists: lists}, nil
}
