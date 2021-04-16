// categoryHandler处理器
package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
	"strconv"
)

// 添加分类
func AddCategory(c *gin.Context){
	var data model.Category //引入User结构体
	_ = c.ShouldBindJSON(&data)	//模型绑定
	// 处理接收的code状态码
	code = model.CheckCategory(data.Name)	//先查询Username是否存在
	if code == errmsg.SUCCESS{
		// 成功：将用户名写入数据库，调用dao中的添加用户方法CreateUser()
		model.CreateCate(&data)
	}
	if code == errmsg.ERROR_CATENAME_USERD{	//2001
		code = errmsg.ERROR_CATENAME_USERD	//2001
	}

	// 返回
	c.JSON(http.StatusOK ,gin.H{	//200，网络传输时
		"status" : code,	//返回前端的错误状态码
		"data"	: data,	//返回到前端去
		"message" : errmsg.GetErrMsg(code),	//返回前端的错误信息
	})
}

//  查询单个分类下的所有图书
func GetCateInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))

	data , code := model.GetCateInfo(id)
	// 返回
	c.JSON(http.StatusOK ,gin.H{	//200，网络传输时
		"status" : code,	//返回前端的错误状态码
		"data"	: data,	//返回到前端去
		"message" : errmsg.GetErrMsg(code),	//返回前端的错误信息
	})
}

// 查询分类列表
func GetCates(c *gin.Context){
	//接收两个query，从前端传入
	pageSize , _ := strconv.Atoi(c.Query("pagesize"))
	pageNum , _ := strconv.Atoi(c.Query("pagenum"))
	// 查询所有用户列表
	if pageSize == 0{
		pageSize = 0
	}
	if pageNum == 0{
		pageNum = 0
	}
	//调用User.go中dao操作数据库方法GetUsers获取查询用户列表
	data,total := model.GetCate(pageSize,pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"data"	: data,
		"total" : total,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 编辑分类
func EditCate(c *gin.Context){
	// 查询用户名是否已被使用
	var data  model.Category

	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	c.ShouldBindJSON(&data)
	// 调用查询用户是否存在CheckUser()
	code = model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS{// 200,可以使用新的用户名
		model.EditCate(id,&data)
	}
	if code == errmsg.ERROR_CATENAME_USERD{
		c.Abort()  //中止
	}
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DeleteCate(c *gin.Context){
	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteCate(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}