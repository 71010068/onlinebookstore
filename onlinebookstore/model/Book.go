package model

import (
	"gorm.io/gorm"
	"onlinebookstore/utils/errmsg"
)

type Book struct {
	Category Category `gorm:"foreignKey:Cid"  json:"category"` //Book从属于Category分类
	gorm.Model
	Cid int  `gorm:"type:int;not null"  json:"cid"`   // 指向Category中ID
	Title  string  `gorm:"type:varchar(100);primary_key;not null"  json:"title"`	// 书名
	Desc string `gorm:type:"varchar(200)"  json:"desc"`	 // 图书描述
	Author string  `gorm:"type:varchar(100);not null"  json:"author"`  // 作者
	Price string  `gorm:"type:longtext"  json:"price"`  // 价格
	Img string  `gorm:"type:varchar(100);not null"  json:"img"`  //图片
}

// -------------------编写数据库对应的DAO操作的方法--------------------


// 新增图书
func CreateBook(data *Book)(code int){

	// gorm中添加的方法：create
	err := db.Create(&data).Error
	if err != nil{ //出错
		return errmsg.ERROR	// 500
	}
	return errmsg.SUCCESS	//200

}

//  查询分类下的所有图书
func GetCateBook(id int ,pageSize int, pageNum int) ([]Book,int ,int64) {
	var cateBookList []Book
	var total int64
	//err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1) * pageSize).
	//	Where("cid = ?",id).Find(&cateBookList).Count(&total).Error
	db.Preload("Category").Where("cid =?", id).Find(&cateBookList).Count(&total)
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where(
		"cid =?", id).Find(&cateBookList).Error
	if err != nil{
		return nil,errmsg.ERROR_CATE_NOT_EXIST,0
	}
	return cateBookList, errmsg.SUCCESS,total
}

//  查询单个图书
func GetBookInfo(id int) (Book,int) {
	var book Book
	err := db.Preload("Category").Where("id = ?",id).First(&book).Error
	if err != nil{
		return book, errmsg.ERROR_BOOK_NOT_EXIST
	}
	return book,errmsg.SUCCESS
}

// 查询图书列表
//pageSize：页数， pageNum：当前多少页     通过前端传
func GetBooks(title string, pageSize int,pageNum int) ([]Book , int ,int64) {
	var bookList []Book
	var err error
	var total int64
	if title == ""{
		// 数据库查找
		//err = db.Preload("Category").Find(&bookList).Count(&total).Limit(pageSize).Offset((pageNum-1) * pageSize).Error
		err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("Created_At DESC").
			Preload("Category").Find(&bookList).Error
		// 单独计数
		db.Model(&bookList).Count(&total)
		//if err != nil && err != gorm.ErrRecordNotFound{
		//	return nil , errmsg.ERROR,0
		//}
		//return bookList,errmsg.SUCCESS,total
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
		return bookList, errmsg.SUCCESS, total
	}
	//err = db.Preload("Category").Where("title LIKE ?",title +"%").Find(&bookList).Count(&total).
	//	Limit(pageSize).Offset((pageNum-1) * pageSize).Error
	//if err != nil && err != gorm.ErrRecordNotFound{
	//	return nil , errmsg.ERROR,0
	//}
	//return bookList,errmsg.SUCCESS,total
	err = db.Limit(pageSize).Offset((pageNum-1)*pageSize).Order("Created_At DESC").
		Preload("Category").Where("title LIKE ?", title+"%", ).Find(&bookList).Error
	// 单独计数
	db.Model(&bookList).Where("title LIKE ?", title+"%").Count(&total)

	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return bookList, errmsg.SUCCESS, total
}


// 编辑图书信息
func EditBook(id int,data *Book) int{
	var book Book
	var maps = make(map[string]interface{})
	//maps["name"] = data.Category.Name
	maps["cid"] = data.Cid
	maps["title"] = data.Title
	maps["author"] = data.Author
	maps["desc"] = data.Desc
	maps["price"] = data.Price
	maps["img"] = data.Img
	//db.Model(&User{}).Where("id = ?",id).Updates(maps)
	err := db.Model(&book).Where("id = ?",id).Updates(maps).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS // 200
}

// 删除图书
func DeleteBook(id int) int {
	var book Book
	err := db.Where("id = ?",id).Delete(&book).Error
	if err != nil{
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200
}
