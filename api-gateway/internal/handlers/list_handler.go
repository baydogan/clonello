package handlers

import (
	"net/http"

	"github.com/baydogan/clonello/api-gateway/internal/models"
	"github.com/baydogan/clonello/api-gateway/internal/services"
	"github.com/gin-gonic/gin"
)

var listService = services.NewListService()

func CreateList(c *gin.Context) {
	var req models.CreateListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	resp, err := listService.CreateList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create list"})
		return
	}

	c.JSON(http.StatusOK, resp)
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

	c.JSON(http.StatusOK, resp)
}
