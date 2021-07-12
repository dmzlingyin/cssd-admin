package routers

import (
	"github.com/gin-gonic/gin"

	"cssd-admin/middleware/jwt"
	"cssd-admin/routers/api"
	v1 "cssd-admin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode("debug")

	r.GET("/auth", api.GetAuth)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT()) // All of the /api/v1 APIs need Aothentication first.
	{
		apiv1.GET("/test", v1.GetTestInfo)
		apiv1.GET("/test_select", v1.TestSelect)
	}

	return r
}
