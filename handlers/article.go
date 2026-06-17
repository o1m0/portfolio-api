package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/db"
	"github.com/o1m0/portfolio-api/models"
)

func GetArticles(c *gin.Context) {
	var articles []models.Article
	query := db.DB

	search := c.Query("search")
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	sort := c.Query("sort")
	if sort != "" {
		query = query.Order(sort)
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt

	query = query.Limit(limitInt).Offset(offset)

	query.Find(&articles)
	c.JSON(http.StatusOK, articles)
}

func CreateArticle(c *gin.Context) {
	var input models.Article
	userID, _ := c.Get("user_id")
	input.UserID = uint(userID.(float64))

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	db.DB.Create(&input)

	c.JSON(http.StatusCreated, input)
}

func DetailArticle(c *gin.Context) {
	var article models.Article

	id := c.Param("id")

	result := db.DB.First(&article, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, article)

}

func UpdateArticle(c *gin.Context) {

	type ArticleInput struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	var article models.Article
	var input ArticleInput

	id := c.Param("id")

	result := db.DB.First(&article, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	article.Title = input.Title
	article.Body = input.Body

	result = db.DB.Save(&article)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "記事が更新できませんでした"})
		return
	}

	c.JSON(http.StatusOK, article)

}

func DeleteArticle(c *gin.Context) {

	var article models.Article

	id := c.Param("id")

	result := db.DB.First(&article, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	result = db.DB.Delete(&article)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "削除できませんでした"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}
