package services

import (
	"context"
	"github.com/baydogan/clonello/list-service/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ListService struct {
	DB *mongo.Collection
}

func NewListService(db *mongo.Database) *ListService {
	return &ListService{
		DB: db.Collection("lists"),
	}
}

func (s *ListService) CreateList(ctx context.Context, list *models.List) error {
	list.CreatedAt = time.Now()
	list.UpdatedAt = time.Now()

	_, err := s.DB.InsertOne(ctx, list)
	return err
}

func (s *ListService) GetListsByBoard(ctx context.Context, boardID string) ([]models.List, error) {
	cursor, err := s.DB.Find(ctx, bson.M{"board_id": boardID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var lists []models.List
	if err := cursor.All(ctx, &lists); err != nil {
		return nil, err
	}

	return lists, nil
}
