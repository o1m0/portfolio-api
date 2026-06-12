package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/handlers"
)

func Init(r *gin.Engine) {
	r.POST("/auth/login", handlers.Login)
}
