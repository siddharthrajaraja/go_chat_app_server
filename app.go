package main

import (
	. "chat_app.com/app/server/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/check-connection", CheckHandler)
	router.POST("/register", RegisterHandler)
	router.Run("localhost:8080")
}
