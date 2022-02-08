package blog

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello post!",
	})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello comment!",
	})
}