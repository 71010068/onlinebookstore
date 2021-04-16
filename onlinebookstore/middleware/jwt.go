// 中间件
package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/utils"
	"onlinebookstore/utils/errmsg"
	"strings"
	"time"
)

// 生成token秘钥
var JwtKey = []byte(utils.JwtKey)

var code int

// 我的请求，结构体用来接收参数
type MyClaims struct {	//username和password和User.go中的保持一致
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string,int) { //返回string是token字符串，int是错误状态码
	expireTime := time.Now().Add(7 * 24 * time.Hour)	 //到期时间
	SetClaims  := MyClaims{
		Username: username,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 有效时间
			Issuer: "onlineookstore" ,
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256,SetClaims)
	token , err := reqClaim.SignedString(JwtKey)
	if err != nil{
		return "",errmsg.ERROR
	}
	return token ,errmsg.SUCCESS
}

// 验证token
//func CheckToken(token string) (*MyClaims,int) {
//	setToken , _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return JwtKey,nil
//	})
//	if key , _ := setToken.Claims.(*MyClaims); setToken.Valid {
//		return key , errmsg.SUCCESS
//	} else {
//		return nil,errmsg.ERROR
//	}
//
//}
func CheckToken(token string) (*MyClaims, int) {
	var claims MyClaims

	setToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (i interface{}, e error) {
		return JwtKey, nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errmsg.ERROR_TOKEN_WRONG
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errmsg.ERROR_TOKEN_RUNTIME
			} else {
				return nil, errmsg.ERROR_TOKEN_TYPE_WRONG
			}
		}
	}
	if setToken != nil {
		if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
			return key, errmsg.SUCCESS
		} else {
			return nil, errmsg.ERROR_TOKEN_WRONG
		}
	}
	return nil, errmsg.ERROR_TOKEN_WRONG
}

// jwt 中间件控制验证写入routers
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")

		if tokenHerder == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder," ",2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key , tCode  := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code =  errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt{
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK,gin.H{
				"code" : code,
				"message" : errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("username",key.Username)
		c.Next()
	}
}