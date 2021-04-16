// userHandler 处理器(用于与前端操作)
package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"onlinebookstore/model"
	"onlinebookstore/utils/errmsg"
	"onlinebookstore/utils/validator"
	"strconv"
)

var code int 	//用来接收数据库dao操作中返回的code状态码


// 添加用户
func AddUser(c *gin.Context){
	var data model.User //引入User结构体
	var msg string
	_ = c.ShouldBindJSON(&data)	//模型绑定

	// 进一步判断，将data传入到validator中验证
	msg , code = validator.Validate(&data)
	if code != errmsg.SUCCESS{
		c.JSON(http.StatusOK,gin.H{
			"status" :	code,
			"message":msg,
		})
		//return
		c.Abort()
	}

	// 处理接收的code状态码
	code = model.CheckUser(data.Username)	//先查询Username是否存在
	if code == errmsg.SUCCESS{
		// 成功：将用户名写入数据库，调用dao中的添加用户方法CreateUser()
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR{	//500
		code = errmsg.ERROR_USERNAME_USED	//1001
	}

	// 返回
	c.JSON(http.StatusOK ,gin.H{	//200，网络传输时
		"status" : code,	//返回前端的错误状态码
		//"data"	: data,	//返回到前端去
		"message" : errmsg.GetErrMsg(code),	//返回前端的错误信息
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context){
	id,_ := strconv.Atoi(c.Param("id"))
	var maps = make(map[string]interface{})
	data , code := model.GetUser(id)
	maps["username"] = data.Username
	maps["phone"] = data.Phone
	maps["email"] = data.Email
	maps["role"] = data.Role
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)

	code = errmsg.SUCCESS
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   total,
			"message": errmsg.GetErrMsg(code),
		},
	)

//func GetUsers(c *gin.Context){
//	//接收两个query，从前端传入
//	pageSize , _ := strconv.Atoi(c.Query("pagesize"))
//	pageNum , _ := strconv.Atoi(c.Query("pagenum"))
//	// 查询所有用户列表
//	if pageSize == 0{
//		pageSize = 0
//	}
//	if pageNum == 0{
//		pageNum = 0
//	}
//	//调用User.go中dao操作数据库方法GetUsers获取查询用户列表
//	data,total := model.GetUsers(pageSize,pageNum)
//	code = errmsg.SUCCESS
//	c.JSON(http.StatusOK,gin.H{
//		"status" :  code,
//		"data"	: 	data,
//		"total" :	total,
//		"message" : errmsg.GetErrMsg(code),
//	})
//}
}


// 编辑用户
func EditUser(c *gin.Context){
	// 查询用户名是否已被使用
	var data  model.User

	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	c.ShouldBindJSON(&data)
	// 调用查询用户是否存在CheckUser()
	code = model.CheckUpUser(id,data.Username)
	if code == errmsg.SUCCESS{// 200,可以使用新的用户名
		model.EditUser(id,&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()  //中止
	}
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}
// 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 删除用户
func DeleteUser(c *gin.Context){
	// 从param接收要删除用户的id
	id , _ := strconv.Atoi(c.Param("id"))

	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status" : code,
		"message" : errmsg.GetErrMsg(code),
	})
}

//---------------------------------------------个人信息设置---------------------
// 获取用户信息
func GetPersonal(c *gin.Context)  {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetPersonal(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 更新个人信息
func UpdateProfile(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.UpdatePersonal(id, &data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}