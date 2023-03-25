package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default
}

func publish(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": codes.CODE_SUCCESS,
	})
}
