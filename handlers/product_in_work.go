package handlers

import (
	"net/http"
	"strconv"
	"myapi/database"
	"myapi/models"
	"github.com/gin-gonic/gin"
)

func GetProductsInWork(c *gin.Context) {
	var piw []models.ProductInWork

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	database.DB.Preload("User").Preload("Product").Limit(limit).Offset(offset).Find(&piw)
	c.JSON(http.StatusOK, piw)
}

func GetProductInWorkByID(c *gin.Context) {
	id := c.Param("id")
	var piw models.ProductInWork

	if err := database.DB.Preload("User").Preload("Product").First(&piw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	c.JSON(http.StatusOK, piw)
}

func CreateProductInWork(c *gin.Context) {
	var piw models.ProductInWork
	if err := c.ShouldBindJSON(&piw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&piw)
	c.JSON(http.StatusCreated, piw)
}

func UpdateProductInWork(c *gin.Context) {
	id := c.Param("id")
	var piw models.ProductInWork
	if err := database.DB.First(&piw, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	var input models.ProductInWork
	if err := c.ShouldBindJSON(&piw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	piw.UserID = input.UserID
	piw.ProductID = input.ProductID
	database.DB.Save(&piw)
	c.JSON(http.StatusOK, piw)
}

func DeleteProductInWork(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.ProductInWork{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Запись не найдена"})
		return
	}
	database.DB.Delete(&piw)
	c.JSON(http.StatusOK, gin.H{"message": "Запись удалена"})
}
