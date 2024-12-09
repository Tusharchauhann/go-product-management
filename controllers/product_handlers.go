package controllers

import (
	"net/http"
	"product-management/config"
	"product-management/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// POST /products
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Bind the incoming JSON to product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists in the user_s table by user_id
	var user models.User
	if err := config.DB.First(&user, "user_id = ?", product.UserID).Error; err != nil {
		// If no user is found, return an error
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found in user_s table"})
		return
	}

	// Ensure that product_images is an array
	if len(product.ProductImages) == 0 {
		product.ProductImages = []string{} // Default to an empty array if no images are provided
	}

	// Debugging output to check how the data is being processed
	logrus.Infof("Product data: %+v", product)

	// Save the product to the database
	if err := config.DB.Create(&product).Error; err != nil {
		logrus.Error("Failed to create product:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	logrus.Info("Product created successfully")
	c.JSON(http.StatusCreated, product)
}

// GET /products/:id
func GetProductByID(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	// Find the product by ID
	if err := config.DB.First(&product, id).Error; err != nil {
		logrus.Warn("Product not found with ID:", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
	}

	logrus.Info("Product retrieved with ID:", id)
	c.JSON(http.StatusOK, product)
}

// GET /products
func GetProducts(c *gin.Context) {
	var products []models.Product

	// Retrieve all products from the productm table
	if err := config.DB.Find(&products).Error; err != nil {
		logrus.Error("Failed to retrieve products:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	logrus.Info("Retrieved all products")
	c.JSON(http.StatusOK, products)
}
