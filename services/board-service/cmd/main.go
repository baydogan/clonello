package main

import (
	"github.com/baydogan/clonello/board-service/internal/database"
	"github.com/baydogan/clonello/board-service/internal/grpc"
	"github.com/baydogan/clonello/board-service/internal/handlers"
	"github.com/baydogan/clonello/board-service/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	uri := "mongodb://mongo:27017"
	username := "root"
	password := "secret"
	dbName := "board_service_db"

	db, err := database.ConnectMongoDB(uri, username, password, dbName)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	boardService := services.NewBoardService(db)
	grpcClients := grpc.NewGRPCClients("list-service:50052", "card-service:50053")
	boardHandler := handlers.NewBoardHandler(boardService, grpcClients)

	router := gin.Default()
	router.POST("/boards", boardHandler.CreateBoard)
	router.GET("/boards", boardHandler.GetAllBoards)

	log.Println("Board Service running on port :8081")
	if err := router.Run(":8081"); err != nil {
		log.Fatalf("Failed to start Board Service: %v", err)
	}
}
