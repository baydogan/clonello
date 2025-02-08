package grpcserver

import (
	"context"
	"log"

	"github.com/baydogan/clonello/board-service/internal/database"
	"github.com/baydogan/clonello/board-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	pb "github.com/baydogan/clonello/proto/pb"
)

type BoardServer struct {
	pb.UnimplementedBoardServiceServer
}

func (s *BoardServer) CreateBoard(ctx context.Context, req *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	collection := database.GetCollection("boards")

	board := models.Board{
		ID:      primitive.NewObjectID(),
		Title:   req.Title,
		OwnerID: req.OwnerId,
	}

	_, err := collection.InsertOne(ctx, board)
	if err != nil {
		log.Printf("Error during create board %v", err)
		return nil, err
	}

	return &pb.CreateBoardResponse{BoardId: board.ID.Hex()}, nil
}

func (s *BoardServer) GetBoards(ctx context.Context, req *pb.GetBoardsRequest) (*pb.GetBoardsResponse, error) {
	collection := database.GetCollection("boards")

	cursor, err := collection.Find(ctx, bson.M{"owner_id": req.OwnerId})
	if err != nil {
		log.Printf("Board'lar getirilirken hata: %v", err)
		return nil, err
	}

	var boards []*pb.Board
	for cursor.Next(ctx) {
		var board models.Board
		if err := cursor.Decode(&board); err != nil {
			log.Printf("Board decode error: %v", err)
			continue
		}
		boards = append(boards, &pb.Board{
			Id:      board.ID.Hex(),
			Title:   board.Title,
			OwnerId: board.OwnerID,
		})
	}

	return &pb.GetBoardsResponse{Boards: boards}, nil
}
