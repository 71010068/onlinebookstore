// loginHandler 处理器
package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/middleware"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
)

//func Login(c *gin.Context) {
//	var data model.User
//	var token string
//	var code int
//	c.ShouldBindJSON(&data)
//	code = model.CheckLogin(data.Username,data.Password)
//	if code == errmsg.SUCCESS{
//		token ,code = middleware.SetToken(data.Username)
//	}
//	c.JSON(http.StatusOK,gin.H{
//		"status":code,
//		"message":errmsg.GetErrMsg(code),
//		"token" : token,
//	})
//}
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCESS {
		token, code = middleware.SetToken(formData.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

// 前台登录

func LoginFront(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	var token string
	var code int

	formData, code = model.CheckLoginFront(formData.Username, formData.Password)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    formData.Username,
		"id":      formData.ID,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}

type UpToken struct {
	Token string `json:"token"`
}

// 验证token
func CheckToken(c *gin.Context) {
	var Token UpToken
	_ = c.ShouldBindJSON(&Token)

	_, code = middleware.CheckToken(Token.Token)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}