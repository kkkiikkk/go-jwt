package database

import (
	"fmt"
	"strconv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"github.com/kkkiikkk/go-jwt/config"
	"github.com/kkkiikkk/go-jwt/model"
)

func ConnectDB() {
	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		panic("failed to parse database port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"))
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{}, &model.Image{})

	user := model.User{
        Username: "testuser",
        Password: "$2a$10$LUh2eLV1JN4Gt/rjTGH1WuM6Pyhc1rLuFVXhYh.IHUPcnzfV1FgfO", // 123456
    }

    result := DB.Create(&user)
    if result.Error != nil {
        fmt.Println("Failed to create test user:", result.Error)
    } else {
        fmt.Println("Test user created successfully")
    }
	fmt.Println("Database Migrated")
}