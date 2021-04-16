package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
)

func UpLoad(c *gin.Context){
	file , fileHeader , _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size

	url , code := model.UpLoadFile(file,fileSize)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
		"url" : url,
	})
}
