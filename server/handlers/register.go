package handler

import (
	"fmt"

	. "chat_app.com/app/server/types"
	"github.com/gin-gonic/gin"
)

func RegisterHandler(c *gin.Context) {

	var registerBody RegisterRequest
	if err := c.ShouldBindJSON(&registerBody); err != nil {
		fmt.Println("INVALID")
		c.JSON(400, gin.H{"message": "Invalid Request Body"})
		return
	}
	fmt.Println(registerBody.Email)
	c.JSON(201, gin.H{"message": "USER CREATED"})
}
