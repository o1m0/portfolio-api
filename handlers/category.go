package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/db"
	"github.com/o1m0/portfolio-api/models"
)

func GetCategories(c *gin.Context) {
	var Categories []models.Category
	db.DB.Find(&Categories)
	c.JSON(http.StatusOK, Categories)
}

func CreateCategory(c *gin.Context) {
	var input models.Category

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	db.DB.Create(&input)

	c.JSON(http.StatusCreated, input)
}
