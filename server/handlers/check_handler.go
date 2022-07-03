package handler

import (
	. "chat_app.com/app/server/types"
	"github.com/gin-gonic/gin"
)

func CheckHandler(c *gin.Context) {
	var response Response
	c.BindJSON(&response)
	c.JSON(200, gin.H{"message": "connected"})
}
