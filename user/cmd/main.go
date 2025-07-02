package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/maxamed-cali/go-microservices/proto/gen/proto"

    "github.com/maxamed-cali/go-microservices/user/config"
    "github.com/maxamed-cali/go-microservices/user/internal/db"
    "github.com/maxamed-cali/go-microservices/user/internal/handler"
)

func main() {
    cfg := config.LoadConfig()
    db.InitDB(cfg)

    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterUserServiceServer(grpcServer, &handler.UserHandler{})

    log.Println("UserService running on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
