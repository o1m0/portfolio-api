package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/db"
	"github.com/o1m0/portfolio-api/models"
)

func GetWorks(c *gin.Context) {
	var works []models.Work
	query := db.DB.Preload("Categories")

	search := c.Query("search")
	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	sort := c.Query("sort")
	if sort != "" {
		query = query.Order(sort)
	}

	categoryID := c.Query("category_id")
	if categoryID != "" {
		query = query.Joins("JOIN work_categories ON work_categories.work_id = works.id").
			Where("work_categories.category_id = ?", categoryID)
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	offset := (pageInt - 1) * limitInt

	query = query.Limit(limitInt).Offset(offset)

	query.Find(&works)
	c.JSON(http.StatusOK, works)
}

func DetailWork(c *gin.Context) {
	var work models.Work
	id := c.Param("id")
	result := db.DB.Preload("Categories").First(&work, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, work)

}

func CreateWork(c *gin.Context) {
	var input models.Work

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	db.DB.Create(&input)

	c.JSON(http.StatusCreated, input)
}

func UpdateWork(c *gin.Context) {

	type WorkInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		GithubURL   string `json:"github_url"`
		DemoURL     string `json:"demo_url"`
		ImageURL    string `json:"image_url"`
	}

	var work models.Work
	var input WorkInput

	id := c.Param("id")

	result := db.DB.First(&work, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	work.Title = input.Title
	work.Description = input.Description
	work.GithubURL = input.GithubURL
	work.DemoURL = input.DemoURL
	work.ImageURL = input.ImageURL

	result = db.DB.Save(&work)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "記事が更新できませんでした"})
		return
	}

	c.JSON(http.StatusOK, work)

}

func DeleteWork(c *gin.Context) {

	var work models.Work

	id := c.Param("id")

	result := db.DB.First(&work, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "記事が見つかりません"})
		return
	}

	result = db.DB.Delete(&work)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "削除できませんでした"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "削除しました"})
}

func AddWorkCategory(c *gin.Context) {

	type AddCategoryInput struct {
		CategoryID uint `json:"category_id"`
	}

	id := c.Param("id")

	var input AddCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "入力が正しくありません"})
		return
	}

	var work models.Work
	result := db.DB.First(&work, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "workが見つかりません"})
		return
	}

	var category models.Category
	result = db.DB.First(&category, input.CategoryID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "カテゴリーが見つかりません"})
		return
	}

	idInt, _ := strconv.Atoi(id)
	WorkCategory := models.WorkCategory{
		WorkID:     uint(idInt),
		CategoryID: input.CategoryID,
	}

	db.DB.Create(&WorkCategory)

	c.JSON(http.StatusCreated, WorkCategory)

}
