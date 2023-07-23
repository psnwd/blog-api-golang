package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/psnwd/blog-api-golang/internal/handlers"
)

func Init(router *gin.Engine) {
    // Define your API routes and handlers here
    router.GET("/hello", handlers.HelloHandler)
}
