package handler

import (
    "context"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/maxamed-cali/go-microservices/api-gateway/internal/service"
    pb "github.com/maxamed-cali/go-microservices/proto/gen/proto"
)

func GetUserByID(c *gin.Context) {
    id := c.Param("id")

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    var resp *pb.UserResponse
    var err error

    for i := 0; i < 3; i++ {
        resp, err = service.UserClient.GetUser(ctx, &pb.UserRequest{Id: id})
        if err == nil {
            break
        }
        time.Sleep(100 * time.Millisecond)
    }

    if err != nil {
        c.JSON(http.StatusServiceUnavailable, gin.H{"error": "user service unavailable"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "id":   resp.Id,
        "name": resp.Name,
    })
}
