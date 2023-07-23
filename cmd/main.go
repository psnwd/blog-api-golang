package main

import (
    "github.com/psnwd/blog-api-golang/internal/server"
    "log"
)

func main() {
    err := server.Start()
    if err != nil {
        log.Fatal("Error starting the server: ", err)
    }
}
