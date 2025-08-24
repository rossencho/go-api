package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rossencho/go-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	godotenv.Load()
	dbHost := os.Getenv("MYSQL_HOST")
	dbUser := os.Getenv("MYSQL_USER")
	dbPass := os.Getenv("MYSQL_PASS")
	dbName := os.Getenv("MYSQL_DBNAME")
	dbPort := os.Getenv("MYSQL_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("db connection failed")
	}

	DB = db

	autoMigrate(db)
}

func autoMigrate(conn *gorm.DB) {
	conn.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Discount{},
	)
}
