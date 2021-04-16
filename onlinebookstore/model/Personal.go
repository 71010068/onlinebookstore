package model

//import "onlinebookstore/utils/errmsg"
//
//type Personal struct {
//	ID     int    `gorm:"foreignKey" json:"id"`
//	Name   string `gorm:"type:varchar(20)" json:"name"`
//	Desc   string `gorm:"type:varchar(200)" json:"desc"`
//	Qqchat string `gorm:"type:varchar(200)" json:"qq_chat"`
//	Wechat string `gorm:"type:varchar(100)" json:"wechat"`
//	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
//}
//
//// 获取个人信息设置
//func GetPersonal(id int) (Personal, int) {
//	var personal Personal
//	err = db.Where("ID = ?", id).First(&personal).Error
//	if err != nil {
//		return personal, errmsg.ERROR
//	}
//	return personal, errmsg.SUCCESS
//}
//
//// 更新个人信息设置
//func UpdatePersonal(id int, data *Personal) int {
//	var personal Personal
//	err = db.Model(&personal).Where("ID = ?", id).Updates(&data).Error
//	if err != nil {
//		return errmsg.ERROR
//	}
//	return errmsg.SUCCESS
//}