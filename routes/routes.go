package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/handlers"
	"github.com/o1m0/portfolio-api/middleware"
)

func Init(r *gin.Engine) {
	r.POST("/auth/login", handlers.Login)
	r.GET("/articles", handlers.GetArticles)
	r.GET("/articles/:id", handlers.DetailArticle)
	r.GET("/categories", handlers.GetCategories)
	r.GET("/works", handlers.GetWorks)
	r.GET("/works/:id", handlers.DetailWork)

	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	auth.POST("/articles", handlers.CreateArticle)
	auth.PUT("/articles/:id", handlers.UpdateArticle)
	auth.DELETE("/articles/:id", handlers.DeleteArticle)
	auth.POST("/categories", handlers.CreateCategory)
	auth.POST("/works", handlers.CreateWork)
	auth.PUT("/works/:id", handlers.UpdateWork)
	auth.DELETE("/works/:id", handlers.DeleteWork)
	auth.POST("/articles/:id/categories", handlers.AddArticleCategory)
	auth.POST("/works/:id/categories", handlers.AddWorkCategory)
}
