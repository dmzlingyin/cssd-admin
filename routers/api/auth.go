/*
* Get Token if the username is exist
* Author: Lv Wenchao
* Date: 2021-07-12
* Last Modified:
* Last Modified Date:
 */
package api

import (
	"net/http"

	"cssd-admin/pkg/e"
	"cssd-admin/pkg/logging"
	"cssd-admin/pkg/util"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// We just use username as part of JWT PAYLOAD now, maybe update later.
type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
}

// username is exist has not validation now, implementate later.
func GetAuth(c *gin.Context) {
	username := c.Query("username")

	a := auth{username}
	valid := validation.Validation{} // 实例化验证对象
	ok, _ := valid.Valid(&a)         // 验证参数是否符合约定

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		token, err := util.GenerateToken(username)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
