package main

import (
	"fmt"
	"github.com/jonecoboy/netCheck/internal/core/domain/entity/task"
	"github.com/jonecoboy/netCheck/internal/presentation/router"
	"github.com/jonecoboy/netCheck/internal/utils/config"
	"log"
	"net/http"
)

func main() {
	// Load Configuration
	config.LoadConfig()

	// Initialize MongoDB
	config.InitMongo()
	defer func() {
		if err := config.MongoClient.Disconnect(nil); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
	}()

	// Start task processing in the background
	// remover do entity! por num service domain ou application
	go entity.StartTaskProcessor()

	// Start HTTP server in a separate goroutine
	go func() {
		roteador := router.BuildRouter()
		fmt.Println("Server is running on port", config.AppConfig.ServerPort)
		log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, roteador))
	}()

	// todo start grpc and graphql

	// this line will make it run forver, end stuck the program. otherwise it would run and exit.
	select {} // Blocks forever
}
