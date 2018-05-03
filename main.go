package main

import (
	"./server"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	server.InitRouter(router)
	router.Run(":8080")
}
