package main

import (
	"github.com/baydogan/clonello/list-service/internal/database"
	grpcserver "github.com/baydogan/clonello/list-service/internal/grpc"
	"github.com/baydogan/clonello/list-service/internal/services"
	"log"
)

func main() {
	uri := "mongodb://mongo-list:27017"
	username := "root"
	password := "secret"
	dbName := "list_service_db"

	db, err := database.ConnectMongoDB(uri, username, password, dbName)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	listService := services.NewListService(db)
	log.Println("List Service running on gRPC port :50052")

	grpcserver.StartGRPCServer(listService)

}
