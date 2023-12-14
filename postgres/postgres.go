package postgres

import (
	"Go-ToDo/config"
	"Go-ToDo/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DatabaseHost, config.DatabasePort, config.DatabaseUserName, config.DatabasePassword, config.DatabaseName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	DB = db
}

func AutoMigrate(connection *gorm.DB) {
	if err := connection.Debug().AutoMigrate(&models.Task{}); err != nil {
		return
	}
}
