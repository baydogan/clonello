package main

import (
	"github.com/baydogan/clonello/api-gateway/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()

	api := router.Group("/api")

	{
		api.POST("/boards", handlers.CreateBoard)
		api.GET("/boards", handlers.GetBoards)

		api.POST("/lists", handlers.CreateList)
		api.GET("/lists", handlers.GetLists)

	}

	log.Println("API Gateway running on port 8082")
	http.ListenAndServe(":8082", router)

}
