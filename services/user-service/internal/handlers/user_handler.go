package handlers

import (
	"github.com/baydogan/clonello/user-service/internal/models"
	"github.com/baydogan/clonello/user-service/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		Service: services.NewUserService(db),
	}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Username: req.Username,
		FullName: req.FullName,
		Email:    req.Email,
	}

	if err := h.Service.RegisterUser(user, req.Password); err != nil {
		if err == services.ErrUserAlreadyExists {
			c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {

	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.LoginUser(req.Username, req.Password)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while logging in"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Login successful",
		"username":  user.Username,
		"full_name": user.FullName,
		"email":     user.Email,
	})
}
