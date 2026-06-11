package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/o1m0/portfolio-api/db"
	"github.com/o1m0/portfolio-api/routes"
)

func main() {
	db.Init()

	r := gin.Default()

	routes.Init(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
