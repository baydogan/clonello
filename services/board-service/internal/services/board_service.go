package services

import (
	"context"
	"github.com/baydogan/clonello/board-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type BoardService struct {
	DB *mongo.Collection
}

func NewBoardService(db *mongo.Database) *BoardService {
	return &BoardService{
		DB: db.Collection("boards"),
	}
}

func (s *BoardService) CreateBoard(ctx context.Context, board *models.Board) error {

	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()

	_, err := s.DB.InsertOne(ctx, board)
	return err
}

func (s *BoardService) GetAllBoards(ctx context.Context) ([]models.Board, error) {
	cursor, err := s.DB.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var boards []models.Board
	if err := cursor.All(ctx, &boards); err != nil {
		return nil, err
	}

	return boards, nil
}
