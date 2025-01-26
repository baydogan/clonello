package handlers

import (
	"context"
	"github.com/baydogan/clonello/board-service/internal/models"
	"github.com/baydogan/clonello/board-service/internal/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type BoardHandler struct {
	Service *services.BoardService
}

func NewBoardHandler(service *services.BoardService) *BoardHandler {
	return &BoardHandler{Service: service}
}

func (h *BoardHandler) CreateBoard(c *gin.Context) {
	var req struct {
		Title   string `json:"title" binding:"required"`
		OwnerID string `json:"owner_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board := &models.Board{
		ID:        primitive.NewObjectID(),
		Title:     req.Title,
		OwnerID:   req.OwnerID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := h.Service.CreateBoard(context.Background(), board); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	c.JSON(http.StatusCreated, board)
}

func (h *BoardHandler) GetAllBoards(c *gin.Context) {
	boards, err := h.Service.GetAllBoards(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch boards"})
		return
	}

	c.JSON(http.StatusOK, boards)
}
