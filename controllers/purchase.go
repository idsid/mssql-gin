package controllers

import (
	"mssql-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTop10Purchases(c *gin.Context) {
	db := models.DB

	var top10 []models.Purchase

	// Execute the SQL query using GORM
	rows, err := db.Table("PCH1").
		Select("ItemCode, Dscription, SUM(Quantity) as Quantity").
		Where("DocDate BETWEEN ? AND ?", "1/1/2022", "1/31/2023").
		Group("ItemCode, Dscription").
		Order("SUM(Quantity) desc").
		Limit(10).
		Rows()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var itemCode, description string
		var quantityUint8 []uint8
		if err := rows.Scan(&itemCode, &description, &quantityUint8); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		quantity, err := strconv.ParseFloat(string(quantityUint8), 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		top10 = append(top10, models.Purchase{ItemCode: itemCode, Dscription: description, Quantity: quantity})
	}

	// Return the result as JSON
	c.JSON(http.StatusOK, top10)
}
