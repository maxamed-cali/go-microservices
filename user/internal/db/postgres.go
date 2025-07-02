package db

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "github.com/maxamed-cali/go-microservices/user/config"
    "github.com/maxamed-cali/go-microservices/user/internal/model"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    DB = db

    if err := db.AutoMigrate(&model.User{}); err != nil {
        log.Fatalf("auto migration failed: %v", err)
    }
}
