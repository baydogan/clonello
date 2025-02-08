package handlers

import (
	"github.com/baydogan/clonello/api-gateway/internal/models"
	"github.com/baydogan/clonello/api-gateway/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

var listService *services.ListService

func CreateList(c *gin.Context) {
	var req models.CreateListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := listService.CreateList(req.BoardID, req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create list"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"list_id": resp.ListId})
}

func GetLists(c *gin.Context) {
	boardID := c.Query("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "board_id is required"})
		return
	}

	resp, err := listService.GetLists(boardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch lists"})
		return
	}

	var lists []models.List
	for _, l := range resp.Lists {
		lists = append(lists, models.List{
			ID:      l.Id,
			Title:   l.Title,
			BoardID: l.BoardId,
		})
	}

	c.JSON(http.StatusOK, models.GetListsResponse{Lists: lists})
}
