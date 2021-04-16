// 解决跨域问题
package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc{

	// return cors.Default()

	return cors.New(
		cors.Config{
			AllowAllOrigins:true,   // 允许所有的跨域
			// AllowOrigins: []string{"*"},

			//AllowMethods:     []string{"*"},  // 允许的请求方法
			AllowMethods:     []string{"GET","POST","PUT","DELETE","OPTIONS"},  // 允许的请求方法

			//AllowHeaders:     []string{"Origin"},
			AllowHeaders:     []string{"*"},

			ExposeHeaders:    []string{"Content-Length","Authorization","Content-Type"},

			AllowCredentials: true,
			
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//},

			MaxAge: 12 * time.Hour,  // 域请求的持续时间
		})

}