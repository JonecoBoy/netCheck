package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JonecoBoy/netCheck/config"
	"github.com/JonecoBoy/netCheck/router"
	"github.com/JonecoBoy/netCheck/task"
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
	tasks.StartTaskProcessor()

	// Initialize router
	roteador := router.NewRouter()

	fmt.Println("Server is running on port", config.AppConfig.ServerPort)
	http.ListenAndServe(":"+config.AppConfig.ServerPort, roteador)
}
