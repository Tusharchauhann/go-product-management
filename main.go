package main

import (
	"log"
	"product-management/config"
	"product-management/controllers"
	"product-management/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logging
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Connect to the database
	config.ConnectDatabase()

	// Migrate the schema
	err := config.DB.AutoMigrate(&models.Product{}, &models.User{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Initialize the Gin router
	r := gin.Default()

	// Define API routes
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProductByID)
	r.GET("/products", controllers.GetProducts)

	// Run the server
	r.Run(":8081") // Server will run on localhost:8080
}
