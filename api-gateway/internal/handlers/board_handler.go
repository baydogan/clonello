package handlers

import (
	"net/http"

	"github.com/baydogan/clonello/api-gateway/internal/models"
	"github.com/baydogan/clonello/api-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

var boardService = services.NewBoardService()

func CreateBoard(c *gin.Context) {
	var req models.CreateBoardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := boardService.CreateBoard(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create board"})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetBoards(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	resp, err := boardService.GetBoards(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch boards"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
