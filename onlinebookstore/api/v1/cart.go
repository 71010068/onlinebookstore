package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
	"strconv"
)

func GetBookByUID(c *gin.Context){
	id , _ := strconv.Atoi(c.Param("id"))

	data,code:= model.GetBookByUID(id)

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data"	: data,
		"message" : errmsg.GetErrMsg(code),
	})
}

func AddBookInUID(c *gin.Context){
	var data model.Cart //引入User结构体
	_ = c.ShouldBindJSON(&data)	//模型绑定

	code = model.AddBookInUID(&data)

	// 返回
	c.JSON(http.StatusOK ,gin.H{	//200，网络传输时
		"status" : code,	//返回前端的错误状态码
		"data"	: data,	//返回到前端去
		"message" : errmsg.GetErrMsg(code),	//返回前端的错误信息
	})

}

func DeleteCart(c *gin.Context)  {
	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCart(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}