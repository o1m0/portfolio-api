package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/handlers"
	"github.com/o1m0/portfolio-api/middleware"
)

func Init(r *gin.Engine) {
	r.POST("/auth/login", handlers.Login)

	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
}
