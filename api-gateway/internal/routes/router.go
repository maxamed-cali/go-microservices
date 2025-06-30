package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/maxamed-cali/go-microservices/api-gateway/internal/handler"
)

func SetupRoutes(r *gin.Engine) {
    r.GET("/healthz", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    user := r.Group("/users")
    {
        user.GET("/:id", handler.GetUserByID)
    }

    // More route groups: /orders, /products, etc.
}
