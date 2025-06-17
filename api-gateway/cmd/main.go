package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pb "github.com/maxamed-cali/go-microservices/proto/gen/proto" // update to match your actual module path
	"google.golang.org/grpc"
)


func main() {
	// Connect to user service over gRPC
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to user service: %v", err)
	}
	defer conn.Close()

	userClient := pb.NewUserServiceClient(conn)

	r := gin.New()
	r.Use(gin.Recovery()) // handles panics gracefully

	// Health check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// User endpoint (REST → gRPC)
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Set timeout for gRPC call
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		var resp *pb.UserResponse
		var err error

		for i := 0; i < 3; i++ { // retry up to 3 times
			resp, err = userClient.GetUser(ctx, &pb.UserRequest{Id: id})
			if err == nil {
				break
			}
			log.Printf("gRPC call failed (attempt %d): %v", i+1, err)
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
	})

	log.Println("API Gateway is running on http://localhost:8080")
	r.Run(":8080")
}
