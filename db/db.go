package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/o1m0/portfolio-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println(host, port, user, password, dbname)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("DBに接続できませんでした")
	}

	db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Category{},
		&models.Work{},
		&models.ArticleCategory{},
		&models.WorkCategory{},
	)

	DB = db
}
