package routes

import (
	"github.com/baydogan/clonello/user-service/internal/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(router *gin.Engine, db *gorm.DB) {
	userHandler := handlers.NewUserHandler(db)

	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userHandler.RegisterUser)
		userRoutes.POST("/login", userHandler.LoginUser)
	}
}
