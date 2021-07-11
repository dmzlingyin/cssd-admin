package v1

import (
	"github.com/gin-gonic/gin"
)

func GetTestInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "what a fucking genius",
	})
}
