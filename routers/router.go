package routers

import (
	"github.com/gin-gonic/gin"

	v1 "cssd-admin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/test", v1.GetTestInfo)
		apiv1.GET("/test_select", v1.TestSelect)
	}

	return r
}
