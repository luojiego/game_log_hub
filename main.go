package main

import (
	"game_log_hub/api"
	"game_log_hub/database"
	"log"
)

func main() {
	// Initialize the database
	database.Initialize()

	// Setup the router
	router := api.SetupRouter()

	// Start the server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
