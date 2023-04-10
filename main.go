package main

import (
	"mssql-gin/controllers"
	"mssql-gin/models"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnnectDatabase()

	// CORS Configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://94.237.79.44"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/purchaseTop10", controllers.GetTop10Purchases)

	r.Run()
}
