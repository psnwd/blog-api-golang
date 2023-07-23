package websocket

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "net/http"
)

var upgrader = websocket.Upgrader{}

func Init(router *gin.Engine) {
    router.GET("/ws", handleWebSocket)
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        // Handle error
        return
    }
    defer conn.Close()

    for {
        // Read message from the client
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            // Handle error
            break
        }

        // Process the received message (e.g., echo it back)
        err = conn.WriteMessage(messageType, message)
        if err != nil {
            // Handle error
            break
        }
    }
}
