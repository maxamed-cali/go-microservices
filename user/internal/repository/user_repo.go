package repository

import (
    "github.com/google/uuid"
    "time"

    "github.com/maxamed-cali/go-microservices/user/internal/db"
    "github.com/maxamed-cali/go-microservices/user/internal/model"
)

func CreateUser(name, email, password string) (*model.User, error) {
    user := &model.User{
        ID:        uuid.New().String(),
        Name:      name,
        Email:     email,
        Password:  password,
        CreatedAt: time.Now(),
    }

    if err := db.DB.Create(user).Error; err != nil {
        return nil, err
    }

    return user, nil
}

func GetUserByID(id string) (*model.User, error) {
    var user model.User
    if err := db.DB.First(&user, "id = ?", id).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
