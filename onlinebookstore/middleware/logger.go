// 日志的中间件
package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	rotalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc{

	//	写入文件中
	filePath := "log/log"
	linkName := "latest_log.log"
	src,err := os.OpenFile(filePath , os.O_RDWR|os.O_CREATE,0755)
	if err != nil{
		fmt.Println("err:",err)
	}

	logger := logrus.New()	// 实例化对象

	logger.Out = src	// 用src输出到log.log中

	logger.SetLevel(logrus.DebugLevel) // 设置日志级别

	// 日志分割
	logWriter ,_ := rotalog.New(
		filePath+"%Y%m%d.log",
		rotalog.WithMaxAge(7*24*time.Hour),
		rotalog.WithRotationTime(24*time.Hour),
		// 软链接
		rotalog.WithLinkName(linkName),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}) // 先实例化

	// 新增一个hock
	logger.AddHook(Hook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms",int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName,err := os.Hostname()
		if err != nil{
			hostName = "unknown"
		}
		// 请求码
		statusCode := c.Writer.Status()
		// 客户端请求的Ip
		clientIp := c.ClientIP()
		// 客户端浏览器信息
		userAgent := c.Request.UserAgent()
		// 请求过来的文件长度
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		// 请求方法
		method := c.Request.Method
		// 请求路径
		path := c.Request.RequestURI


		entry := logger.WithFields(logrus.Fields{
			"HostName":hostName,
			"status" : statusCode,
			"SpendTime" : spendTime,
			"Ip" : clientIp,
			"Method" : method,
			"Path"	: path,
			"DataSize"	: dataSize,
			"Agent" :	userAgent,
		})
		if len(c.Errors) > 0{	// 记录系统内部的错误
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {	// 记录状态码的错误
			entry.Error()
		}else if statusCode >= 400{	// 记录警告
			entry.Warn()
		}else {
			entry.Info()
		}
	}
}
