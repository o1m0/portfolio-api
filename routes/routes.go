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

	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	auth.POST("/articles", handlers.CreateArticle)
	auth.PUT("/articles/:id", handlers.UpdateArticle)
	auth.DELETE("/articles/:id", handlers.DeleteArticle)
}
