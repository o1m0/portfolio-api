package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/o1m0/portfolio-api/db"
	"github.com/o1m0/portfolio-api/models"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	godotenv.Load()

	db.Init()

	password := os.Getenv("SEED_PASSWORD")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := models.User{
		Email:    os.Getenv("SEED_EMAIL"),
		Password: string(hashed),
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("ユーザー作成成功")
}
