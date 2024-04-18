package main

import (
	"cashpay/database"
	"cashpay/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	err := database.InitializeDB("username:password@tcp(hostname:port)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
