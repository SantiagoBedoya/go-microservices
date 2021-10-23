package database

import (
	"fmt"
	"os"

	"github.com/SantiagoBedoya/go-microservices/users/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		panic("could not connect with the database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
