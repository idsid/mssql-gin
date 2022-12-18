package controllers

import (
	"fmt"
	"mssql-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /books
// Get all products
func GetProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(200, gin.H{
		"data": products,
	})
}

// POST /product
// Create new product
func CreateProduct(c *gin.Context) {
	// Validate input
	var input models.ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product
	product := models.Product{Name: input.Name, Price: input.Price, Quantity: input.Quantity, Status: input.Status}
	models.DB.Create(&product)

	c.JSON(http.StatusCreated, gin.H{
		"data": product,
	})
}
