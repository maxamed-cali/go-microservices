package handler

import (
    "context"

    pb "github.com/maxamed-cali/go-microservices/proto/gen/proto"
    "github.com/maxamed-cali/go-microservices/user/internal/service"
)

type UserHandler struct {
    pb.UnimplementedUserServiceServer
}

func (s *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
    user, err := service.CreateUser(req.Name, req.Email, req.Password)
    if err != nil {
        return nil, err
    }

    return &pb.UserResponse{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}

func (s *UserHandler) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
    user, err := service.GetUser(req.Id)
    if err != nil {
        return nil, err
    }

    return &pb.UserResponse{
        Id:    user.ID,
        Name:  user.Name,
        Email: user.Email,
    }, nil
}
