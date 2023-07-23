package server

import (
    "github.com/gin-gonic/gin"
    "github.com/psnwd/blog-api-golang/internal/routes"
    "github.com/psnwd/blog-api-golang/internal/websocket"
    "github.com/joho/godotenv"
    "log"
    "net/http"
    "os"
)

func Start() error {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Create a Gin router
    router := gin.Default()

    // Initialize API routes
    routes.Init(router)

    // Initialize WebSocket route
    websocket.Init(router)

    // Start the server
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if PORT environment variable is not set
    }
    return router.Run(":" + port)
}
