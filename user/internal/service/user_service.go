package service

import (
    "github.com/maxamed-cali/go-microservices/user/internal/repository"
    "github.com/maxamed-cali/go-microservices/user/internal/utils"
    "github.com/maxamed-cali/go-microservices/user/internal/model"
)

func CreateUser(name, email, password string) (*model.User, error) {
    hashedPassword, err := utils.HashPassword(password)
    if err != nil {
        return nil, err
    }
    return repository.CreateUser(name, email, hashedPassword)
}

func GetUser(id string) (*model.User, error) {
    return repository.GetUserByID(id)
}
