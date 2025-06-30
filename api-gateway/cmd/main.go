package main

import (
    "log"

    "github.com/gin-gonic/gin"
    "github.com/maxamed-cali/go-microservices/api-gateway/config"
    "github.com/maxamed-cali/go-microservices/api-gateway/internal/routes"
    "github.com/maxamed-cali/go-microservices/api-gateway/internal/service"
)

func main() {
    cfg := config.LoadConfig()

    service.InitUserClient(cfg.UserSvcAddr)

    r := gin.New()
    r.Use(gin.Recovery())

    routes.SetupRoutes(r)

    log.Printf("API Gateway running at http://localhost%s", cfg.Port)
    r.Run(cfg.Port)
}
