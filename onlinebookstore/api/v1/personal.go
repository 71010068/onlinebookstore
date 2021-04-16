package v1

//import (
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"onlinebookstore/model"
//	"onlinebookstore/utils/errmsg"
//	"strconv"
//)
//
//func GetPersonal(c *gin.Context)  {
//	id, _ := strconv.Atoi(c.Param("id"))
//	data, code := model.GetPersonal(id)
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"data":    data,
//		"message": errmsg.GetErrMsg(code),
//	})
//}
//
//// 更新个人信息
//func UpdateProfile(c *gin.Context) {
//	var data model.Personal
//	id, _ := strconv.Atoi(c.Param("id"))
//	_ = c.ShouldBindJSON(&data)
//
//	code = model.UpdatePersonal(id, &data)
//
//	c.JSON(http.StatusOK, gin.H{
//		"status":  code,
//		"message": errmsg.GetErrMsg(code),
//	})
//}