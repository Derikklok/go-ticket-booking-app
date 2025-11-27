package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/Derikklok/go-ticket-booking-app/config"
	"github.com/Derikklok/go-ticket-booking-app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect to MongoDB
	config.ConnectMongo()

	// Create router
	router := gin.Default()

	// Enable CORS
	router.Use(cors.Default())

	// Register routes
	routes.UserRoutes(router)

	// Health check
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Go Mongo API running ✅"})
	})

	// Start server
	router.Run(fmt.Sprintf(":%s", port))
}
