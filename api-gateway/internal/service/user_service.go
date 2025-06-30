package service

import (
    "context"
    "log"
    "time"

    pb "github.com/maxamed-cali/go-microservices/proto/gen/proto"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

var UserClient pb.UserServiceClient

func InitUserClient(addr string) {
    // Use context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Dial with context
    conn, err := grpc.DialContext(
        ctx,
        addr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithBlock(), // waits for connection
    )
    if err != nil {
        log.Fatalf("failed to connect to user service: %v", err)
    }

    UserClient = pb.NewUserServiceClient(conn)
}
