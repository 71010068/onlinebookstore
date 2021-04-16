package model

import "onlinebookstore/utils/errmsg"

type Cart struct {
	ID uint `gorm:"primary_key;auto_increment"  json:"id"`

	//User User	`gorm:"foreignKey:UID"  json:"user"`
	UID int `grom:"type:int;not null" json:"uid"`

	Book Book	`gorm:"foreignKey:BID" json:"book"`
	BID int `gorm:"type:int;not null" json:"bid"`
}

// 获取当前用户所有的图书
func GetBookByUID(uid int)([]Cart,int){
	var bookList []Cart
	//db.Preload("Category").Where("cid =?", id).Find(&cateBookList).Count(&total)
	//db.Select("BID , UID",).Find(&Book{}).Where("uid = ?",uid)
	err := db.Select("BID","UID","id").Preload("Book").Where("uid = ?",uid).Find(&bookList).Error
	if err != nil{
		return nil,errmsg.ERROR_USER_NOT_EXIST
	}
	return bookList, errmsg.SUCCESS
}

// 添加图书到当前用户
func AddBookInUID(data *Cart)(code int){
	// gorm中添加的方法：create
	err := db.Create(&data).Error
	if err != nil{ //出错
		return errmsg.ERROR	// 500
	}
	return errmsg.SUCCESS	//200
}

// 删除购物项
// 删除分类
func DeleteCart(id int) int {
	var cart Cart
	err = db.Where("id = ?",id).Delete(&cart).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
