package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
	"strconv"
)

// 添加图书
func AddBook(c *gin.Context){
	var data model.Book //引入User结构体
	_ = c.ShouldBindJSON(&data)	//模型绑定

	code = model.CreateBook(&data)

	// 返回
	c.JSON(http.StatusOK ,gin.H{	//200，网络传输时
		"status" : code,	//返回前端的错误状态码
		"data"	: data,	//返回到前端去
		"message" : errmsg.GetErrMsg(code),	//返回前端的错误信息
	})
}

//  查询分类下所有图书
func GetCateBook(c *gin.Context){
	pageSize , _ := strconv.Atoi(c.Query("pagesize"))
	pageNum , _ := strconv.Atoi(c.Query("pagenum"))
	id , _ := strconv.Atoi(c.Param("id"))
	// 查询所有用户列表
	if pageSize == 0{
		pageSize = 0
	}
	if pageNum == 0{
		pageNum = 0
	}
	data,code,total := model.GetCateBook(id ,pageSize,pageNum)

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data"	: data,
		"total"	:total,
		"message" : errmsg.GetErrMsg(code),
	})
}
//  根据查询单个图书
func GetBook(c *gin.Context){
	id , _ := strconv.Atoi(c.Param("id"))
	data , code := model.GetBookInfo(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data"	: data,
		"message" : errmsg.GetErrMsg(code),
	})
}


// 查询图书列表
func GetBooks(c *gin.Context){
	//接收两个query，从前端传入
	pageSize , _ := strconv.Atoi(c.Query("pagesize"))
	pageNum , _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	// 查询所有用户列表
	if pageSize == 0{
		pageSize = 0
	}
	if pageNum == 0{
		pageNum = 0
	}
	//调用User.go中dao操作数据库方法GetUsers获取查询用户列表
	data , code ,total := model.GetBooks(title,pageSize,pageNum)

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data"	: data,
		"total" : total,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 编辑图书
func EditBook(c *gin.Context){
	// 查询用户名是否已被使用
	var data  model.Book

	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	c.ShouldBindJSON(&data)

	code = model.EditBook(id,&data)

	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 删除图书
func DeleteBook(c *gin.Context){
	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteBook(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}