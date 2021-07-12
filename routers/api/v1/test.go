package v1

import (
	"cssd-admin/models"

	"github.com/gin-gonic/gin"
)

func GetTestInfo(c *gin.Context) {

	data := models.TestCreate()
	c.JSON(200, gin.H{
		"data": data,
	})
}

func TestSelect(c *gin.Context) {
	// data := make(map[string]interface{})
	// data["name"] = models.TestSelect()
	data := models.TestSelect()

	c.JSON(200, gin.H{
		"data": data,
	})
}
