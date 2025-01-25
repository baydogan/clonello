package main

import (
	"github.com/baydogan/clonello/user-service/internal/database"
	"github.com/baydogan/clonello/user-service/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigFile("./configs/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error config file %v", err)
	}

	db := database.Connect()

	router := gin.Default()
	routes.Setup(router, db)

	port := viper.GetString("server.port")
	log.Printf("Server starting in %s port", port)
	log.Fatal(router.Run(port))
}
