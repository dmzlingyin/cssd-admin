/*
* Gin Middleware for aothentication use JWT
* Author: Lv Wenchao
* Date: 2021-07-12
* Last Modified:
* Last Modified Date:
 */
package jwt

import (
	"net/http"
	"time"

	"cssd-admin/pkg/e"
	"cssd-admin/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token, err := c.Cookie("token")

		if err != nil {
			code = e.ERROR_AUTH
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
