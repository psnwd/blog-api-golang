package handlers

import (
    "github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Hello, World!",
    })
}
